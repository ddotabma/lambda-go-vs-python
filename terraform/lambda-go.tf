resource "aws_s3_bucket" "blog_bucket" {
  bucket = "${var.user_name}-go-blog"
}


resource "aws_lambda_function" "go_requests" {
  s3_bucket = aws_s3_bucket.blog_bucket.bucket
  s3_key = "go-requests.zip"
  function_name = "api-calls-go"
  handler = "requests"
  runtime = "go1.x"
  role = aws_iam_role.requests.arn
  timeout = 30

  environment {
    variables = {
      API = var.API
    }
  }
}


