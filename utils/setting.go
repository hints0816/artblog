package utils

import (
	"fmt"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string
	Domain   string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisHost     string
	RedisPort     string
	RedisUser     string
	RedisPassWord string

	QiniuAccessKey   string
	QiniuSecretKey   string
	Scope      string

	AccessKey   string
	SecretKey   string
	Bucket      string
	MinioServer string
	
	ClientId     string
	ClientSecret string
	RedirectUri  string

	WebMaster  string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检查文件路径:", err)
	}
	LoadServer(file)
	LoadData(file)
	LoadMinio(file)
	LoadQiniu(file)
	LoadGithub(file)
	LoadBase(file)
}

func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("89js82js72")
	Domain = file.Section("server").Key("Domain").String()
}

func LoadData(file *ini.File) {
	Db = file.Section("database").Key("Db").MustString("debug")
	DbHost = file.Section("database").Key("DbHost").MustString("localhost")
	DbPort = file.Section("database").Key("DbPort").MustString("3306")
	DbUser = file.Section("database").Key("DbUser").MustString("ginblog")
	DbPassWord = file.Section("database").Key("DbPassWord").MustString("admin123")
	DbName = file.Section("database").Key("DbName").MustString("ginblog")

	RedisHost = file.Section("redis").Key("RedisHost").MustString("localhost")
	RedisPort = file.Section("redis").Key("RedisPort").MustString("6379")
	RedisUser = file.Section("redis").Key("RedisUser").String()
	RedisPassWord = file.Section("redis").Key("RedisPassWord").String()
}

func LoadMinio(file *ini.File) {
	AccessKey = file.Section("minio").Key("AccessKey").String()
	SecretKey = file.Section("minio").Key("SecretKey").String()
	Bucket = file.Section("minio").Key("Bucket").String()
	MinioServer = file.Section("minio").Key("MinioServer").String()
}

func LoadQiniu(file *ini.File) {
	QiniuAccessKey = file.Section("qiniu").Key("AccessKey").String()
	QiniuSecretKey = file.Section("qiniu").Key("SecretKey").String()
	Scope = file.Section("qiniu").Key("Scope").String()
}

func LoadGithub(file *ini.File) {
	ClientId = file.Section("github").Key("ClientId").String()
	ClientSecret = file.Section("github").Key("ClientSecret").String()
	RedirectUri = file.Section("github").Key("RedirectUri").String()
}

func LoadBase(file *ini.File) {
	WebMaster = file.Section("base").Key("WebMaster").String()
}
