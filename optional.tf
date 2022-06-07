variable "name" {
  description = "Name of the cloudfront function module. If null, will default to application_tag."
  default     = null
  type        = string
}

variable "runtime" {
  description = "The runtime to use for the function."
  type        = string
  default     = "cloudfront-js-1.0"
}

variable "comment" {
  description = "Description of the function."
  type        = string
  default     = null
}

variable "publish" {
  description = "Whether to publish creation/change as Live CloudFront Function Version. Defaults to true."
  type        = bool
  default     = true
}
