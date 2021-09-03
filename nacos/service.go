package nacos

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"time"
)

var namingClient naming_client.INamingClient
var configClient config_client.IConfigClient

// ServiceSetup... 创建服务发现客户端
func ServiceSetup() {
	var err error
	namingClient, err = clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println("创建服务发现客户端失败: ", err.Error())
	}

	configClient, err = clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	if err != nil {
		fmt.Println("创建client失败: ", err.Error())
	}
}

// serviceRegister...注册实例
func serviceRegister(nacosIp string, nacosPort uint64, appIp string, appPort uint64, serverName string) {
	fmt.Printf("\n----- NACOS REGISTER START -----\n")
	_, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          appIp,
		Port:        appPort,
		ServiceName: serverName,
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		// Metadata:    map[string]string{"idc":"shanghai"},
		// ClusterName: "cluster-a", // 默认值DEFAULT
		// GroupName:   "group-a",  // 默认值DEFAULT_GROUP
	})
	fmt.Printf("NACOS IP: %s \n", nacosIp)
	fmt.Printf("NACOS PORT: %d \n", nacosPort)
	fmt.Printf("MINE IP: %s \n", appIp)
	fmt.Printf("MINE PORT: %d \n", appPort)
	fmt.Printf("MINE servername: %s \n", serverName)

	if err != nil {
		fmt.Printf("ERR: %s\n", err.Error())
		fmt.Printf("----- NACOS REGISTER END -----\n")
		fmt.Println("WAIT TO RETRY IN 30 SECONDS...")
		time.Sleep(30 * time.Second)
		serviceRegister(nacosIp, nacosPort, appIp, appPort, serverName)
	} else {
		fmt.Println("SUCCESS !!!")
		fmt.Printf("----- NACOS REGISTER END -----\n")
	}
}

// ServiceGetInfo... 获取服务信息
func ServiceGetInfo(serviceName string) (services model.Service, err error) {
	services, err = namingClient.GetService(vo.GetServiceParam{
		ServiceName: serviceName,
		// Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		// GroupName:   "group-a",             // 默认值DEFAULT_GROUP
	})
	if err != nil {
		err = fmt.Errorf("获取服务信息失败: %s", err.Error())
		fmt.Println(err.Error())
		return
	}

	return
}

func GetConfig(dataId string) (config string, err error) {
	config, err = configClient.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group: "DEFAULT_GROUP",
	})
	return
}
