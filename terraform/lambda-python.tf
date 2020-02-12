resource "aws_lambda_function" "python_requests" {
  s3_bucket = aws_s3_bucket.blog_bucket.bucket
  s3_key = "python-requests.zip"
  function_name = "api-calls-python"
  handler = "api_calls.handler"
  runtime = "python3.7"
  role = aws_iam_role.requests.arn
  timeout = 30
  memory_size = 320
  //  layers = [
  //    var.python_requests_layer_arn]
  environment {
    variables = {
      API = var.API
    }
  }
}


resource "aws_lambda_layer_version" "python_requests_layer" {
  layer_name = "python-requests-layer"
  s3_bucket = aws_s3_bucket.blog_bucket.bucket
  s3_key = "python-requests-layer.zip"
  compatible_runtimes = [
    "python3.7"]
}