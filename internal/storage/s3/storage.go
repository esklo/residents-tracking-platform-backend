package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	def "github.com/esklo/residents-tracking-platform/internal/storage"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"
)

var _ def.Storage = (*Storage)(nil)

type Storage struct {
	endpoint, region, keyId, accessKey, bucket string
	client                                     *s3.Client
}

func NewStorage(endpoint, region, keyId, accessKey, bucket string) *Storage {
	log.Println("storage init", endpoint, region, keyId, accessKey, bucket)
	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:           endpoint,
			SigningRegion: region,
		}, nil
	})

	cfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(keyId, accessKey, "")),
		awsConfig.WithRegion(region),
		awsConfig.WithEndpointResolverWithOptions(customResolver),
	)
	if err != nil {
		log.Fatalf("can not create new file repository: %s", err)
	}

	return &Storage{
		endpoint:  endpoint,
		region:    region,
		keyId:     keyId,
		accessKey: accessKey,
		bucket:    bucket,
		client:    s3.NewFromConfig(cfg),
	}
}

func (s Storage) PutFile(r io.ReadSeeker) (path string, mime *mimetype.MIME, err error) {
	if mime, err = getMimeInfo(r); err != nil {
		return "", nil, errors.Wrap(err, "can not get mime info")
	}

	path = fmt.Sprintf("storage/%s%s", uuid.New().String(), mime.Extension())
	log.Printf("path: %#v", path)
	if _, err = r.Seek(0, io.SeekStart); err != nil {
		return
	}
	log.Printf("s.b: %#v", s.bucket)
	_, err = s.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      aws.String(s.bucket),
		Key:         aws.String(path),
		Body:        r,
		ContentType: aws.String(mime.String()),
	})
	if err != nil {
		return "", nil, errors.Wrap(err, "can not put object")
	}
	return path, mime, nil
}

func (s Storage) GetFile(path string) (r io.ReadCloser, err error) {
	id := getMD5Hash(path)
	tmpFilePath := filepath.Join(os.TempDir(), id)

	_, err = os.Stat(tmpFilePath)
	if err == nil {
		file, err := os.Open(tmpFilePath)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	pr := s3.NewPresignClient(s.client)
	_, err = pr.PresignGetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	}, func(opts *s3.PresignOptions) {
		opts.Expires = time.Duration(1 * int64(time.Second))
	})
	if err != nil {
		return nil, err
	}

	getObjectOutput, err := s.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})

	if err != nil {
		return nil, err
	}

	tmpFile, err := os.Create(tmpFilePath)
	if err != nil {
		return nil, err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, getObjectOutput.Body)
	if err != nil {
		return nil, err
	}

	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

func (s Storage) GetFileUrl(path string) (url string, err error) {
	return fmt.Sprintf("%s/%s/%s", s.endpoint, s.bucket, path), nil
}

func (s Storage) DeleteFile(path string) error {
	_, err := s.client.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(path),
	})
	if err != nil {
		return err
	}
	return nil
}
