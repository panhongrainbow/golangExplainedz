package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var prefix string    // For instance, list all ojbects under /bar/foo/ ...
var delimeter string // ... using "/" as delimeter

func main() {
	paginator := s3.NewListObjectsV2Paginator(awsS3Client, &s3.ListObjectsV2Input{
		Bucket:    aws.String(AWS_S3_BUCKET),
		Prefix:    aws.String(prefix),
		Delimiter: aws.String(delimeter),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(context.TODO())
		if err != nil {
			// Error handling goes here
		}
		for _, obj := range page.Contents {
			// Do whatever you need with each object "obj"
			fmt.Println(obj.Key)
		}
	}
}
