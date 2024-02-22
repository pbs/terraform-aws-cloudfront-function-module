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
| <a name="requirement_terraform"></a> [terraform](#requirement\_terraform) | >= 1.3.2 |
| <a name="requirement_aws"></a> [aws](#requirement\_aws) | >= 4.5.0 |

## Providers

| Name | Version |
|------|---------|
| <a name="provider_aws"></a> [aws](#provider\_aws) | 5.24.0 |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [aws_cloudfront_function.function](https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/cloudfront_function) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_code"></a> [code](#input\_code) | The code of the function | `string` | n/a | yes |
| <a name="input_product"></a> [product](#input\_product) | Tag used to group resources according to product | `string` | n/a | yes |
| <a name="input_comment"></a> [comment](#input\_comment) | Description of the function. | `string` | `null` | no |
| <a name="input_name"></a> [name](#input\_name) | Name of the cloudfront function module. If null, will default to application\_tag. | `string` | `null` | no |
| <a name="input_publish"></a> [publish](#input\_publish) | Whether to publish creation/change as Live CloudFront Function Version. Defaults to true. | `bool` | `true` | no |
| <a name="input_runtime"></a> [runtime](#input\_runtime) | The runtime to use for the function. | `string` | `"cloudfront-js-1.0"` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_arn"></a> [arn](#output\_arn) | ARN of the CloudFront Function |
| <a name="output_name"></a> [name](#output\_name) | Name of the CloudFront Function |
