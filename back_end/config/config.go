package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	Server      server      `yaml:"server"`
	Database    database    `yaml:"database"`
	Jwt         jwt         `yaml:"jwt"`
	Minio       minio       `yaml:"minio"`
	Filebrowser filebrowser `yaml:"filebrowser"`
	Proxy       proxy       `yaml:"proxy"`
	Redis       redis       `yaml:"redis"`
	Email       email       `yaml:"email"`
	Miniprogram miniprogram `yaml:"miniprogram"`
	Path        string
}

type server struct {
	Port          int    `yaml:"port"`
	QrcodeRelease string `yaml:"qrcode_release"`
}

type database struct {
	Driver string `yaml:"driver"`
	User   string `yaml:"user"`
	Passwd string `yaml:"password"`
	Addr   string `yaml:"addr"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbname"`
	Zone   string `yaml:"zone"`
}

type jwt struct {
	Key string `yaml:"key"`
}

type minio struct {
	Endpoint     string `yaml:"endpoint"`
	AccessKey    string `yaml:"accesskey"`
	AccessSecret string `yaml:"accesssecret"`
	UseSSL       bool   `yaml:"ssl"`
	Bucket       string `yaml:"bucket"`
}

type filebrowser struct {
	Url      string `yaml:"url"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Bucket   string `yaml:"bucket"`
	ShareID  string `yaml:"share_id"`
}

type proxy struct {
	Url string `yaml:"url"`
}

type redis struct {
	Host     string `yaml:"host"`
	Password string `yaml:"password"`
	Database int    `yaml:"database"`
}

type email struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
}

type miniprogram struct {
	AppID                      string `yaml:"appid"`
	AppSecret                  string `yaml:"appsecret"`
	Mchid                      string `yaml:"mchid"`
	MchCertificateSerialNumber string `yaml:"mchcertificateserialnumber"`
	MchAPIv3Key                string `yaml:"mchapiv3key"`
}

var Config config

func Init(filepath string) {
	filePath := filepath + "/config.yaml"
	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Println("读取配置文件失败", err.Error())
		return
	}
	err = yaml.Unmarshal(file, &Config)
	if err != nil {
		log.Println("序列化配置文件失败", err.Error())
		return
	}
	Config.Path = filePath
}
