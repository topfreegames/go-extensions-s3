package s3

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/spf13/viper"
)

// Client is a wrapper over the official aws s3 package
// only implements used functions
type Client struct {
	client s3iface.S3API
	bucket string
	folder string
}

// NewClient ctor
func NewClient(prefix string, conf *viper.Viper) (*Client, error) {
	region := conf.GetString(fmt.Sprintf("%s.region", prefix))
	accessKey := conf.GetString(fmt.Sprintf("%s.accessKey", prefix))
	secretAccessKey := conf.GetString(fmt.Sprintf("%s.secretAccessKey", prefix))
	credentials := credentials.NewStaticCredentials(accessKey, secretAccessKey, "")

	forcePathStyle := true
	awsConfig := &aws.Config{
		Region:           &region,
		Credentials:      credentials,
		S3ForcePathStyle: &forcePathStyle,
	}

	endpoint := conf.GetString(fmt.Sprintf("%s.endpoint", prefix))
	if endpoint != "" {
		awsConfig.Endpoint = &endpoint
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, err
	}
	svc := s3.New(sess)
	s3 := s3iface.S3API(svc)
	return &Client{
		client: s3,
		bucket: conf.GetString(fmt.Sprintf("%s.bucket", prefix)),
		folder: conf.GetString(fmt.Sprintf("%s.folder", prefix)),
	}, nil
}

func streamToByte(stream *io.ReadCloser) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(*stream)
	return buf.Bytes()
}

// GetObject gets an object from s3
func (c Client) GetObject(path string) ([]byte, error) {
	splittedString := strings.SplitN(path, "/", 2)
	if len(splittedString) < 2 {
		return nil, fmt.Errorf("Invalid path")
	}
	bucket := splittedString[0]
	objKey := splittedString[1]
	params := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &objKey,
	}
	resp, err := c.client.GetObject(params)
	if err != nil {
		return nil, err
	}
	return streamToByte(&resp.Body), nil
}

// PutObject puts an object into s3
func (c Client) PutObject(path string, body *[]byte) error {
	b := bytes.NewReader(*body)
	params := &s3.PutObjectInput{
		Bucket: &c.bucket,
		Key:    &path,
		Body:   b,
	}
	_, err := c.client.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}

// MakePath concatenates folder with key
func (c Client) MakePath(k string) string {
	return fmt.Sprintf("%s/%s", c.folder, k)
}

// PutObjectRequest return a presigned url for uploading a file to s3
func (c Client) PutObjectRequest(key, acl string) (string, http.Header, error) {
	path := c.MakePath(key)
	params := &s3.PutObjectInput{
		ACL:    &acl,
		Bucket: &c.bucket,
		Key:    &path,
	}
	req, _ := c.client.PutObjectRequest(params)
	url, header, err := req.PresignRequest(900 * time.Second)
	if err != nil {
		return "", nil, err
	}
	return url, header, nil
}

// DeleteObject puts an object into s3
func (c Client) DeleteObject(key string) error {
	path := c.MakePath(key)
	params := &s3.DeleteObjectInput{
		Bucket: &c.bucket,
		Key:    &path,
	}
	_, err := c.client.DeleteObject(params)
	if err != nil {
		return err
	}
	return nil
}

// PutObjectInput puts an object into s3, if params.Bucket or params.Body
// are equal nil, they will be overwrite
func (c Client) PutObjectInput(params *s3.PutObjectInput, body *[]byte) error {
	b := bytes.NewReader(*body)
	if params.Bucket == nil {
		params.Bucket = &c.bucket
	}
	if params.Body == nil {
		params.Body = b
	}
	_, err := c.client.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}