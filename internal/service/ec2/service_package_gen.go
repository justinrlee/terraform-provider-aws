// Code generated by internal/generate/servicepackage/main.go; DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/types"
	"github.com/hashicorp/terraform-provider-aws/names"
)

type servicePackage struct{}

func (p *servicePackage) FrameworkDataSources(ctx context.Context) []*types.ServicePackageFrameworkDataSource {
	return []*types.ServicePackageFrameworkDataSource{
		{
			Factory: newCapacityBlockOfferingDataSource,
			Name:    "Capacity Block Offering",
		},
		{
			Factory: newDataSourceSpotDataFeedSubscription,
			Name:    "Spot Data Feed Subscription Data Source",
		},
		{
			Factory: newSecurityGroupRuleDataSource,
			Name:    "Security Group Rule",
		},
		{
			Factory: newSecurityGroupRulesDataSource,
			Name:    "Security Group Rules",
		},
	}
}

func (p *servicePackage) FrameworkResources(ctx context.Context) []*types.ServicePackageFrameworkResource {
	return []*types.ServicePackageFrameworkResource{
		{
			Factory: newCapacityBlockReservationResource,
			Name:    "Capacity Block Reservation",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory: newEBSFastSnapshotRestoreResource,
			Name:    "EBS Fast Snapshot Restore",
		},
		{
			Factory: newEIPDomainNameResource,
			Name:    "EIP Domain Name",
		},
		{
			Factory: newInstanceConnectEndpointResource,
			Name:    "Instance Connect Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory: newInstanceMetadataDefaultsResource,
			Name:    "Instance Metadata Defaults",
		},
		{
			Factory: newNetworkInterfacePermission,
			Name:    "Network Interface Permission",
		},
		{
			Factory: newResourceSecurityGroupVPCAssociation,
			Name:    "Security Group VPC Association",
		},
		{
			Factory: newSecurityGroupEgressRuleResource,
			Name:    "Security Group Egress Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory: newSecurityGroupIngressRuleResource,
			Name:    "Security Group Ingress Rule",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory: newTransitGatewayDefaultRouteTableAssociationResource,
			Name:    "Transit Gateway Default Route Table Association",
		},
		{
			Factory: newTransitGatewayDefaultRouteTablePropagationResource,
			Name:    "Transit Gateway Default Route Table Propagation",
		},
		{
			Factory: newVPCBlockPublicAccessExclusionResource,
			Name:    "VPC Block Public Access Exclusion",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory: newVPCBlockPublicAccessOptionsResource,
			Name:    "VPC Block Public Access Options",
		},
		{
			Factory: newVPCEndpointPrivateDNSResource,
			Name:    "VPC Endpoint Private DNS",
		},
		{
			Factory: newVPCEndpointServicePrivateDNSVerificationResource,
			Name:    "VPC Endpoint Service Private DNS Verification",
		},
	}
}

func (p *servicePackage) SDKDataSources(ctx context.Context) []*types.ServicePackageSDKDataSource {
	return []*types.ServicePackageSDKDataSource{
		{
			Factory:  dataSourceAMI,
			TypeName: "aws_ami",
			Name:     "AMI",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceAMIIDs,
			TypeName: "aws_ami_ids",
			Name:     "AMI IDs",
		},
		{
			Factory:  dataSourceAvailabilityZone,
			TypeName: "aws_availability_zone",
			Name:     "Availability Zone",
		},
		{
			Factory:  dataSourceAvailabilityZones,
			TypeName: "aws_availability_zones",
			Name:     "Availability Zones",
		},
		{
			Factory:  dataSourceCustomerGateway,
			TypeName: "aws_customer_gateway",
			Name:     "Customer Gateway",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEBSDefaultKMSKey,
			TypeName: "aws_ebs_default_kms_key",
			Name:     "EBS Default KMS Key",
		},
		{
			Factory:  dataSourceEBSEncryptionByDefault,
			TypeName: "aws_ebs_encryption_by_default",
			Name:     "EBS Encryption By Default",
		},
		{
			Factory:  dataSourceEBSSnapshot,
			TypeName: "aws_ebs_snapshot",
			Name:     "EBS Snapshot",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEBSSnapshotIDs,
			TypeName: "aws_ebs_snapshot_ids",
			Name:     "EBS Snapshot IDs",
		},
		{
			Factory:  dataSourceEBSVolume,
			TypeName: "aws_ebs_volume",
			Name:     "EBS Volume",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEBSVolumes,
			TypeName: "aws_ebs_volumes",
			Name:     "EBS Volumes",
		},
		{
			Factory:  dataSourceClientVPNEndpoint,
			TypeName: "aws_ec2_client_vpn_endpoint",
			Name:     "Client VPN Endpoint",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceCoIPPool,
			TypeName: "aws_ec2_coip_pool",
			Name:     "COIP Pool",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceCoIPPools,
			TypeName: "aws_ec2_coip_pools",
			Name:     "COIP Pools",
		},
		{
			Factory:  dataSourceHost,
			TypeName: "aws_ec2_host",
			Name:     "Host",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceInstanceType,
			TypeName: "aws_ec2_instance_type",
			Name:     "Instance Type",
		},
		{
			Factory:  dataSourceInstanceTypeOffering,
			TypeName: "aws_ec2_instance_type_offering",
			Name:     "Instance Type Offering",
		},
		{
			Factory:  dataSourceInstanceTypeOfferings,
			TypeName: "aws_ec2_instance_type_offerings",
			Name:     "Instance Type Offering",
		},
		{
			Factory:  dataSourceInstanceTypes,
			TypeName: "aws_ec2_instance_types",
			Name:     "Instance Types",
		},
		{
			Factory:  dataSourceLocalGateway,
			TypeName: "aws_ec2_local_gateway",
			Name:     "Local Gateway",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceLocalGatewayRouteTable,
			TypeName: "aws_ec2_local_gateway_route_table",
			Name:     "Local Gateway Route Table",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceLocalGatewayRouteTables,
			TypeName: "aws_ec2_local_gateway_route_tables",
			Name:     "Local Gateway Route Table",
		},
		{
			Factory:  dataSourceLocalGatewayVirtualInterface,
			TypeName: "aws_ec2_local_gateway_virtual_interface",
			Name:     "Local Gateway Virtual Interface",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceLocalGatewayVirtualInterfaceGroup,
			TypeName: "aws_ec2_local_gateway_virtual_interface_group",
			Name:     "Local Gateway Virtual Interface Group",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceLocalGatewayVirtualInterfaceGroups,
			TypeName: "aws_ec2_local_gateway_virtual_interface_groups",
			Name:     "Local Gateway Virtual Interface Groups",
		},
		{
			Factory:  dataSourceLocalGateways,
			TypeName: "aws_ec2_local_gateways",
			Name:     "Local Gateways",
		},
		{
			Factory:  dataSourceManagedPrefixList,
			TypeName: "aws_ec2_managed_prefix_list",
			Name:     "Managed Prefix List",
		},
		{
			Factory:  dataSourceManagedPrefixLists,
			TypeName: "aws_ec2_managed_prefix_lists",
			Name:     "Managed Prefix Lists",
		},
		{
			Factory:  dataSourceNetworkInsightsAnalysis,
			TypeName: "aws_ec2_network_insights_analysis",
			Name:     "Network Insights Analysis",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceNetworkInsightsPath,
			TypeName: "aws_ec2_network_insights_path",
			Name:     "Network Insights Path",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourcePublicIPv4Pool,
			TypeName: "aws_ec2_public_ipv4_pool",
			Name:     "Public IPv4 Pool",
		},
		{
			Factory:  dataSourcePublicIPv4Pools,
			TypeName: "aws_ec2_public_ipv4_pools",
			Name:     "Public IPv4 Pools",
		},
		{
			Factory:  dataSourceSerialConsoleAccess,
			TypeName: "aws_ec2_serial_console_access",
			Name:     "Serial Console Access",
		},
		{
			Factory:  dataSourceSpotPrice,
			TypeName: "aws_ec2_spot_price",
			Name:     "Spot Price",
		},
		{
			Factory:  dataSourceTransitGateway,
			TypeName: "aws_ec2_transit_gateway",
			Name:     "Transit Gateway",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayAttachment,
			TypeName: "aws_ec2_transit_gateway_attachment",
			Name:     "Transit Gateway Attachment",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayAttachments,
			TypeName: "aws_ec2_transit_gateway_attachments",
			Name:     "Transit Gateway Attachments",
		},
		{
			Factory:  dataSourceTransitGatewayConnect,
			TypeName: "aws_ec2_transit_gateway_connect",
			Name:     "Transit Gateway Connect",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayConnectPeer,
			TypeName: "aws_ec2_transit_gateway_connect_peer",
			Name:     "Transit Gateway Connect Peer",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayDxGatewayAttachment,
			TypeName: "aws_ec2_transit_gateway_dx_gateway_attachment",
			Name:     "Transit Gateway Direct Connect Gateway Attachment",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayMulticastDomain,
			TypeName: "aws_ec2_transit_gateway_multicast_domain",
			Name:     "Transit Gateway Multicast Domain",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayPeeringAttachment,
			TypeName: "aws_ec2_transit_gateway_peering_attachment",
			Name:     "Transit Gateway Peering Attachment",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayPeeringAttachments,
			TypeName: "aws_ec2_transit_gateway_peering_attachments",
			Name:     "Transit Gateway Peering Attachments",
		},
		{
			Factory:  dataSourceTransitGatewayRouteTable,
			TypeName: "aws_ec2_transit_gateway_route_table",
			Name:     "Transit Gateway Route Table",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayRouteTableAssociations,
			TypeName: "aws_ec2_transit_gateway_route_table_associations",
			Name:     "Transit Gateway Route Table Associations",
		},
		{
			Factory:  dataSourceTransitGatewayRouteTablePropagations,
			TypeName: "aws_ec2_transit_gateway_route_table_propagations",
			Name:     "Transit Gateway Route Table Propagations",
		},
		{
			Factory:  dataSourceTransitGatewayRouteTableRoutes,
			TypeName: "aws_ec2_transit_gateway_route_table_routes",
			Name:     "Transit Gateway Route Table Routes",
		},
		{
			Factory:  dataSourceTransitGatewayRouteTables,
			TypeName: "aws_ec2_transit_gateway_route_tables",
			Name:     "Transit Gateway Route Tables",
		},
		{
			Factory:  dataSourceTransitGatewayVPCAttachment,
			TypeName: "aws_ec2_transit_gateway_vpc_attachment",
			Name:     "Transit Gateway VPC Attachment",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceTransitGatewayVPCAttachments,
			TypeName: "aws_ec2_transit_gateway_vpc_attachments",
			Name:     "Transit Gateway VPC Attachments",
		},
		{
			Factory:  dataSourceTransitGatewayVPNAttachment,
			TypeName: "aws_ec2_transit_gateway_vpn_attachment",
			Name:     "Transit Gateway VPN Attachment",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEIP,
			TypeName: "aws_eip",
			Name:     "EIP",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceEIPs,
			TypeName: "aws_eips",
			Name:     "EIPs",
		},
		{
			Factory:  dataSourceInstance,
			TypeName: "aws_instance",
			Name:     "Instance",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceInstances,
			TypeName: "aws_instances",
			Name:     "Instances",
		},
		{
			Factory:  dataSourceInternetGateway,
			TypeName: "aws_internet_gateway",
			Name:     "Internet Gateway",
		},
		{
			Factory:  dataSourceKeyPair,
			TypeName: "aws_key_pair",
			Name:     "Key Pair",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceLaunchTemplate,
			TypeName: "aws_launch_template",
			Name:     "Launch Template",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceNATGateway,
			TypeName: "aws_nat_gateway",
			Name:     "NAT Gateway",
		},
		{
			Factory:  dataSourceNATGateways,
			TypeName: "aws_nat_gateways",
			Name:     "NAT Gateways",
		},
		{
			Factory:  dataSourceNetworkACLs,
			TypeName: "aws_network_acls",
			Name:     "Network ACLs",
		},
		{
			Factory:  dataSourceNetworkInterface,
			TypeName: "aws_network_interface",
			Name:     "Network Interface",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceNetworkInterfaces,
			TypeName: "aws_network_interfaces",
			Name:     "Network Interfaces",
		},
		{
			Factory:  dataSourcePrefixList,
			TypeName: "aws_prefix_list",
			Name:     "Prefix List",
		},
		{
			Factory:  dataSourceRoute,
			TypeName: "aws_route",
			Name:     "Route",
		},
		{
			Factory:  dataSourceRouteTable,
			TypeName: "aws_route_table",
			Name:     "Route Table",
		},
		{
			Factory:  dataSourceRouteTables,
			TypeName: "aws_route_tables",
			Name:     "Route Tables",
		},
		{
			Factory:  dataSourceSecurityGroup,
			TypeName: "aws_security_group",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceSecurityGroups,
			TypeName: "aws_security_groups",
			Name:     "Security Groups",
		},
		{
			Factory:  dataSourceSubnet,
			TypeName: "aws_subnet",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceSubnets,
			TypeName: "aws_subnets",
		},
		{
			Factory:  dataSourceVPC,
			TypeName: "aws_vpc",
			Name:     "VPC",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceVPCDHCPOptions,
			TypeName: "aws_vpc_dhcp_options",
		},
		{
			Factory:  dataSourceVPCEndpoint,
			TypeName: "aws_vpc_endpoint",
			Name:     "Endpoint",
		},
		{
			Factory:  dataSourceVPCEndpointService,
			TypeName: "aws_vpc_endpoint_service",
			Name:     "Endpoint Service",
		},
		{
			Factory:  dataSourceIPAMPool,
			TypeName: "aws_vpc_ipam_pool",
			Name:     "IPAM Pool",
			Tags:     &types.ServicePackageResourceTags{},
		},
		{
			Factory:  dataSourceIPAMPoolCIDRs,
			TypeName: "aws_vpc_ipam_pool_cidrs",
			Name:     "IPAM Pool CIDRs",
		},
		{
			Factory:  dataSourceIPAMPools,
			TypeName: "aws_vpc_ipam_pools",
			Name:     "IPAM Pools",
		},
		{
			Factory:  dataSourceIPAMPreviewNextCIDR,
			TypeName: "aws_vpc_ipam_preview_next_cidr",
			Name:     "IPAM Preview Next CIDR",
		},
		{
			Factory:  dataSourceVPCPeeringConnection,
			TypeName: "aws_vpc_peering_connection",
			Name:     "VPC Peering Connection",
		},
		{
			Factory:  dataSourceVPCPeeringConnections,
			TypeName: "aws_vpc_peering_connections",
			Name:     "VPC Peering Connections",
		},
		{
			Factory:  dataSourceVPCs,
			TypeName: "aws_vpcs",
			Name:     "VPCs",
		},
		{
			Factory:  dataSourceVPNGateway,
			TypeName: "aws_vpn_gateway",
			Name:     "VPN Gateway",
			Tags:     &types.ServicePackageResourceTags{},
		},
	}
}

func (p *servicePackage) SDKResources(ctx context.Context) []*types.ServicePackageSDKResource {
	return []*types.ServicePackageSDKResource{
		{
			Factory:  resourceAMI,
			TypeName: "aws_ami",
			Name:     "AMI",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceAMICopy,
			TypeName: "aws_ami_copy",
			Name:     "AMI",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceAMIFromInstance,
			TypeName: "aws_ami_from_instance",
			Name:     "AMI",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceAMILaunchPermission,
			TypeName: "aws_ami_launch_permission",
			Name:     "AMI Launch Permission",
		},
		{
			Factory:  resourceCustomerGateway,
			TypeName: "aws_customer_gateway",
			Name:     "Customer Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultNetworkACL,
			TypeName: "aws_default_network_acl",
			Name:     "Network ACL",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultRouteTable,
			TypeName: "aws_default_route_table",
			Name:     "Route Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultSecurityGroup,
			TypeName: "aws_default_security_group",
			Name:     "Security Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultSubnet,
			TypeName: "aws_default_subnet",
			Name:     "Subnet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultVPC,
			TypeName: "aws_default_vpc",
			Name:     "Default VPC",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceDefaultVPCDHCPOptions,
			TypeName: "aws_default_vpc_dhcp_options",
			Name:     "DHCP Options",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEBSDefaultKMSKey,
			TypeName: "aws_ebs_default_kms_key",
			Name:     "EBS Default KMS Key",
		},
		{
			Factory:  resourceEBSEncryptionByDefault,
			TypeName: "aws_ebs_encryption_by_default",
			Name:     "EBS Encryption By Default",
		},
		{
			Factory:  resourceEBSSnapshot,
			TypeName: "aws_ebs_snapshot",
			Name:     "EBS Snapshot",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEBSSnapshotBlockPublicAccess,
			TypeName: "aws_ebs_snapshot_block_public_access",
			Name:     "EBS Snapshot Block Public Access",
		},
		{
			Factory:  resourceEBSSnapshotCopy,
			TypeName: "aws_ebs_snapshot_copy",
			Name:     "EBS Snapshot Copy",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEBSSnapshotImport,
			TypeName: "aws_ebs_snapshot_import",
			Name:     "EBS Snapshot Import",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEBSVolume,
			TypeName: "aws_ebs_volume",
			Name:     "EBS Volume",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceAvailabilityZoneGroup,
			TypeName: "aws_ec2_availability_zone_group",
			Name:     "Availability Zone Group",
		},
		{
			Factory:  resourceCapacityReservation,
			TypeName: "aws_ec2_capacity_reservation",
			Name:     "Capacity Reservation",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceCarrierGateway,
			TypeName: "aws_ec2_carrier_gateway",
			Name:     "Carrier Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceClientVPNAuthorizationRule,
			TypeName: "aws_ec2_client_vpn_authorization_rule",
			Name:     "Client VPN Authorization Rule",
		},
		{
			Factory:  resourceClientVPNEndpoint,
			TypeName: "aws_ec2_client_vpn_endpoint",
			Name:     "Client VPN Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceClientVPNNetworkAssociation,
			TypeName: "aws_ec2_client_vpn_network_association",
			Name:     "Client VPN Network Association",
		},
		{
			Factory:  resourceClientVPNRoute,
			TypeName: "aws_ec2_client_vpn_route",
			Name:     "Client VPN Route",
		},
		{
			Factory:  resourceFleet,
			TypeName: "aws_ec2_fleet",
			Name:     "Fleet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceHost,
			TypeName: "aws_ec2_host",
			Name:     "Host",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceImageBlockPublicAccess,
			TypeName: "aws_ec2_image_block_public_access",
			Name:     "Image Block Public Access",
		},
		{
			Factory:  resourceInstanceState,
			TypeName: "aws_ec2_instance_state",
			Name:     "Instance State",
		},
		{
			Factory:  resourceLocalGatewayRoute,
			TypeName: "aws_ec2_local_gateway_route",
		},
		{
			Factory:  resourceLocalGatewayRouteTableVPCAssociation,
			TypeName: "aws_ec2_local_gateway_route_table_vpc_association",
			Name:     "Local Gateway Route Table VPC Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceManagedPrefixList,
			TypeName: "aws_ec2_managed_prefix_list",
			Name:     "Managed Prefix List",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceManagedPrefixListEntry,
			TypeName: "aws_ec2_managed_prefix_list_entry",
			Name:     "Managed Prefix List Entry",
		},
		{
			Factory:  resourceNetworkInsightsAnalysis,
			TypeName: "aws_ec2_network_insights_analysis",
			Name:     "Network Insights Analysis",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceNetworkInsightsPath,
			TypeName: "aws_ec2_network_insights_path",
			Name:     "Network Insights Path",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceSerialConsoleAccess,
			TypeName: "aws_ec2_serial_console_access",
			Name:     "Serial Console Access",
		},
		{
			Factory:  resourceSubnetCIDRReservation,
			TypeName: "aws_ec2_subnet_cidr_reservation",
			Name:     "Subnet CIDR Reservation",
		},
		{
			Factory:  resourceTag,
			TypeName: "aws_ec2_tag",
			Name:     "EC2 Resource Tag",
		},
		{
			Factory:  resourceTrafficMirrorFilter,
			TypeName: "aws_ec2_traffic_mirror_filter",
			Name:     "Traffic Mirror Filter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTrafficMirrorFilterRule,
			TypeName: "aws_ec2_traffic_mirror_filter_rule",
			Name:     "Traffic Mirror Filter Rule",
		},
		{
			Factory:  resourceTrafficMirrorSession,
			TypeName: "aws_ec2_traffic_mirror_session",
			Name:     "Traffic Mirror Session",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTrafficMirrorTarget,
			TypeName: "aws_ec2_traffic_mirror_target",
			Name:     "Traffic Mirror Target",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGateway,
			TypeName: "aws_ec2_transit_gateway",
			Name:     "Transit Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayConnect,
			TypeName: "aws_ec2_transit_gateway_connect",
			Name:     "Transit Gateway Connect",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayConnectPeer,
			TypeName: "aws_ec2_transit_gateway_connect_peer",
			Name:     "Transit Gateway Connect Peer",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayMulticastDomain,
			TypeName: "aws_ec2_transit_gateway_multicast_domain",
			Name:     "Transit Gateway Multicast Domain",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayMulticastDomainAssociation,
			TypeName: "aws_ec2_transit_gateway_multicast_domain_association",
			Name:     "Transit Gateway Multicast Domain Association",
		},
		{
			Factory:  resourceTransitGatewayMulticastGroupMember,
			TypeName: "aws_ec2_transit_gateway_multicast_group_member",
			Name:     "Transit Gateway Multicast Group Member",
		},
		{
			Factory:  resourceTransitGatewayMulticastGroupSource,
			TypeName: "aws_ec2_transit_gateway_multicast_group_source",
			Name:     "Transit Gateway Multicast Group Source",
		},
		{
			Factory:  resourceTransitGatewayPeeringAttachment,
			TypeName: "aws_ec2_transit_gateway_peering_attachment",
			Name:     "Transit Gateway Peering Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayPeeringAttachmentAccepter,
			TypeName: "aws_ec2_transit_gateway_peering_attachment_accepter",
			Name:     "Transit Gateway Peering Attachment Accepter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayPolicyTable,
			TypeName: "aws_ec2_transit_gateway_policy_table",
			Name:     "Transit Gateway Policy Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayPolicyTableAssociation,
			TypeName: "aws_ec2_transit_gateway_policy_table_association",
			Name:     "Transit Gateway Policy Table Association",
		},
		{
			Factory:  resourceTransitGatewayPrefixListReference,
			TypeName: "aws_ec2_transit_gateway_prefix_list_reference",
			Name:     "Transit Gateway Prefix List Reference",
		},
		{
			Factory:  resourceTransitGatewayRoute,
			TypeName: "aws_ec2_transit_gateway_route",
			Name:     "Transit Gateway Route",
		},
		{
			Factory:  resourceTransitGatewayRouteTable,
			TypeName: "aws_ec2_transit_gateway_route_table",
			Name:     "Transit Gateway Route Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayRouteTableAssociation,
			TypeName: "aws_ec2_transit_gateway_route_table_association",
			Name:     "Transit Gateway Route Table Association",
		},
		{
			Factory:  resourceTransitGatewayRouteTablePropagation,
			TypeName: "aws_ec2_transit_gateway_route_table_propagation",
			Name:     "Transit Gateway Route Table Propagation",
		},
		{
			Factory:  resourceTransitGatewayVPCAttachment,
			TypeName: "aws_ec2_transit_gateway_vpc_attachment",
			Name:     "Transit Gateway VPC Attachment",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceTransitGatewayVPCAttachmentAccepter,
			TypeName: "aws_ec2_transit_gateway_vpc_attachment_accepter",
			Name:     "Transit Gateway VPC Attachment Accepter",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEgressOnlyInternetGateway,
			TypeName: "aws_egress_only_internet_gateway",
			Name:     "Egress-Only Internet Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEIP,
			TypeName: "aws_eip",
			Name:     "EIP",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceEIPAssociation,
			TypeName: "aws_eip_association",
			Name:     "EIP Association",
		},
		{
			Factory:  resourceFlowLog,
			TypeName: "aws_flow_log",
			Name:     "Flow Log",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceInstance,
			TypeName: "aws_instance",
			Name:     "Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceInternetGateway,
			TypeName: "aws_internet_gateway",
			Name:     "Internet Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceInternetGatewayAttachment,
			TypeName: "aws_internet_gateway_attachment",
			Name:     "Internet Gateway Attachment",
		},
		{
			Factory:  resourceKeyPair,
			TypeName: "aws_key_pair",
			Name:     "Key Pair",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "key_pair_id",
			},
		},
		{
			Factory:  resourceLaunchTemplate,
			TypeName: "aws_launch_template",
			Name:     "Launch Template",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceMainRouteTableAssociation,
			TypeName: "aws_main_route_table_association",
			Name:     "Main Route Table Association",
		},
		{
			Factory:  resourceNATGateway,
			TypeName: "aws_nat_gateway",
			Name:     "NAT Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceNetworkACL,
			TypeName: "aws_network_acl",
			Name:     "Network ACL",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceNetworkACLAssociation,
			TypeName: "aws_network_acl_association",
			Name:     "Network ACL Association",
		},
		{
			Factory:  resourceNetworkACLRule,
			TypeName: "aws_network_acl_rule",
			Name:     "Network ACL Rule",
		},
		{
			Factory:  resourceNetworkInterface,
			TypeName: "aws_network_interface",
			Name:     "Network Interface",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceNetworkInterfaceAttachment,
			TypeName: "aws_network_interface_attachment",
			Name:     "Network Interface Attachment",
		},
		{
			Factory:  resourceNetworkInterfaceSGAttachment,
			TypeName: "aws_network_interface_sg_attachment",
			Name:     "Network Interface SG Attachement",
		},
		{
			Factory:  resourcePlacementGroup,
			TypeName: "aws_placement_group",
			Name:     "Placement Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: "placement_group_id",
			},
		},
		{
			Factory:  resourceRoute,
			TypeName: "aws_route",
			Name:     "Route",
		},
		{
			Factory:  resourceRouteTable,
			TypeName: "aws_route_table",
			Name:     "Route Table",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceRouteTableAssociation,
			TypeName: "aws_route_table_association",
			Name:     "Route Table Association",
		},
		{
			Factory:  resourceSecurityGroup,
			TypeName: "aws_security_group",
			Name:     "Security Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceSecurityGroupRule,
			TypeName: "aws_security_group_rule",
			Name:     "Security Group Rule",
		},
		{
			Factory:  resourceSnapshotCreateVolumePermission,
			TypeName: "aws_snapshot_create_volume_permission",
			Name:     "EBS Snapshot CreateVolume Permission",
		},
		{
			Factory:  resourceSpotDataFeedSubscription,
			TypeName: "aws_spot_datafeed_subscription",
			Name:     "Spot Datafeed Subscription",
		},
		{
			Factory:  resourceSpotFleetRequest,
			TypeName: "aws_spot_fleet_request",
			Name:     "Spot Fleet Request",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceSpotInstanceRequest,
			TypeName: "aws_spot_instance_request",
			Name:     "Spot Instance Request",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceSubnet,
			TypeName: "aws_subnet",
			Name:     "Subnet",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVerifiedAccessEndpoint,
			TypeName: "aws_verifiedaccess_endpoint",
			Name:     "Verified Access Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVerifiedAccessGroup,
			TypeName: "aws_verifiedaccess_group",
			Name:     "Verified Access Group",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVerifiedAccessInstance,
			TypeName: "aws_verifiedaccess_instance",
			Name:     "Verified Access Instance",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVerifiedAccessInstanceLoggingConfiguration,
			TypeName: "aws_verifiedaccess_instance_logging_configuration",
			Name:     "Verified Access Instance Logging Configuration",
		},
		{
			Factory:  resourceVerifiedAccessInstanceTrustProviderAttachment,
			TypeName: "aws_verifiedaccess_instance_trust_provider_attachment",
			Name:     "Verified Access Instance Trust Provider Attachment",
		},
		{
			Factory:  resourceVerifiedAccessTrustProvider,
			TypeName: "aws_verifiedaccess_trust_provider",
			Name:     "Verified Access Trust Provider",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVolumeAttachment,
			TypeName: "aws_volume_attachment",
			Name:     "EBS Volume Attachment",
		},
		{
			Factory:  resourceVPC,
			TypeName: "aws_vpc",
			Name:     "VPC",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCDHCPOptions,
			TypeName: "aws_vpc_dhcp_options",
			Name:     "DHCP Options",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCDHCPOptionsAssociation,
			TypeName: "aws_vpc_dhcp_options_association",
			Name:     "VPC DHCP Options Association",
		},
		{
			Factory:  resourceVPCEndpoint,
			TypeName: "aws_vpc_endpoint",
			Name:     "VPC Endpoint",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCEndpointConnectionAccepter,
			TypeName: "aws_vpc_endpoint_connection_accepter",
			Name:     "VPC Endpoint Connection Accepter",
		},
		{
			Factory:  resourceVPCEndpointConnectionNotification,
			TypeName: "aws_vpc_endpoint_connection_notification",
			Name:     "VPC Endpoint Connection Notification",
		},
		{
			Factory:  resourceVPCEndpointPolicy,
			TypeName: "aws_vpc_endpoint_policy",
		},
		{
			Factory:  resourceVPCEndpointRouteTableAssociation,
			TypeName: "aws_vpc_endpoint_route_table_association",
			Name:     "VPC Endpoint Route Table Association",
		},
		{
			Factory:  resourceVPCEndpointSecurityGroupAssociation,
			TypeName: "aws_vpc_endpoint_security_group_association",
			Name:     "VPC Endpoint Security Group Association",
		},
		{
			Factory:  resourceVPCEndpointService,
			TypeName: "aws_vpc_endpoint_service",
			Name:     "VPC Endpoint Service",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCEndpointServiceAllowedPrincipal,
			TypeName: "aws_vpc_endpoint_service_allowed_principal",
			Name:     "Endpoint Service Allowed Principal",
		},
		{
			Factory:  resourceVPCEndpointSubnetAssociation,
			TypeName: "aws_vpc_endpoint_subnet_association",
			Name:     "VPC Endpoint Subnet Association",
		},
		{
			Factory:  resourceIPAM,
			TypeName: "aws_vpc_ipam",
			Name:     "IPAM",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceIPAMOrganizationAdminAccount,
			TypeName: "aws_vpc_ipam_organization_admin_account",
			Name:     "IPAM Organization Admin Account",
		},
		{
			Factory:  resourceIPAMPool,
			TypeName: "aws_vpc_ipam_pool",
			Name:     "IPAM Pool",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceIPAMPoolCIDR,
			TypeName: "aws_vpc_ipam_pool_cidr",
			Name:     "IPAM Pool CIDR",
		},
		{
			Factory:  resourceIPAMPoolCIDRAllocation,
			TypeName: "aws_vpc_ipam_pool_cidr_allocation",
			Name:     "IPAM Pool CIDR Allocation",
		},
		{
			Factory:  resourceIPAMPreviewNextCIDR,
			TypeName: "aws_vpc_ipam_preview_next_cidr",
			Name:     "IPAM Preview Next CIDR",
		},
		{
			Factory:  resourceIPAMResourceDiscovery,
			TypeName: "aws_vpc_ipam_resource_discovery",
			Name:     "IPAM Resource Discovery",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceIPAMResourceDiscoveryAssociation,
			TypeName: "aws_vpc_ipam_resource_discovery_association",
			Name:     "IPAM Resource Discovery Association",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceIPAMScope,
			TypeName: "aws_vpc_ipam_scope",
			Name:     "IPAM Scope",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCIPv4CIDRBlockAssociation,
			TypeName: "aws_vpc_ipv4_cidr_block_association",
			Name:     "VPC IPV4 CIDR Block Association",
		},
		{
			Factory:  resourceVPCIPv6CIDRBlockAssociation,
			TypeName: "aws_vpc_ipv6_cidr_block_association",
			Name:     "VPC IPV6 CIDR Block Association",
		},
		{
			Factory:  resourceNetworkPerformanceMetricSubscription,
			TypeName: "aws_vpc_network_performance_metric_subscription",
			Name:     "VPC Network Performance Metric Subscription",
		},
		{
			Factory:  resourceVPCPeeringConnection,
			TypeName: "aws_vpc_peering_connection",
			Name:     "VPC Peering Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCPeeringConnectionAccepter,
			TypeName: "aws_vpc_peering_connection_accepter",
			Name:     "VPC Peering Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPCPeeringConnectionOptions,
			TypeName: "aws_vpc_peering_connection_options",
			Name:     "VPC Peering Connection Options",
		},
		{
			Factory:  resourceVPNConnection,
			TypeName: "aws_vpn_connection",
			Name:     "VPN Connection",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPNConnectionRoute,
			TypeName: "aws_vpn_connection_route",
			Name:     "VPN Connection Route",
		},
		{
			Factory:  resourceVPNGateway,
			TypeName: "aws_vpn_gateway",
			Name:     "VPN Gateway",
			Tags: &types.ServicePackageResourceTags{
				IdentifierAttribute: names.AttrID,
			},
		},
		{
			Factory:  resourceVPNGatewayAttachment,
			TypeName: "aws_vpn_gateway_attachment",
			Name:     "VPN Gateway Attachment",
		},
		{
			Factory:  resourceVPNGatewayRoutePropagation,
			TypeName: "aws_vpn_gateway_route_propagation",
			Name:     "VPN Gateway Route Propagation",
		},
	}
}

func (p *servicePackage) ServicePackageName() string {
	return names.EC2
}

func ServicePackage(ctx context.Context) conns.ServicePackage {
	return &servicePackage{}
}
