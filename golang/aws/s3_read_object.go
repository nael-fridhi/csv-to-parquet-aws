package aws

import (
	"os"
)

func WriteObjectToFile() {

	sess := session.New(&aws.Config{Region: aws.String("us-east-1")})
	s3svc := s3.New(sess)
	result, err := s3svc.ListBuckets(&s3.ListBucketsInput{})
	if err != nil {
		fmt.Println("Failed to list buckets", err)
		return
	}

	fmt.Println("Buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("%s : %s\n", aws.StringValue(bucket.Name), bucket.CreationDate)
	}
	out, err := s3svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	fmt.Println(out)
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	// Create a file to write the S3 Object contents to.
	path := "/tmp/csvfile.csv"
	f, err := os.Create("/tmp/csvfile.csv")
	if err != nil {
		fmt.Println("failed to create file", err)
		return
	}

	// Write the contents of S3 Object to the file
	n, err := downloader.Download(f, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		fmt.Println("failed to download file", err)
		return
	}
	fmt.Printf("file downloaded, %d bytes\n", n)

}
