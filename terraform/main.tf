provider "aws" {
    region = var.aws_region
}

# Create Bucket Configuration AWS S3 Bucket 

resource "random_id" "bucket_suffix_id" {
  byte_length = 2
}


# Create DataLake S3 bucket

resource "aws_s3_bucket" "bucket" {
  bucket        = "${var.bucket_domain_name}-${random_id.bucket_suffix_id.dec}"
  acl           = "private"
  force_destroy = true

  tags = {
    Name = "data_bucket"
  }
}

# Create Prefix for CSV Files
resource "aws_s3_bucket_object" "input_csv" {
    bucket = aws_s3_bucket.bucket.id
    acl    = "private"
    key    = "csv/"
    source = "/dev/null"
}

# Create prefix for parquet Files
resource "aws_s3_bucket_object" "output_parquet" {
    bucket = aws_s3_bucket.bucket.id
    acl    = "private"
    key    = "parquet/"
    source = "/dev/null"
}

# IAM Role for lambda function
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
      "Effect": "Allow"
    }
  ]
}
EOF
}

# Permission for triggering lambda

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = aws_lambda_function.lambda_func.arn
  principal     = "s3.amazonaws.com"
  source_arn    = aws_s3_bucket.bucket.arn
}

# Lambda Function 

resource "aws_lambda_function" "lambda_func" {
  s3_bucket     = var.lambda_zip_bucket
  s3_key = var.lambda_zip_file
  function_name = var.lambda_function_name
  role          = aws_iam_role.iam_for_lambda.arn
  handler       = var.handler
  runtime       = "go1.x"
}

# AWS S3 LAMBDA TRIGGER EVENT

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = aws_s3_bucket.bucket.id

  lambda_function {
    lambda_function_arn = aws_lambda_function.lambda_func.arn
    events              = ["s3:ObjectCreated:*"]
    filter_prefix       = "csv/"
    filter_suffix       = ".csv"
  }

  depends_on = [aws_lambda_permission.allow_bucket]
}