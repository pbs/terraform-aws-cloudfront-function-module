resource "aws_cloudfront_function" "function" {
  name    = local.name
  runtime = var.runtime
  comment = local.comment
  publish = var.publish
  code    = var.code
}
