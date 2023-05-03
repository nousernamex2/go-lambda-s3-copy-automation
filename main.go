package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var sourceBucket = "sourcebucket"
var destinationBucket = "arn:aws:s3:::destination-bucket/SUBFOLDER/"
var prefix = "my-aws-prefix"

func CopyJobS3() (string, error) {

	// Set up an AWS session local
	sessManagementProd := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("eu-central-1"),
	}))
	// Create an S3 client
	svcManagementProd := s3.New(sessManagementProd)

	// Create the input for the list objects call
	findObject := &s3.ListObjectsInput{
		Bucket: &sourceBucket,
		Prefix: &prefix,
	}

	// List the objects in the bucket
	listBucket, err := svcManagementProd.ListObjects(findObject)
	if err != nil {
		return "Failed to list objects: ", err
	}
	fmt.Sprint("Files in Bucket: ", listBucket)

	// Find the latest uploaded file
	var sourceObjectKeyLatestFile *s3.Object
	var latestTime time.Time
	for _, obj := range listBucket.Contents {
		if obj.LastModified.After(latestTime) {
			sourceObjectKeyLatestFile = obj
			latestTime = *obj.LastModified
		}
	}

	// Set the S3 COPY operation input parameters
	input := &s3.CopyObjectInput{
		Bucket:     aws.String(destinationBucket),
		CopySource: aws.String(fmt.Sprintf("%s/%s", sourceBucket, *sourceObjectKeyLatestFile.Key)),
		Key:        aws.String(*sourceObjectKeyLatestFile.Key),
		ACL:        aws.String("bucket-owner-full-control"),
	}

	// Copy the object
	result, err := svcManagementProd.CopyObjectWithContext(context.Background(), input)
	if err != nil {
		return "Failed to copy object: ", err
	}
	returnMessageCopyObject := fmt.Sprint("Copied object with ID: ", result.CopyObjectResult.ETag)
	returnMessageCopyObject += fmt.Sprint("Name: ", *sourceObjectKeyLatestFile.Key)
	returnMessageCopyObject += fmt.Sprint("from source Bucket: ", sourceBucket)
	returnMessageCopyObject += fmt.Sprint("to destination Bucket: ", destinationBucket)
	returnMessageCopyObject += fmt.Sprint("Copy Succeeded!", nil)
	return returnMessageCopyObject, nil
}

func main() {
	lambda.Start(CopyJobS3)
}
