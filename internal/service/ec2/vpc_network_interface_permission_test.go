// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package ec2_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/YakDriver/regexache"
	sdkacctest "github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	tfec2 "github.com/hashicorp/terraform-provider-aws/internal/service/ec2"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccVPCNetworkInterfacePermission_basic(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_interface_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		CheckDestroy:             testAccCheckVPCNetworkInterfacePermissionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInterfacePermissionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccVPCNetworkInterfacePermissionExists(ctx, resourceName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrAccountID),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrNetworkInterfaceID),
					resource.TestCheckResourceAttrSet(resourceName, "permission"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{names.AttrSkipDestroy},
			},
		},
	})
}

func TestAccVPCNetworkInterfacePermission_disappears(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_interface_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		CheckDestroy:             testAccCheckVPCNetworkInterfacePermissionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInterfacePermissionConfig_basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccVPCNetworkInterfacePermissionExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkInterfacePermission(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func TestAccVPCNetworkInterfacePermission_ownerExpectError(t *testing.T) {
	ctx := acctest.Context(t)
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
		},
		ErrorCheck:               acctest.ErrorCheck(t, names.EC2ServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVPCNetworkInterfacePermissionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config:      testAccVPCNetworkInterfacePermissionConfig_accountOwner(rName),
				ExpectError: regexache.MustCompile(`OperationNotPermitted`),
			},
		},
	})
}

func testAccCheckVPCNetworkInterfacePermissionDestroy(ctx context.Context) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Client(ctx)

		for _, rs := range s.RootModule().Resources {
			if rs.Type != "aws_network_interface_permission" {
				continue
			}

			_, err := tfec2.FindNetworkInterfacePermissionByID(ctx, conn, rs.Primary.Attributes["id"])

			if tfresource.NotFound(err) {
				continue
			}

			if err != nil {
				return err
			}

			return fmt.Errorf("VPC Network Interface Permission %s still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccVPCNetworkInterfacePermissionExists(ctx context.Context, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		conn := acctest.Provider.Meta().(*conns.AWSClient).EC2Client(ctx)

		_, err := tfec2.FindNetworkInterfacePermissionByID(ctx, conn, rs.Primary.Attributes["id"])

		return err
	}
}

func testAccVPCNetworkInterfacePermissionConfig_basic(rName string) string {
	return acctest.ConfigCompose(
		acctest.ConfigAlternateAccountProvider(),
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block           = "172.16.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_id            = aws_vpc.test.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
}

data "aws_caller_identity" "test" {
  provider = "awsalternate"
}

resource "aws_network_interface_permission" "test" {
    network_interface_id = aws_network_interface.test.id
    account_id           = data.aws_caller_identity.test.account_id
    permission           = "INSTANCE-ATTACH"
}
`, rName))
}

func testAccVPCNetworkInterfacePermissionConfig_accountOwner(rName string) string {
	return acctest.ConfigCompose(acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block           = "172.16.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_id            = aws_vpc.test.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
}

data "aws_caller_identity" "test" {}

resource "aws_network_interface_permission" "test" {
    network_interface_id = aws_network_interface.test.id
    account_id           = data.aws_caller_identity.test.account_id
    permission           = "INSTANCE-ATTACH"
}
`, rName))
}

func TestAccVPCNetworkInterfacePermission_basic_l(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_interface_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			// acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck: acctest.ErrorCheck(t, names.EC2ServiceID),
		// ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVPCNetworkInterfacePermissionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInterfacePermissionConfig_basic_l(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccVPCNetworkInterfacePermissionExists(ctx, resourceName),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrAccountID),
					resource.TestCheckResourceAttrSet(resourceName, names.AttrNetworkInterfaceID),
					resource.TestCheckResourceAttrSet(resourceName, "permission"),
				),
			},
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{names.AttrSkipDestroy},
			},
		},
	})
}

func TestAccVPCNetworkInterfacePermission_disappears_l(t *testing.T) {
	ctx := acctest.Context(t)
	resourceName := "aws_network_interface_permission.test"
	rName := sdkacctest.RandomWithPrefix(acctest.ResourcePrefix)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheck(ctx, t)
			// acctest.PreCheckAlternateAccount(t)
		},
		ErrorCheck: acctest.ErrorCheck(t, names.EC2ServiceID),
		// ProtoV5ProviderFactories: acctest.ProtoV5FactoriesAlternate(ctx, t),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		CheckDestroy:             testAccCheckVPCNetworkInterfacePermissionDestroy(ctx),
		Steps: []resource.TestStep{
			{
				Config: testAccVPCNetworkInterfacePermissionConfig_basic_l(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccVPCNetworkInterfacePermissionExists(ctx, resourceName),
					acctest.CheckResourceDisappears(ctx, acctest.Provider, tfec2.ResourceNetworkInterfacePermission(), resourceName),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}

func testAccVPCNetworkInterfacePermissionConfig_basic_l(rName string) string {
	return acctest.ConfigCompose(
		// acctest.ConfigAlternateAccountProvider(),
		acctest.ConfigAvailableAZsNoOptIn(),
		fmt.Sprintf(`
resource "aws_vpc" "test" {
  cidr_block           = "172.16.0.0/16"
  enable_dns_hostnames = true

  tags = {
    Name = %[1]q
  }
}

resource "aws_subnet" "test" {
  vpc_id            = aws_vpc.test.id
  cidr_block        = "172.16.10.0/24"
  availability_zone = data.aws_availability_zones.available.names[0]

  tags = {
    Name = %[1]q
  }
}

resource "aws_network_interface" "test" {
  subnet_id = aws_subnet.test.id
}

// data "aws_caller_identity" "test" {
//   provider = "awsalternate"
// }

resource "aws_network_interface_permission" "test" {
    network_interface_id = aws_network_interface.test.id
    account_id           = "12345"
    // account_id           = data.aws_caller_identity.test.account_id
    permission           = "INSTANCE-ATTACH"
}
`, rName))
}
