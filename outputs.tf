output "arn" {
  description = "ARN of the CloudFront Function"
  value       = aws_cloudfront_function.function.arn
}

output "name" {
  description = "Name of the CloudFront Function"
  value       = aws_cloudfront_function.function.name
}
