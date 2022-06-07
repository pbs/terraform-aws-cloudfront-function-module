locals {
  name    = var.name != null ? var.name : var.product
  comment = var.comment != null ? var.comment : "CloudFront function: ${local.name}"
}
