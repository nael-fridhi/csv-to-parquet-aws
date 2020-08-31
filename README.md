# AWS S3 & Lambda CSV to Parquet using Golang and Spark Scala

This repository contains sample of converting a CSV file which is uploaded into AWS S3 bucket to Parquet format.

The upload of a CSV file into S3 bucket will trigger a lambda function to convert this object into parquet and then write the result to another prefix in the bucket as shown in the image below.

![Csv To Parquet AWS Architecture](./data/csv_to_parquet_aws_simple_architecture.jpg)


We have implemented this feature using two different programming language.

## Golang

For the golang you have to:

1. build the binary for your module
    `GOARCH=amd64 GOOS=linux go build -gcflags='-N -l' -o . .`

2. Package the binary:
    `zip function.zip binaryFile`

3. Also sometinmes you'll need to set the executable bit in the zipfile. There are a bunch of ways to do this, if you want to do it on windows, you'll need to run a [python script](https://stackoverflow.com/questions/57964626/permissions-denied-when-trying-to-invoke-go-aws-lambda-function) which i find it in the stackoverflow.

## Spark Scala

For spark scala:

1. package a JAR/ZIP file using sbt or maven including the dependencies

2. Pass this Jar file to the lambda function


# AWS 

1. Create A Role that allow lambda to interact with S3 bucket and Log to CloudWatch

2. Create Lambda function

3. Create Bucket and the two folder (prefix) for csv files and parquet files

4. Create Event from s3 properties window to trigger lambda function on upload 

## DevOps

- Using terraform as an infrastructure as a code tool to automate the creation and configuration of the bucket and the lambda function. 

- Using the AWS CodeBuild for the Continous Integration: 
 1. Build the zip file containing the code of our lambda function
 2. execute the scripts terraform to build the infrastructure

- The state of our infrastructure is saved in the bucket  

## Data in Depth

What if we have a big volume of csv files. In this case, lambda functions can not be the best choice. In fact, Lambda function 
can't be run more than 15 min so in the case where we process for instance 1 TB of data lambda will timeout. So, we have to think about building an ETL to manage our data pipeline such as using Spark-based ETL processing running on Amazon's Elastic Map Reduce (EMR) Hadoop platform.

### ETL CHOICE 

This table below shows two different solutions for ETL and the difference between them.
1. AWS Glue
2. AWS EMR 


| Criteria                 | Amazon Glue          | Amazon EMR                 |
| -------------------------| -------------------- |:--------------------:      |
| Deployment Types         | Serverless           | Server Platform ( Cluster )|
| Pricing                  | High                 | Low                        |
| Flexibility & Scalability| Flexible             | Harder to scale            |
| ETL operations           | Better               | Not so good                |  
| Performance              | Slower & less stable | Faster and more stable     |




![Csv To Parquet AWS Architecture Using EMR](./data/aws_glue_EMR.jpg)

- We can orchestrate the piepline using AWS Data Pipeline

## Optimization (Thinking!!!!)

- We can improve this project with others feature:
  - Adding a pipeline to deploy version of lambda automatically after commit using travisCI,gitlabCI or CodeBuild ..
  - What if a lambda function fails or a spark job fails when talking about running or processing in EMR cluster ..

  - We can even change the way we process the data and using AWS Glue instead which is the serverless ETL service of AWS or we can also use EMR cluster (Hadoop) under the hood to process data.
  

# Resources: 

https://www.rittmanmead.com/blog/page/13/

https://docs.aws.amazon.com/datapipeline/latest/DeveloperGuide/dp-copydata-s3.html

https://www.terraform.io/docs/index.html

https://github.com/xitongsys/parquet-go