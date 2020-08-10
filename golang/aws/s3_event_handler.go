package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

// GetObjectMetadata used to get the uploaded object key and the bucket Name
func GetObjectMetadata(ctx context.Context, s3Event events.S3Event) (string, string) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", record.EventSource, record.EventTime, s3.Bucket.Name, s3.Object.Key)
		return s3.Bucket.Name, s3.Object.Key
	}
	return "", ""
}
