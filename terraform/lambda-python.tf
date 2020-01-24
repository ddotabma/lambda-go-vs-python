resource "aws_lambda_function" "python_requests" {
  s3_bucket = aws_s3_bucket.blog_bucket.bucket
  s3_key = "python-requests.zip"
  function_name = "api-calls-python"
  handler = "api_calls.handler"
  runtime = "python3.7"
  role = aws_iam_role.go_requests.arn
  timeout = 30

  environment {
    variables = {
      API = var.API
    }
  }
}


