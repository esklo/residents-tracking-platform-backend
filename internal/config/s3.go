package config

import (
	"errors"
	"os"
)

const (
	s3EndpointEnvName  = "AWS_S3_ENDPOINT"
	s3RegionEnvName    = "AWS_S3_REGION"
	s3KeyIdEnvName     = "AWS_S3_KEY_ID"
	s3AccessKeyEnvName = "AWS_S3_ACCESS_KEY"
	s3BucketEnvName    = "AWS_S3_BUCKET"
)

type S3Config interface {
	Endpoint() string
	Region() string
	KeyId() string
	AccessKey() string
	Bucket() string
}

type s3Config struct {
	endpoint  string
	region    string
	keyId     string
	accessKey string
	bucket    string
}

func NewS3Config() (S3Config, error) {
	endpoint := os.Getenv(s3EndpointEnvName)
	if len(endpoint) == 0 {
		return nil, errors.New("endpoint not found")
	}

	region := os.Getenv(s3RegionEnvName)
	if len(region) == 0 {
		return nil, errors.New("region not found")
	}

	keyId := os.Getenv(s3KeyIdEnvName)
	if len(keyId) == 0 {
		return nil, errors.New("keyId not found")
	}

	accessKey := os.Getenv(s3AccessKeyEnvName)
	if len(accessKey) == 0 {
		return nil, errors.New("accessKey not found")
	}

	bucket := os.Getenv(s3BucketEnvName)
	if len(bucket) == 0 {
		return nil, errors.New("bucket not found")
	}

	return &s3Config{
		endpoint:  endpoint,
		region:    region,
		keyId:     keyId,
		accessKey: accessKey,
		bucket:    bucket,
	}, nil
}

func (cfg *s3Config) Endpoint() string {
	return cfg.endpoint
}
func (cfg *s3Config) Region() string {
	return cfg.region
}
func (cfg *s3Config) KeyId() string {
	return cfg.keyId
}
func (cfg *s3Config) AccessKey() string {
	return cfg.accessKey
}
func (cfg *s3Config) Bucket() string {
	return cfg.bucket
}
