rm -r python
rm -r lambda-packaged/

set -e  # bash don't continue on error

docker run --rm -v "$(pwd)":/foo -w /foo lambci/lambda:build-python3.7 \
 pip install -r requirements.txt -t python/lib/python3.7/site-packages/

zip -rq9 python-requests-layer.zip python

aws s3 cp python-requests-layer.zip s3://bdr-go-blog/python-requests-layer.zip

export LAYER_VERSION_ARN=$(aws lambda publish-layer-version \
  --layer-name python-requests-layer \
  --content S3Bucket=bdr-go-blog,S3Key=python-requests-layer.zip \
  --compatible-runtimes python3.7 | jq -r ".LayerVersionArn")

mkdir lambda-packaged
cp *.py lambda-packaged/
cd lambda-packaged && zip -rq python-requests.zip *

aws s3 cp python-requests.zip s3://bdr-go-blog

aws ambda update-function-code --function-name api-calls-python --s3-bucket bdr-go-blog --s3-key python-requests.zip > /tmp/dumpr

aws lambda update-function-configuration \
--function-name api-calls-python \
--layers ${LAYER_VERSION_ARN}