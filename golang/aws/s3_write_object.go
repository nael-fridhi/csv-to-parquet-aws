package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// WriteObjectToBucket Write object to s3 bucket given the Object Key and Object value
func WriteObjectToBucket(sess string) {

	// Create an uploader with the session and default options
    uploader := s3manager.NewUploader(sess)
    f, err  := os.Open(path)
    if err != nil {
            fmt.Println("failed to open file")
        return
    }

    //Take the filename from path
    file := filepath.Base(path)
    objectKey := "//parquet//"+file
	
	// Upload the file to S3.
    _, err := uploader.Upload(&s3manager.UploadInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(objectKey),
        Body:   f,
    })
    if err != nil {
        fmt.Errorf("failed to upload file, %v", err)
        return
    }
}