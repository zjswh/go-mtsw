package main

import (
	"mtsw/config"
	"mtsw/core"
	"mtsw/global"
	"mtsw/initialize"
	"mtsw/nacos"
)

func main() {
	nacosIp, nacosPort,appIp, appPort, serverName := "47.110.242.138", uint64(8848), "127.0.0.1", uint64(8002), "go-mtsw"
	nacos.Setup(nacosIp, nacosPort,appIp, appPort, serverName)

	//加载config配置
	config.SetUp()

	//加载数据库
	switch global.GVA_CONFIG.System.DbType {
		case "mysql":
			initialize.Mysql()
		default:
			initialize.Mysql()
	}

	//注册数据库表
	initialize.DBTables()

	//加载redis
	initialize.Redis()

	core.RunServer()
}
