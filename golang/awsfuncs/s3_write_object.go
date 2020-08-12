package awsfuncs

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// WriteObjectToBucket Write object to s3 bucket given the Object Key and Object value
func WriteObjectToBucket(bucketName, parquetObjectKey string) {
	sess := session.New(&aws.Config{Region: aws.String("eu-west-1")})

	// Create an uploader with the session and default options
    uploader := s3manager.NewUploader(sess)
    f, err  := os.Open("/tmp/parquetFile.parquet")
    if err != nil {
            fmt.Println("failed to open file")
        return
    }

	
	// Upload the file to S3.
    _, err = uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(parquetObjectKey),
        Body:   f,
    })
    if err != nil {
        fmt.Errorf("failed to upload file, %v", err)
        return
    }
}