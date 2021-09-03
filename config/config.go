package config

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"github.com/zjswh/go-tool/nacos"
	"gorm.io/gorm"
	"os"
)

var (
	GVA_CONFIG Server
	GVA_VP     *viper.Viper
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
	GVA_OSS_BUCKET *oss.Bucket
)

type UserInfo struct {
	AccountId int `json:"accountId"`
	Aid int `json:"aid"`
	Uin int `json:"uin"`
	Name string `json:"name"`
}

type Server struct {
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
	System System `json:"system"`
	Jwt    JWT    `json:"jwt"`
	Log    Log    `json:"log"`
	Email  Email  `json:"email"`
	Param  Param  `json:"param"`
	Oss    Oss    `json:"oss"`
}

type Mysql struct {
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Path         string `mapstructure:"path" json:"path" yaml:"path"`
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"`
	LogMode      bool   `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`
}

type Redis struct {
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

type System struct {
	Env    string `mapstructure:"env" json:"env" yaml:"env"`
	Addr   int    `mapstructure:"addr" json:"addr" yaml:"addr"`
	DbType string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`
}

type JWT struct {
	SigningKey string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`
}

type Log struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}

type Email struct {
	User      string `mapstructure:"user" json:"user" yaml:"user"`
	Host      string `mapstructure:"host" json:"host" yaml:"host"`
	Port      int    `mapstructure:"port" json:"port" yaml:"port"`
	Pass      string `mapstructure:"pass" json:"pass" yaml:"pass"`
	AdminUser string `mapstructure:"admin-user" json:"admin-user" yaml:"admin-user"`
}

type Param struct {
	WebHost         string `mapstructure:"web-host" json:"web-host" yaml:"web-host"`
	ActivityHost    string `mapstructure:"activity-host" json:"activity-host" yaml:"activity-host"`
	ProjectHost     string `mapstructure:"project-host" json:"project-host" yaml:"project-host"`
	CronHost        string `mapstructure:"cron-host" json:"cron-host" yaml:"cron-host"`
	MessageHost     string `mapstructure:"message-host" json:"message-host" yaml:"message-host"`
	EsHost          string `mapstructure:"es-host" json:"es-host" yaml:"es-host"`
	BGatewayHost    string `mapstructure:"b-gateway-host" json:"b-gateway-host" yaml:"b-gateway-host"`
	DirtyFilterHost string `mapstructure:"dirty-filter-host" json:"dirty-filter-host" yaml:"dirty-filter-host"`
	XCaStage        string `mapstructure:"x-ca-stage" json:"x-ca-stage" yaml:"x-ca-stage"`
}

type Oss struct {
	AccessKeyId     string `mapstructure:"access_key_id" json:"access_key_id" yaml:"access_key_id"`
	AccessKeySecret string `mapstructure:"access_key_secret" json:"access_key_secret" yaml:"access_key_secret"`
	Endpoint        string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
	Bucket          string `mapstructure:"bucket" json:"bucket" yaml:"bucket"`
}

type Nacos struct {
	NacosIp    string `mapstructure:"nacosIp" json:"nacosIp" yaml:"nacosIp"`
	NacosPort  uint64 `mapstructure:"nacosPort" json:"nacosPort" yaml:"nacosPort"`
	AppIp      string `mapstructure:"appIp" json:"appIp" yaml:"appIp"`
	AppPort    uint64 `mapstructure:"appPort" json:"appPort" yaml:"appPort"`
	ServerName string `mapstructure:"serverName" json:"serverName" yaml:"serverName"`
}


var configFilePath = "config"

func SetUp() {
	//读取环境变量
	args := os.Args

	//判断环境变量是否存在
	environment := os.Getenv("environment")
	if len(args) > 1 {
		configFilePath = configFilePath + "." + args[1]
	} else if environment != "" {
		configFilePath = configFilePath + "." + environment
	}

	configFilePath = configFilePath + ".yaml"

	//判断配置文件是否已存在 不存在则去环境变量中读取
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		config, err := nacos.GetConfig("gdy")
		if err != nil {
			fmt.Println(err)
			return
		}
		f, er := os.OpenFile(configFilePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0600)

		defer f.Close()
		if er == nil {
			f.Write([]byte(config))
		}
	}

	v := viper.New()
	v.SetConfigFile(configFilePath)
	err := v.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&GVA_CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&GVA_CONFIG); err != nil {
		fmt.Println(err)
	}
	GVA_VP = v
}

