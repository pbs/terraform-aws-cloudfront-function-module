# PBS TF CloudFront Function

## Installation

### Using the Repo Source

Use this URL for the source of the module. See the usage examples below for more details.

```hcl
github.com/pbs/terraform-aws-cloudfront-function-module?ref=x.y.z
```

### Alternative Installation Methods

More information can be found on these install methods and more in [the documentation here](./docs/general/install).

## Usage

A CloudFront function that integrates with one or more CloudFront behaviors.

Integrate this module like so:

```hcl
module "cloudfront-function" {
  source = "github.com/pbs/terraform-aws-cloudfront-function-module?ref=x.y.z"

  code = file("function.js")

  # Optional Parameters
}
```

## Adding This Version of the Module

If this repo is added as a subtree, then the version of the module should be close to the version shown here:

`x.y.z`

Note, however that subtrees can be altered as desired within repositories.

Further documentation on usage can be found [here](./docs).

Below is automatically generated documentation on this Terraform module using [terraform-docs][terraform-docs]

---

[terraform-docs]: https://github.com/terraform-docs/terraform-docs

## Requirements

| Name | Version |
|------|---------|
| terraform | >= 1.3.2 |
| aws | >= 4.5.0 |

## Providers

| Name | Version |
|------|---------|
| aws | >= 4.5.0 |

## Modules

No Modules.

## Resources

| Name |
|------|
| [aws_cloudfront_function](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_function) |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| code | The code of the function | `string` | n/a | yes |
| product | Tag used to group resources according to product | `string` | n/a | yes |
| comment | Description of the function. | `string` | `null` | no |
| name | Name of the cloudfront function module. If null, will default to application\_tag. | `string` | `null` | no |
| publish | Whether to publish creation/change as Live CloudFront Function Version. Defaults to true. | `bool` | `true` | no |
| runtime | The runtime to use for the function. | `string` | `"cloudfront-js-1.0"` | no |

## Outputs

| Name | Description |
|------|-------------|
| arn | ARN of the CloudFront Function |
| name | Name of the CloudFront Function |
