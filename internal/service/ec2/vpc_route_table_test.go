package ec2_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	sdkacctest "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
)

func TestAccVPCRouteTable_basic(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRouteTable(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVPCRouteTable_Disappears_subnetAssociation(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_subnetAssociation(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceRouteTable(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"
	destinationCidr3 := "10.4.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr2, destinationCidr3),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "gateway_id", igwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToInstance(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	instanceResourceName := "aws_instance.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4Instance(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "instance_id", instanceResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv6ToEgressOnlyInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv6EgressOnlyInternetGateway(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr, "egress_only_gateway_id", eoigwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Verify that expanded form of the destination CIDR causes no diff.
				Config:   testAccVPCRouteTableConfig_ipv6EgressOnlyInternetGateway(rName, "::0/0"),
				PlanOnly: true,
			},
		},
	})
}

func TestAccVPCRouteTable_tags(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_tags1(rName, "key1", "value1"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVPCRouteTableConfig_tags2(rName, "key1", "value1updated", "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "2"),
					resource.TestCheckResourceAttr(resourceName, "tags.key1", "value1updated"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_tags1(rName, "key2", "value2"),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.key2", "value2"),
				),
			},
		},
	})
}

func TestAccVPCRouteTable_requireRouteDestination(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccVPCRouteTableConfig_noDestination(rName),
				ExpectError: regexp.MustCompile("error creating route: one of `cidr_block"),
			},
		},
	})
}

func TestAccVPCRouteTable_requireRouteTarget(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccVPCRouteTableConfig_noTarget(rName),
				ExpectError: regexp.MustCompile(`error creating route: one of .*\begress_only_gateway_id\b`),
			},
		},
	})
}

func TestAccVPCRouteTable_Route_mode(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVPCRouteTableConfig_modeNoBlocks(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "gateway_id", igwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVPCRouteTableConfig_modeZeroed(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToTransitGateway(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	tgwResourceName := "aws_ec2_transit_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4TransitGateway(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "transit_gateway_id", tgwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	if testing.Short() {
		t.Skip("skipping long-running test in short mode")
	}

	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	vpceResourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheckELBv2GatewayLoadBalancer(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID, "elasticloadbalancing"),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4EndpointID(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "vpc_endpoint_id", vpceResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToCarrierGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	cgwResourceName := "aws_ec2_carrier_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheckWavelengthZoneAvailable(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4CarrierGateway(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "carrier_gateway_id", cgwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToLocalGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	lgwDataSourceName := "data.aws_ec2_local_gateway.first"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "0.0.0.0/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); acctest.PreCheckOutpostsOutposts(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4LocalGateway(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "local_gateway_id", lgwDataSourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToVPCPeeringConnection(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4PeeringConnection(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "vpc_peering_connection_id", pcxResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_vgwRoutePropagation(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	vgwResourceName1 := "aws_vpn_gateway.test1"
	vgwResourceName2 := "aws_vpn_gateway.test2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_vgwPropagation(rName, vgwResourceName1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(resourceName, "propagating_vgws.*", vgwResourceName1, "id"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_vgwPropagation(rName, vgwResourceName2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "1"),
					resource.TestCheckTypeSetElemAttrPair(resourceName, "propagating_vgws.*", vgwResourceName2, "id"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_conditionalCIDRBlock(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"
	destinationIpv6Cidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, false),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "gateway_id", igwResourceName, "id"),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr, true),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationIpv6Cidr, "gateway_id", igwResourceName, "id"),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_ipv4ToNatGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	ngwResourceName := "aws_nat_gateway.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "10.2.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv4NATGateway(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr, "nat_gateway_id", ngwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_IPv6ToNetworkInterface_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eniResourceName := "aws_network_interface.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr := "::/0"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_ipv6NetworkInterfaceUnattached(rName, destinationCidr),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr, "network_interface_id", eniResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_IPv4ToNetworkInterfaces_unattached(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eni1ResourceName := "aws_network_interface.test1"
	eni2ResourceName := "aws_network_interface.test2"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "0"),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr1, destinationCidr2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "network_interface_id", eni1ResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", eni2ResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr2, destinationCidr1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 3),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "2"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "network_interface_id", eni1ResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "network_interface_id", eni2ResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_modeZeroed(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 1),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
		},
	})
}

func TestAccVPCRouteTable_vpcMultipleCIDRs(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_multipleCIDRs(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_vpcClassicLink(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_classicLink(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_gatewayVPCEndpoint(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	var vpce ec2.VpcEndpoint
	resourceName := "aws_route_table.test"
	vpceResourceName := "aws_vpc_endpoint.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_gatewayEndpoint(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckVPCEndpointExists(ctx, vpceResourceName, &vpce),
					testAccCheckRouteTableWaitForVPCEndpointRoute(ctx, &routeTable, &vpce),
					// Refresh the route table once the VPC endpoint route is present.
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_multipleRoutes(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	eoigwResourceName := "aws_egress_only_internet_gateway.test"
	igwResourceName := "aws_internet_gateway.test"
	instanceResourceName := "aws_instance.test"
	pcxResourceName := "aws_vpc_peering_connection.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)
	destinationCidr1 := "10.2.0.0/16"
	destinationCidr2 := "10.3.0.0/16"
	destinationCidr3 := "10.4.0.0/16"
	destinationCidr4 := "2001:db8::/122"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_multiples(rName,
					"cidr_block", destinationCidr1, "gateway_id", igwResourceName,
					"cidr_block", destinationCidr2, "instance_id", instanceResourceName,
					"ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "instance_id", instanceResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_multiples(rName,
					"cidr_block", destinationCidr1, "vpc_peering_connection_id", pcxResourceName,
					"cidr_block", destinationCidr3, "instance_id", instanceResourceName,
					"ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr1, "vpc_peering_connection_id", pcxResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "instance_id", instanceResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr4, "egress_only_gateway_id", eoigwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				Config: testAccVPCRouteTableConfig_multiples(rName,
					"ipv6_cidr_block", destinationCidr4, "vpc_peering_connection_id", pcxResourceName,
					"cidr_block", destinationCidr3, "gateway_id", igwResourceName,
					"cidr_block", destinationCidr2, "instance_id", instanceResourceName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 5),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "3"),
					testAccCheckRouteTableRoute(resourceName, "ipv6_cidr_block", destinationCidr4, "vpc_peering_connection_id", pcxResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr3, "gateway_id", igwResourceName, "id"),
					testAccCheckRouteTableRoute(resourceName, "cidr_block", destinationCidr2, "instance_id", instanceResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccVPCRouteTable_prefixListToInternetGateway(t *testing.T) {
	ctx := acctest.Context(t)
	var routeTable ec2.RouteTable
	resourceName := "aws_route_table.test"
	igwResourceName := "aws_internet_gateway.test"
	plResourceName := "aws_ec2_managed_prefix_list.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(ctx, t); testAccPreCheckManagedPrefixList(ctx, t) },
		ErrorCheck:               acctest.ErrorCheck(t, ec2.EndpointsID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckRouteTableDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCRouteTableConfig_prefixListInternetGateway(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckRouteTableExists(ctx, resourceName, &routeTable),
					testAccCheckRouteTableNumberOfRoutes(&routeTable, 2),
					acctest.MatchResourceAttrRegionalARN(resourceName, "arn", "ec2", regexp.MustCompile(`route-table/.+$`)),
					acctest.CheckResourceAttrAccountID(resourceName, "owner_id"),
					resource.TestCheckResourceAttr(resourceName, "propagating_vgws.#", "0"),
					resource.TestCheckResourceAttr(resourceName, "route.#", "1"),
					testAccCheckRouteTablePrefixListRoute(resourceName, plResourceName, "gateway_id", igwResourceName, "id"),
					resource.TestCheckResourceAttr(resourceName, "tags.%", "1"),
					resource.TestCheckResourceAttr(resourceName, "tags.Name", rName),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckRouteTableExists(ctx context.Context, n string, v *ec2.RouteTable) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No ID is set")
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn()

		routeTable, err := tfec2.FindRouteTableByID(ctx, conn, rs.Primary.ID)

		if err != nil {
			return err
		}

		*v = *routeTable

		return nil
	}
}

func testAccCheckRouteTableDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn()

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_route_table" {
				continue
			}

			_, err := tfec2.FindRouteTableByID(ctx, conn, rs.Primary.ID)

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("Route table %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckRouteTableNumberOfRoutes(routeTable *ec2.RouteTable, n int) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		if len := len(routeTable.Routes); len != n {
			return fmt.Errorf("Route Table has incorrect number of routes (Expected=%d, Actual=%d)\n", n, len)
		}

		return nil
	}
}

func testAccCheckRouteTableRoute(resourceName, destinationAttr, destination, targetAttr, targetResourceName, targetResourceAttr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[targetResourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", targetResourceName)
		}

		target := rs.Primary.Attributes[targetResourceAttr]
		if target == "" {
			return fmt.Errorf("Not found: %s.%s", targetResourceName, targetResourceAttr)
		}

		return resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
			destinationAttr: destination,
			targetAttr:      target,
		})(s)
	}
}

func testAccCheckRouteTablePrefixListRoute(resourceName, prefixListResourceName, targetAttr, targetResourceName, targetResourceAttr string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rsPrefixList, ok := s.RootModule().Resources[prefixListResourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", prefixListResourceName)
		}

		destination := rsPrefixList.Primary.Attributes["id"]
		if destination == "" {
			return fmt.Errorf("Not found: %s.id", prefixListResourceName)
		}

		rsTarget, ok := s.RootModule().Resources[targetResourceName]
		if !ok {
			return fmt.Errorf("Not found: %s", targetResourceName)
		}

		target := rsTarget.Primary.Attributes[targetResourceAttr]
		if target == "" {
			return fmt.Errorf("Not found: %s.%s", targetResourceName, targetResourceAttr)
		}

		return resource.TestCheckTypeSetElemNestedAttrs(resourceName, "route.*", map[string]string{
			"destination_prefix_list_id": destination,
			targetAttr:                   target,
		})(s)
	}
}

// testAccCheckRouteTableWaitForVPCEndpointRoute returns a TestCheckFunc which waits for
// a route to the specified VPC endpoint's prefix list to appear in the specified route table.
func testAccCheckRouteTableWaitForVPCEndpointRoute(ctx context.Context, routeTable *ec2.RouteTable, vpce *ec2.VpcEndpoint) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Conn()

		resp, err := conn.DescribePrefixListsWithContext(ctx, &ec2.DescribePrefixListsInput{
			Filters: tfec2.BuildAttributeFilterList(map[string]string{
				"prefix-list-name": aws.StringValue(vpce.ServiceName),
			}),
		})
		if err != nil {
			return err
		}

		if resp == nil || len(resp.PrefixLists) == 0 {
			return fmt.Errorf("Prefix List not found")
		}

		plId := aws.StringValue(resp.PrefixLists[0].PrefixListId)

		err = retry.RetryContext(ctx, 3*time.Minute, func() *retry.RetryError {
			resp, err := conn.DescribeRouteTablesWithContext(ctx, &ec2.DescribeRouteTablesInput{
				RouteTableIds: []*string{routeTable.RouteTableId},
			})
			if err != nil {
				return retry.NonRetryableError(err)
			}
			if resp == nil || len(resp.RouteTables) == 0 {
				return retry.NonRetryableError(fmt.Errorf("Route Table not found"))
			}

			for _, route := range resp.RouteTables[0].Routes {
				if aws.StringValue(route.DestinationPrefixListId) == plId {
					return nil
				}
			}

			return retry.RetryableError(fmt.Errorf("Route not found"))
		})

		return err
	}
}

func testAccVPCRouteTableConfig_basic(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id
}
`, rName)
}

func testAccVPCRouteTableConfig_subnetAssociation(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block        = "10.1.1.0/24"
  vpc_id            = aws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table_association" "test" {
  route_table_id = aws_route_table.test.id
  subnet_id      = aws_subnet.test.id
}
`, rName))
}

func testAccVPCRouteTableConfig_ipv4InternetGateway(rName, destinationCidr1, destinationCidr2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block = %[2]q
    gateway_id = aws_internet_gateway.test.id
  }

  route {
    cidr_block = %[3]q
    gateway_id = aws_internet_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr1, destinationCidr2)
}

func testAccVPCRouteTableConfig_ipv6EgressOnlyInternetGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block                       = "10.1.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    ipv6_cidr_block        = %[2]q
    egress_only_gateway_id = aws_egress_only_internet_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_ipv4Instance(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
		testAccLatestAmazonNatInstanceAMIConfig(),
		acctest.ConfigAvailableAZsNoOptIn(),
		acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
		fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block        = "10.1.1.0/24"
  vpc_id            = aws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_instance" "test" {
  ami           = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_id     = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block  = %[2]q
    instance_id = aws_instance.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr))
}

func testAccVPCRouteTableConfig_tags1(rName, tagKey1, tagValue1 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    %[2]q = %[3]q
  }
}
`, rName, tagKey1, tagValue1)
}

func testAccVPCRouteTableConfig_tags2(rName, tagKey1, tagValue1, tagKey2, tagValue2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    %[2]q = %[3]q
    %[4]q = %[5]q
  }
}
`, rName, tagKey1, tagValue1, tagKey2, tagValue2)
}

func testAccVPCRouteTableConfig_ipv4PeeringConnection(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id      = aws_vpc.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block                = %[2]q
    vpc_peering_connection_id = aws_vpc_peering_connection.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_vgwPropagation(rName, vgwResourceName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpn_gateway" "test1" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_vpn_gateway" "test2" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_vpn_gateway_attachment" "test" {
  vpc_id         = aws_vpc.test.id
  vpn_gateway_id = %[2]s.id
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  propagating_vgws = [aws_vpn_gateway_attachment.test.vpn_gateway_id]

  tags = {
    Name = %[1]q
  }
}
`, rName, vgwResourceName)
}

func testAccVPCRouteTableConfig_noDestination(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAvailableAZsNoOptIn(),
		acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
		acctest.ConfigLatestAmazonLinuxHVMEBSAMI(),
		fmt.Sprintf(`
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    instance_id = aws_instance.test.id
  }

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block        = "10.1.1.0/24"
  vpc_id            = aws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_instance" "test" {
  ami           = data.aws_ami.amzn-ami-minimal-hvm-ebs.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_id     = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}
`, rName))
}

func testAccVPCRouteTableConfig_noTarget(rName string) string {
	return fmt.Sprintf(`
resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block = "10.1.0.0/16"
  }

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc" "test" {
  cidr_block = "10.2.0.0/16"

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccVPCRouteTableConfig_modeNoBlocks(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccVPCRouteTableConfig_modeZeroed(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route = []

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccVPCRouteTableConfig_ipv4TransitGateway(rName, destinationCidr string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptInDefaultExclude(), fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = "10.1.1.0/24"
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway" "test" {
  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_transit_gateway_vpc_attachment" "test" {
  subnet_ids         = [aws_subnet.test.id]
  transit_gateway_id = aws_ec2_transit_gateway.test.id
  vpc_id             = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block         = %[2]q
    transit_gateway_id = aws_ec2_transit_gateway_vpc_attachment.test.transit_gateway_id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr))
}

func testAccVPCRouteTableConfig_ipv4EndpointID(rName, destinationCidr string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
data "aws_caller_identity" "current" {}

resource "aws_vpc" "test" {
  cidr_block = "10.10.10.0/25"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  availability_zone = data.aws_availability_zones.available.names[0]
  cidr_block        = cidrsubnet(aws_vpc.test.cidr_block, 2, 0)
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_lb" "test" {
  load_balancer_type = "gateway"
  name               = %[1]q

  subnet_mapping {
    subnet_id = aws_subnet.test.id
  }
}

resource "aws_vpc_endpoint_service" "test" {
  acceptance_required        = false
  allowed_principals         = [data.aws_caller_identity.current.arn]
  gateway_load_balancer_arns = [aws_lb.test.arn]

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc_endpoint" "test" {
  service_name      = aws_vpc_endpoint_service.test.service_name
  subnet_ids        = [aws_subnet.test.id]
  vpc_endpoint_type = aws_vpc_endpoint_service.test.service_type
  vpc_id            = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block      = %[2]q
    vpc_endpoint_id = aws_vpc_endpoint.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr))
}

func testAccVPCRouteTableConfig_ipv4CarrierGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_carrier_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block         = %[2]q
    carrier_gateway_id = aws_ec2_carrier_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_ipv4LocalGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
data "aws_ec2_local_gateways" "all" {}

data "aws_ec2_local_gateway" "first" {
  id = tolist(data.aws_ec2_local_gateways.all.ids)[0]
}

data "aws_ec2_local_gateway_route_tables" "all" {}

data "aws_ec2_local_gateway_route_table" "first" {
  local_gateway_route_table_id = tolist(data.aws_ec2_local_gateway_route_tables.all.ids)[0]
}

resource "aws_vpc" "test" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.0.0.0/24"
  vpc_id     = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_local_gateway_route_table_vpc_association" "example" {
  local_gateway_route_table_id = data.aws_ec2_local_gateway_route_table.first.id
  vpc_id                       = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block       = %[2]q
    local_gateway_id = data.aws_ec2_local_gateway.first.id
  }

  tags = {
    Name = %[1]q
  }

  depends_on = [aws_ec2_local_gateway_route_table_vpc_association.example]
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_conditionalIPv4IPv6(rName, destinationCidr, destinationIpv6Cidr string, ipv6Route bool) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  assign_generated_ipv6_cidr_block = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

locals {
  ipv6             = %[4]t
  destination      = %[2]q
  destination_ipv6 = %[3]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block      = local.ipv6 ? null : local.destination
    ipv6_cidr_block = local.ipv6 ? local.destination_ipv6 : null
    gateway_id      = aws_internet_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr, destinationIpv6Cidr, ipv6Route)
}

func testAccVPCRouteTableConfig_ipv4NATGateway(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_id     = aws_vpc.test.id

  map_public_ip_on_launch = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_eip" "test" {
  vpc = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_nat_gateway" "test" {
  allocation_id = aws_eip.test.id
  subnet_id     = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }

  depends_on = [aws_internet_gateway.test]
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block     = %[2]q
    nat_gateway_id = aws_nat_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_ipv6NetworkInterfaceUnattached(rName, destinationCidr string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block                       = "10.1.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block      = "10.1.1.0/24"
  vpc_id          = aws_vpc.test.id
  ipv6_cidr_block = cidrsubnet(aws_vpc.test.ipv6_cidr_block, 8, 1)

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    ipv6_cidr_block      = %[2]q
    network_interface_id = aws_network_interface.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr)
}

func testAccVPCRouteTableConfig_ipv4TwoNetworkInterfacesUnattached(rName, destinationCidr1, destinationCidr2 string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block = "10.1.1.0/24"
  vpc_id     = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test1" {
  subnet_id = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test2" {
  subnet_id = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    cidr_block           = %[2]q
    network_interface_id = aws_network_interface.test1.id
  }

  route {
    cidr_block           = %[3]q
    network_interface_id = aws_network_interface.test2.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationCidr1, destinationCidr2)
}

func testAccVPCRouteTableConfig_multipleCIDRs(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc_ipv4_cidr_block_association" "test" {
  vpc_id     = aws_vpc.test.id
  cidr_block = "172.2.0.0/16"
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc_ipv4_cidr_block_association.test.vpc_id

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccVPCRouteTableConfig_classicLink(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block         = "10.1.0.0/16"
  enable_classiclink = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}
`, rName)
}

func testAccVPCRouteTableConfig_gatewayEndpoint(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

data "aws_region" "current" {}

resource "aws_vpc_endpoint" "test" {
  vpc_id          = aws_vpc.test.id
  service_name    = "com.amazonaws.${data.aws_region.current.name}.s3"
  route_table_ids = [aws_route_table.test.id]
}
`, rName)
}

func testAccVPCRouteTableConfig_multiples(rName,
	destinationAttr1, destinationValue1, targetAttribute1, targetValue1,
	destinationAttr2, destinationValue2, targetAttribute2, targetValue2,
	destinationAttr3, destinationValue3, targetAttribute3, targetValue3 string) string {
	return acctest.ConfigCompose(
		testAccLatestAmazonNatInstanceAMIConfig(),
		acctest.ConfigAvailableAZsNoOptInDefaultExclude(),
		acctest.AvailableEC2InstanceTypeForAvailabilityZone("data.aws_availability_zones.available.names[0]", "t3.micro", "t2.micro"),
		fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block                       = "10.1.0.0/16"
  assign_generated_ipv6_cidr_block = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc" "target" {
  cidr_block = "10.0.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  cidr_block        = "10.1.1.0/24"
  vpc_id            = aws_vpc.test.id
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_egress_only_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_instance" "test" {
  ami           = data.aws_ami.amzn-ami-nat-instance.id
  instance_type = data.aws_ec2_instance_type_offering.available.instance_type
  subnet_id     = aws_subnet.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_vpc_peering_connection" "test" {
  vpc_id      = aws_vpc.test.id
  peer_vpc_id = aws_vpc.target.id
  auto_accept = true

  tags = {
    Name = %[1]q
  }
}

locals {
  routes = [
    {
      destination_attr  = %[2]q
      destination_value = %[3]q
      target_attr       = %[4]q
      target_value      = %[5]s.id
    },
    {
      destination_attr  = %[6]q
      destination_value = %[7]q
      target_attr       = %[8]q
      target_value      = %[9]s.id
    },
    {
      destination_attr  = %[10]q
      destination_value = %[11]q
      target_attr       = %[12]q
      target_value      = %[13]s.id
    }
  ]
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  dynamic "route" {
    for_each = local.routes
    content {
      # Destination.
      cidr_block      = (route.value["destination_attr"] == "cidr_block") ? route.value["destination_value"] : null
      ipv6_cidr_block = (route.value["destination_attr"] == "ipv6_cidr_block") ? route.value["destination_value"] : null

      # Target.
      carrier_gateway_id        = (route.value["target_attr"] == "carrier_gateway_id") ? route.value["target_value"] : null
      egress_only_gateway_id    = (route.value["target_attr"] == "egress_only_gateway_id") ? route.value["target_value"] : null
      gateway_id                = (route.value["target_attr"] == "gateway_id") ? route.value["target_value"] : null
      instance_id               = (route.value["target_attr"] == "instance_id") ? route.value["target_value"] : null
      local_gateway_id          = (route.value["target_attr"] == "local_gateway_id") ? route.value["target_value"] : null
      nat_gateway_id            = (route.value["target_attr"] == "nat_gateway_id") ? route.value["target_value"] : null
      network_interface_id      = (route.value["target_attr"] == "network_interface_id") ? route.value["target_value"] : null
      transit_gateway_id        = (route.value["target_attr"] == "transit_gateway_id") ? route.value["target_value"] : null
      vpc_endpoint_id           = (route.value["target_attr"] == "vpc_endpoint_id") ? route.value["target_value"] : null
      vpc_peering_connection_id = (route.value["target_attr"] == "vpc_peering_connection_id") ? route.value["target_value"] : null
    }
  }

  tags = {
    Name = %[1]q
  }
}
`, rName, destinationAttr1, destinationValue1, targetAttribute1, targetValue1, destinationAttr2, destinationValue2, targetAttribute2, targetValue2, destinationAttr3, destinationValue3, targetAttribute3, targetValue3))
}

// testAccLatestAmazonNatInstanceAMIConfig returns the configuration for a data source that
// describes the latest Amazon NAT instance AMI.
// See https://docs.aws.amazon.com/vpc/latest/userguide/VPC_NAT_Instance.html#nat-instance-ami.
// The data source is named 'amzn-ami-nat-instance'.
func testAccLatestAmazonNatInstanceAMIConfig() string {
	return `
data "aws_ami" "amzn-ami-nat-instance" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "name"
    values = ["amzn-ami-vpc-nat-*"]
  }
}
`
}

func testAccVPCRouteTableConfig_prefixListInternetGateway(rName string) string {
	return fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block = "10.1.0.0/16"

  tags = {
    Name = %[1]q
  }
}

resource "aws_internet_gateway" "test" {
  vpc_id = aws_vpc.test.id

  tags = {
    Name = %[1]q
  }
}

resource "aws_ec2_managed_prefix_list" "test" {
  address_family = "IPv4"
  max_entries    = 1
  name           = %[1]q
}

resource "aws_route_table" "test" {
  vpc_id = aws_vpc.test.id

  route {
    destination_prefix_list_id = aws_ec2_managed_prefix_list.test.id
    gateway_id                 = aws_internet_gateway.test.id
  }

  tags = {
    Name = %[1]q
  }
}
`, rName)
}
