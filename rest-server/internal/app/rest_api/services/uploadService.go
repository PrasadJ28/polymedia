package services

import (
    "context"
    "fmt"
    "time"

    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/PrasadJ28/gin-rest-server/internal/app/rest_api/storage"
)
type CompletedPart struct {
    PartNumber int    `json:"partNumber"`
    ETag       string `json:"etag"`
}
type Upload struct {
    store         *storage.MinioStorage
    presignExpiry time.Duration
}

func NewUploadService(store *storage.MinioStorage) *Upload {
    return &Upload{
        store:         store,
        presignExpiry: 15 * time.Minute,
    }
}

func (u *Upload) StartUpload(filename string, filesize int64) (string, int, int64, error) {
    result, err := u.store.InternalClient.CreateMultipartUpload(
        context.Background(),
        &s3.CreateMultipartUploadInput{
            Bucket: aws.String(u.store.Bucket),
            Key:    aws.String(filename),
        },
    )
    if err != nil {
        return "", 0, 0, fmt.Errorf("failed to start multipart upload: %w", err)
    }

    const partSize = 10 * 1024 * 1024 // 10MB per chunk
    totalParts := int((filesize + partSize - 1) / partSize)

    return *result.UploadId, totalParts, partSize, nil
}

func (u *Upload) PresignPart(uploadId string, key string, partNumber int) (string, error) {
    result, err := u.store.PresignClient.PresignUploadPart(
        context.Background(),
        &s3.UploadPartInput{
            Bucket:     aws.String(u.store.Bucket),
            Key:        aws.String(key),
            UploadId:   aws.String(uploadId),
            PartNumber: aws.Int32(int32(partNumber)),
        },
        s3.WithPresignExpires(u.presignExpiry),
    )
    if err != nil {
        return "", fmt.Errorf("failed to presign part %d: %w", partNumber, err)
    }

    return result.URL, nil
}

func (u *Upload) CompleteUpload(uploadId string, key string, parts []CompletedPart) error {
    var completedParts []types.CompletedPart

    for _, part := range parts {
        completedParts = append(completedParts, types.CompletedPart{
            PartNumber: aws.Int32(int32(part.PartNumber)),
            ETag:       aws.String(part.ETag),
        })
    }

    _, err := u.store.InternalClient.CompleteMultipartUpload(
        context.Background(),
        &s3.CompleteMultipartUploadInput{
            Bucket:   aws.String(u.store.Bucket),
            Key:      aws.String(key),
            UploadId: aws.String(uploadId),
            MultipartUpload: &types.CompletedMultipartUpload{
                Parts: completedParts,
            },
        },
    )
    if err != nil {
        return fmt.Errorf("failed to complete multipart upload: %w", err)
    }

    return nil
}
