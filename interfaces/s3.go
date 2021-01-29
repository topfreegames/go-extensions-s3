package interfaces

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"net/http"
)

// S3 is the minimum interface a S3Client must implement
type S3 interface {
	DeleteObject(key string) error
	PutObjectRequest(key, acl string) (string, http.Header, error)
	PutObject(key string, body *[]byte) error
	PutObjectInput(params *s3.PutObjectInput, body *[]byte) error
}