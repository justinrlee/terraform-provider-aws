---
subcategory: "QuickSight"
layout: "aws"
page_title: "AWS: aws_quicksight_group"
description: |-
  Use this data source to fetch information about a QuickSight Group.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_quicksight_group

This data source can be used to fetch information about a specific
QuickSight group. By using this data source, you can reference QuickSight group
properties without having to hard code ARNs or unique IDs as input.

## Example Usage

### Basic Usage

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { DataAwsQuicksightGroup } from "./.gen/providers/aws/data-aws-quicksight-group";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsQuicksightGroup(this, "example", {
      groupName: "example",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `groupName` - (Required) The name of the group that you want to match.

The following arguments are optional:

* `awsAccountId` - (Optional) AWS account ID.
* `namespace` - (Optional) QuickSight namespace. Defaults to `default`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) for the group.
* `description` - The group description.
* `principalId` - The principal ID of the group.

<!-- cache-key: cdktf-0.20.8 input-2c348640ef79405d8a0bd0ed47ae13ff12c796f3ad5a7922552ef44e21969deb -->