# Examples for Go and Python Lambda deployments.

AWS Lambda implementations of a data factory in Go and Python. 
- Extract data from an rest api
- Validate results
- Write to parquet

## Content
- AWS Infrastructure in [terraform/](terraform)  
- The rest api is located in [python-api/](python-api)  
- Lambda function in python in [python-requests/](python-requests)  
- Lambda function in Go in [go-requests/](go-requests) 

## Deployment
See READMEs in the particular directories:
1. Terraform
1. Go-requests
1. Python-requests
1. Python-api

#Warning:
Here and there hardcoded S3 bucket names exist here and there. Replace these with you own S3 bucket name.

