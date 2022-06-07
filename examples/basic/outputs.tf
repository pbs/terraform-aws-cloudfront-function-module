output "arn" {
  description = "ARN of the function"
  value       = module.cloudfront_function.arn
}

output "name" {
  description = "Name of the function"
  value       = module.cloudfront_function.name
}
