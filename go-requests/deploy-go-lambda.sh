export AWS_DEFAULT_REGION=eu-west-1
export AWS_PROFILE=bdr

GOOS=linux GOARCH=amd64 go build requests.go
zip requests.zip requests
aws s3 cp requests.zip s3://bdr-go-blog

aws lambda update-function-code --function-name api-calls-go --s3-bucket bdr-go-blog --s3-key go-requests.zip