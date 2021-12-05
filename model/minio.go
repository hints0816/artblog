package model

import (
	"bytes"
	"context"
	"fmt"
	"hello/utils"
	"hello/utils/errormsg"
	"hello/utils/fileUtils"
	"hello/utils/uuid"
	"io/ioutil"
	"log"
	"strconv"

	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *minio.Client
var minioErr error

func InitMinio() {
	endpoint := utils.QiniuSever
	accessKeyID := utils.AccessKey
	secretAccessKey := utils.SecretKey
	useSSL := false

	// Initialize minio client object.
	minioClient, minioErr = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if minioErr != nil {
		log.Fatalln(minioErr)
	}

	log.Printf("%#v\n", minioClient) // minioClient is now set up
}

func CreateBucket(bucketName string) {
	ctx := context.Background()
	location := "us-east-1"

	err = minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
	}
}

func UpLoadFile(file multipart.File, contentType string, fileSize int64) (string, int) {
	fSrc, _ := ioutil.ReadAll(file)
	defer file.Close()
	buf := bytes.NewBuffer(fSrc)
	uploadfile := ioutil.NopCloser(buf)
	defer uploadfile.Close()

	fileType := fileUtils.GetFileType(fSrc[:10])
	uuid := uuid.GetSnowFlakeID()
	fileName := strconv.FormatInt(uuid, 10) + "." + fileType
	fmt.Println(minioClient.ListBuckets(context.Background()))
	uploadInfo, err := minioClient.PutObject(context.Background(), "blog", fileName, uploadfile, fileSize, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		return "", errormsg.ERROR
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return "http://" + utils.QiniuSever + "/blog/" + fileName, errormsg.SUCCSE
}
