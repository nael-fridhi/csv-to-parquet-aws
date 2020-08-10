# AWS S3 & Lambda CSV to Parquet using Golang and Spark Scala

This repository contains sample of converting  a file which is putted into AWS S3 bucket from Csv to Parquet.

The upload of a CSV file into S3 bucket will trigger a lambda function to convert this object into parquet and then write the result to another prefix in the bucket as shown in the image below.

![Csv To Parquet AWS Architecture](./data/csv_to_parquet_aws_simple_architecture.jpg)

We have implemented this feature using two different programming language.

## Golang

For the golang you have to:

1. build the binary for your module
    `GOOS=linux go build golang/csvToParquet.go`

2. Package the binary:
    `zip function.zip binaryFile`

3. Also you'll need to set the executable bit in the zipfile. There are a bunch of ways to do this, if you want to do it on windows, you'll need to run a [python script](./golang/script.py) which i find it in the stackoverflow.

## Spark Scala

For spark scala:

1. package a JAR file using sbt including the spark dependencies

2. Pass this Jar file to the lambda function

## Optimization

This task is not optimized it's only a POC of a simple architecture in AWS to convert csv to parquet.

- Some features which could be add:
  - What if a lambda function fails? there is a default 3 times retry configured in AWS
  - We can use AWS Glue which is an ETL working using Hadoop Framework to process the data
  - We can orchestrate the piepline using AWS CodePipeline
  - What if we have too much csv files uploaded in the same time?
