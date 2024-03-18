package model

import (
	"bytes"
	"context"
	"fmt"
	"hello/utils"
	"hello/utils/errormsg"
	"hello/utils/fileUtils"
	"hello/utils/uuid"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"strconv"

	"mime/multipart"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"

	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
)


var (
	minioClient *minio.Client
	minioErr error

	upToken string
	formUploader *storage.FormUploader
)

func InitMinio() {
	endpoint := utils.MinioServer
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

func InitQiniu() {
	accessKey := utils.QiniuAccessKey
	secretKey := utils.QiniuSecretKey
	scope := utils.Scope

	putPolicy := storage.PutPolicy{
		Scope: scope,
	}
	mac := auth.New(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	cfg.Region = &storage.ZoneHuanan
	cfg.UseHTTPS = true
	cfg.UseCdnDomains = false

	formUploader = storage.NewFormUploader(&cfg)
}

func UpLoadFileQiniu(file multipart.File, contentType string, fileSize int64) (string, int) {
	fSrc, _ := ioutil.ReadAll(file)
	defer file.Close()

	fileType := fileUtils.GetFileType(fSrc[:10])
	uuid := uuid.GetSnowFlakeID()
	fileName := strconv.FormatInt(uuid, 10) + "." + fileType
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	
	err := formUploader.Put(context.Background(), &ret, upToken, fileName, bytes.NewReader(fSrc), fileSize, &putExtra)
	if err != nil {
		return "", errormsg.ERROR
	}
	return "http://cdn.ceccc.space/" + fileName, errormsg.SUCCSE
}

func UpLoadFileImgQiniu(img image.Image) (string, int) {
	emptyBuff := bytes.NewBuffer(nil)
	png.Encode(emptyBuff, img)
	fSrc := emptyBuff.Bytes()
	fileSize := int64(len(fSrc))

	uuid := uuid.GetSnowFlakeID()
	fileName := strconv.FormatInt(uuid, 10) + ".jpg"
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "github logo",
		},
	}
	
	err := formUploader.Put(context.Background(), &ret, upToken, fileName, bytes.NewReader(fSrc), fileSize, &putExtra)
	if err != nil {
		return "", errormsg.ERROR
	}
	return "http://cdn.ceccc.space/" + fileName + "?imageView2/2/w/400/format/jpeg", errormsg.SUCCSE
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
	return "http://" + utils.MinioServer + "/blog/" + fileName, errormsg.SUCCSE
}

func UpLoadFileImg(img image.Image) (string, int) {
	emptyBuff := bytes.NewBuffer(nil)
	png.Encode(emptyBuff, img)
	fSrc := emptyBuff.Bytes()
	buf := bytes.NewBuffer(fSrc)
	uploadfile := ioutil.NopCloser(buf)
	fileSize := int64(len(fSrc))
	defer uploadfile.Close()

	uuid := uuid.GetSnowFlakeID()
	fileName := strconv.FormatInt(uuid, 10) + ".webp"
	fmt.Println(minioClient.ListBuckets(context.Background()))
	uploadInfo, err := minioClient.PutObject(context.Background(), "blog", fileName, uploadfile, fileSize, minio.PutObjectOptions{ContentType: "image/webp"})
	if err != nil {
		return "", errormsg.ERROR
	}
	fmt.Println("Successfully uploaded bytes: ", uploadInfo)
	return "http://" + utils.MinioServer + "/blog/" + fileName, errormsg.SUCCSE
}
