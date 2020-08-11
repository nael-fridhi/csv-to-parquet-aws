package main 

import (
	"context"
	"regexp"
	"fmt"

	"github.com/nael-fridhi/csv-to-parquet-aws/golang/awsfuncs"
	"github.com/nael-fridhi/csv-to-parquet-aws/golang/csvparquet"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)


func handler(ctx context.Context, s3Event events.S3Event) {
	fmt.Println("oooooooooooooooooo START ooooooooooooooooooo")
	bucketName, csvObjectKey := awsfuncs.GetObjectMetadata(ctx, s3Event)
	
	fmt.Println(bucketName)
	fmt.Println("ooooooooooooooooooooooooooooooooooooooo")
	fmt.Println(csvObjectKey)
	fmt.Println("oooooooooooooooooooooooooooooooooooooooo")

	// Establish the parquet objectKey 
	//csvPath := "csv/titanic.csv"
	pref := regexp.MustCompile(`(^csv)`)
	aux := pref.ReplaceAllString(csvObjectKey, "parquet")
	suf := regexp.MustCompile(`(\.csv$)`)
	parquetObjectKey := suf.ReplaceAllString(aux, ".parquet")

	fmt.Println("oooooooooooooooooooooooooooooooooooooooo")
	fmt.Println(parquetObjectKey)
	fmt.Println("oooooooooooooooooooooooooooooooooooooooo")

	awsfuncs.WriteObjectToFile(bucketName, csvObjectKey)
	csvparquet.ConvertCsvToParquet()
	awsfuncs.WriteObjectToBucket(bucketName, parquetObjectKey)
	fmt.Println("oooooooooooooooooo END ooooooooooooooooooo")
}
func main() {
	lambda.Start(handler)
}