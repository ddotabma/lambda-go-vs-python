resource "aws_lambda_function" "python_requests" {
  s3_bucket = aws_s3_bucket.blog-bucket.bucket
  s3_key = "python-requests.zip"
  function_name = "api-calls-python"
  handler = "api_calls.handler"
  runtime = "python3.7"
  role = aws_iam_role.iam_for_lambda.arn
  timeout = 30
}


