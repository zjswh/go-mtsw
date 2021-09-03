package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/common/constant"
)

var (
	serverConfigs []constant.ServerConfig

	clientConfig = constant.ClientConfig{}
)

func Setup(nacosIp string, nacosPort uint64, appIp string, appPort uint64, serverName string) {
	serverConfigs = []constant.ServerConfig{
		{
			IpAddr: nacosIp,
			Port:   nacosPort,
		},
	}

	clientConfig = constant.ClientConfig{
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		RotateTime:          "1h",
		MaxAge:              3,
		LogLevel:            "debug",
	}
	ServiceSetup()
	serviceRegister(nacosIp, nacosPort, appIp, appPort, serverName)
}
