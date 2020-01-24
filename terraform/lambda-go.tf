resource "aws_s3_bucket" "blog-bucket" {
  bucket = "${var.user_name}-go-blog"
}


resource "aws_lambda_function" "frequent_copy" {
  s3_bucket = aws_s3_bucket.blog-bucket.bucket
  s3_key = "requests.zip"
  function_name = "api-calls-go"
  handler = "requests"
  runtime = "go1.x"
  role = aws_iam_role.iam_for_lambda.arn
  timeout = 30
}


resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}


resource "aws_iam_role_policy_attachment" "execution_role" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  role = aws_iam_role.iam_for_lambda.id
}