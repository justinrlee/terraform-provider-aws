---
subcategory: "IVS (Interactive Video)"
layout: "aws"
page_title: "AWS: aws_ivs_stream_key"
description: |-
  Terraform data source for managing an AWS IVS (Interactive Video) Stream Key.
---


<!-- Please do not edit this file, it is generated. -->
# Data Source: aws_ivs_stream_key

Terraform data source for managing an AWS IVS (Interactive Video) Stream Key.

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
import { DataAwsIvsStreamKey } from "./.gen/providers/aws/data-aws-ivs-stream-key";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new DataAwsIvsStreamKey(this, "example", {
      channelArn: "arn:aws:ivs:us-west-2:326937407773:channel/0Y1lcs4U7jk5",
    });
  }
}

```

## Argument Reference

The following arguments are required:

* `channelArn` - (Required) ARN of the Channel.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `arn` - ARN of the Stream Key.
* `tags` - Map of tags assigned to the resource.
* `value` - Stream Key value.

<!-- cache-key: cdktf-0.20.8 input-07ea21d8ae54ff1a061e2ce5f7053d91f4807ee4e6b32e670b8c70e8582e5c98 -->