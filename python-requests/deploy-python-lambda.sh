export AWS_DEFAULT_REGION=eu-west-1
export AWS_PROFILE=bdr
rm -r lambda-packaged/
docker run --rm -v "$(pwd)":/foo -w /foo lambci/lambda:build-python3.7 \
 pip install -r requirements.txt -t lambda-packaged/


cp api_calls.py lambda-packaged/
cd lambda-packaged && zip -rq python-requests.zip *

aws s3 cp python-requests.zip s3://bdr-go-blog

aws --profile=bdr lambda update-function-code --function-name api-calls-python --s3-bucket bdr-go-blog --s3-key python-requests.zip