package main 

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-lambda-go/events"
)

func main() {

	sess := session.New(&aws.Config{Region: aws.String("us-east-1")})

	bucketName, csvObjectKey := aws.GetObjectMetadata(ctx context.Context, s3Event events.S3Event)

	// Establish the parquet objectKey 
	pref := regexp.MustCompile(`(^csv)`)
	aux := pref.ReplaceAllString(csvPath, "parquet")
	suf := regexp.MustCompile(`(\.csv$)`)
	parquetObjectKey := suf.ReplaceAllString(aux, ".parquet")


	aws.writeObjectToFile(bucketName, objectKey, sess)
	parquet.ConvertCsvToParquet()
	aws.writeObjectToBucket(bucketName, parquetObjectKey, sess)
}