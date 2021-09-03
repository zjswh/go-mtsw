package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"mtsw/config"
)

var (
	GVA_CONFIG config.Server
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
