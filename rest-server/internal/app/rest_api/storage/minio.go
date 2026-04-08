package storage

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"         //for aws.String() which provides pointers ot string instead of direct string
	"github.com/aws/aws-sdk-go-v2/credentials" // NewStaticCredentialProvider holds access key and secrect
	"github.com/aws/aws-sdk-go-v2/service/s3"  // the actual S3 service client with all operations :CreateMultiPartUpload, UploadPart, CompleteMultipartUpload
)

// S3 client holds connections internally so we only pass pointers to ensure single instance
type MinioStorage struct {
	InternalClient *s3.Client        //serverside operations
	ExternalClient *s3.Client        // Build presigned URLs
	PresignClient  *s3.PresignClient //a wrapper for ExternalClient to generate time limited signed URLs
	Bucket         string            // the name of the Bucket
}

func NewMinioStorage(
	internalEndpoint string,
	externalEndpoint string,
	accessKey string,
	secretKey string,
	bucket string,
) (*MinioStorage, error) {

	//credentials object to sign every request using static keys from docker compose
	creds := credentials.NewStaticCredentialsProvider(
		accessKey,
		secretKey,
		"",
	)

	internalClient := s3.NewFromConfig(
		aws.Config{
			Region:      "us-east-1", //required for AWS, although minio does not have region for signature algorithm
			Credentials: creds,
		}, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(internalEndpoint) // talk to minio endpoint instead of AWS
			o.UsePathStyle = true                         //bucket name part of URL path
		},
	)

	externalClient := s3.NewFromConfig(
		aws.Config{
			Region:      "us-east-1",
			Credentials: creds,
		}, func(o *s3.Options) {
			o.BaseEndpoint = aws.String(externalEndpoint)
			o.UsePathStyle = true
		},
	)

	presignClient := s3.NewPresignClient(externalClient)
	_, err := internalClient.HeadBucket(context.Background(), &s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})

	if err != nil {
		return nil, fmt.Errorf("minio bucket '%s' not found or unreachable: %w", bucket, err)
	}

	return &MinioStorage{
		InternalClient: internalClient,
		ExternalClient: externalClient,
		PresignClient:  presignClient,
		Bucket:         bucket,
	}, nil

}
