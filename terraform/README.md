# Terraform

Deployment steps:
- Terraform init
- Terraform apply

Terraform will create the S3 bucket that will contain the lambda code.  
Then it will try to update the lambda function using that bucket. This is not possible
since the zips are not in place yet. Run the bash scripts in go-requests and python-requests
to place the zips. Then run `terraform apply` again.

To apply terraform once the lambda layer is in place, run:  
`export layer_arn=$(aws lambda list-layer-versions --layer-name python-requests-layer  | jq -r ".LayerVersions[0].LayerVersionArn")`  
`terraform apply -var="python_requests_layer_arn=$layer_arn"`