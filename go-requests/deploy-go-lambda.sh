export AWS_DEFAULT_REGION=eu-west-1
export AWS_PROFILE=github

GOOS=linux GOARCH=amd64 go build requests.go
zip -q go-requests.zip requests
aws s3 cp go-requests.zip s3://bdr-go-blog

aws lambda update-function-code --function-name api-calls-go --s3-bucket bdr-go-blog --s3-key go-requests.zip > /tmp/dumpgo