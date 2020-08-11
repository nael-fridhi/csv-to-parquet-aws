package main 

import (
	"context"
	"regexp"

	"github.com/nael-fridhi/csv-to-parquet-aws/golang/awsfuncs"
	"github.com/nael-fridhi/csv-to-parquet-aws/golang/csvparquet"
	"github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/xitongsys/parquet-go/parquet"
)

func main() {

    var ctx context.Context
    var s3Event events.S3Event
	bucketName, csvObjectKey := awsfuncs.GetObjectMetadata(ctx, s3Event)

	// Establish the parquet objectKey 
	csvPath := "csv/titanic.csv"
	pref := regexp.MustCompile(`(^csv)`)
	aux := pref.ReplaceAllString(csvPath, "parquet")
	suf := regexp.MustCompile(`(\.csv$)`)
	parquetObjectKey := suf.ReplaceAllString(aux, ".parquet")


	awsfuncs.writeObjectToFile(bucketName, objectKey)
	csvparquet.ConvertCsvToParquet()
	awsfuncs.writeObjectToBucket(bucketName, parquetObjectKey)
}