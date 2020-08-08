# AWS S3 & Lambda CSV to Parquet using Golang and Spark Scala

This repository contains sample of converting  a file which is putted into AWS S3 bucket from Csv to Parquet.

The upload of a CSV file into S3 bucket will trigger a lambda function to convert this object into parquet and then write the result to another path in the bucket.

We implement this feature using two different programming language.

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
