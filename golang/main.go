package main 

import (
	"context"
	"regexp"
	"fmt"

	"github.com/nael-fridhi/csv-to-parquet-aws/golang/awsfuncs"
	"github.com/nael-fridhi/csv-to-parquet-aws/golang/csvparquet"
	"github.com/aws/aws-lambda-go/events"
)

func main() {

	var ctx context.Context
	var s3Event events.S3Event
    bucketName, csvObjectKey := lambda.Start(awsfuncs.GetObjectMetadata(ctx, s3Event))
	
	fmt.Println(bucketName)
	fmt.Println("ooooooooooooooooooooooooooooooooooooooo")
	fmt.Println(csvObjectKey)
	// Establish the parquet objectKey 
	//csvPath := "csv/titanic.csv"
	pref := regexp.MustCompile(`(^csv)`)
	aux := pref.ReplaceAllString(csvObjectKey, "parquet")
	suf := regexp.MustCompile(`(\.csv$)`)
	parquetObjectKey := suf.ReplaceAllString(aux, ".parquet")


	awsfuncs.WriteObjectToFile(bucketName, csvObjectKey)
	csvparquet.ConvertCsvToParquet()
	awsfuncs.WriteObjectToBucket(bucketName, parquetObjectKey)
}