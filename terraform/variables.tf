variable "user_name" {}
variable "region" {}
variable "API" {}

variable "python_requests_layer_arn"{
 default = "arn:aws:lambda:eu-west-1:756285606505:layer:python-requests-layer:18"
}