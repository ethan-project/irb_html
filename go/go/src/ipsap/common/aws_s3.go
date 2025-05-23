
package common

import (
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
  "net/url"
	"bytes"
	"log"
	"fmt"
	"time"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
  "github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

const (
  S3_REGION = "ap-northeast-2"
  S3_SECRET_ID = "AKIAS2PIUWGJ6PYZR44A"
  S3_SECRET_KEY = "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"
)

func getAwsSession() (s *session.Session, err error) {
	s, err = session.NewSession(&aws.Config{
													Region: aws.String(S3_REGION),
													Credentials : credentials.NewStaticCredentials (
																				S3_SECRET_ID, //aws.String(S3_SECRET_ID), // "secret-id", // id
																				S3_SECRET_KEY, //aws.String(S3_SECRET_KEY), //"secret-key", // secret
																				""), // token은 지금은 비워 둘 수 있음
																					})

	return
}

//  S3 파일 삭제
func RemoveFileToS3(svrDir string) error {
	s, err := getAwsSession()
	if nil != err {
		log.Println(err)
		return err
	}

  svc := s3.New(session.Must(s, nil))
  _, err = svc.DeleteObject(&s3.DeleteObjectInput{
    Bucket: aws.String(Config.S3.BUCKET),
    Key:    aws.String(svrDir),
  })
  if err != nil {
    log.Printf("Unable to delete object %q from bucket %q, %v", svrDir, Config.S3.BUCKET, err)
    return err
  }

  err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
    Bucket: aws.String(Config.S3.BUCKET),
    Key:    aws.String(svrDir),
  })
  return err
}

//  S3 파일 존재 유무 확인
func IsFileExists(s *session.Session, svrDir string) bool {
  svc := s3.New(session.Must(s, nil))
  _, err := svc.HeadObject(&s3.HeadObjectInput{
    Bucket: aws.String(Config.S3.BUCKET),
    Key:    aws.String(svrDir),
  })

  if err != nil {
		log.Println(err)
    return false
  }

  return true
}

func AddFileToS3(key string, file multipart.File, fileHeader *multipart.FileHeader) error {
	s, err := getAwsSession()
	if nil != err {
		log.Println(err)
		return err
	}

	size := fileHeader.Size
  buffer := make([]byte, size)
  file.Read(buffer)

  _, err = s3.New(s).PutObject(&s3.PutObjectInput{
    Bucket:               aws.String(Config.S3.BUCKET),
    Key:                  aws.String(key),
    ACL:                  aws.String("private"),
    Body:                 bytes.NewReader(buffer),
    ContentLength:        aws.Int64(size),
    ContentType:          aws.String(http.DetectContentType(buffer)),
    ContentDisposition:   aws.String("attachment"),
    ServerSideEncryption: aws.String("AES256"),
  })
  return err
}

func DownloadFileToS3(c *gin.Context, key string, org_file_name string) {
	s, err2 := getAwsSession()
	if nil != err2 {
		log.Println(err2)
		return
	}

	if !IsFileExists(s, key) {
		FinishApiWithErrCd(c, Api_status_bad_request, Error_none_file)
		return
	}

	buff := &aws.WriteAtBuffer{}
	downloader := s3manager.NewDownloader(s)
	_, err := downloader.Download(buff,
		 &s3.GetObjectInput{
				 Bucket: aws.String(Config.S3.BUCKET),
				 Key:    aws.String(key),
		 })
	if err != nil {
			fmt.Println(err)
	}

	FileContentType := http.DetectContentType(buff.Bytes())
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%v", org_file_name))
	c.Writer.Header().Add("Content-Type", FileContentType)
	c.Data(Api_status_ok, FileContentType, buff.Bytes())
}

func GetPresignUrl(key string) (urlStr string) {
	s, err := getAwsSession()
	if nil != err {
		log.Println(err)
		return
	}

	svc := s3.New(s)
	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(Config.S3.BUCKET),
			Key:    aws.String(key),
	})
	urlStr, err = req.Presign(15 * time.Minute)
	if err != nil {
		log.Println("Failed to sign request", err)
	}
	return
}

func CopyToS3(copySrc string, newSrc string) error {
	s, err := getAwsSession()
	if nil != err {
		log.Println(err)
		return err
	}

	// copy할 대상 object에는 버킷 이름이 포함되야됨!
	e := url.QueryEscape(Config.S3.BUCKET + "/" + copySrc)
	svc := s3.New(s)
	input := &s3.CopyObjectInput{
    Bucket:     aws.String(Config.S3.BUCKET),
    CopySource: aws.String(e),
    Key:        aws.String(newSrc),
	}

	_, err = svc.CopyObject(input)
	if nil != err {
		log.Println(err)
	}
	return err
}
