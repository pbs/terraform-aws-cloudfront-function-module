module "cloudfront_function" {
  source = "../.."

  product = var.product

  code = file("function.js")
}
