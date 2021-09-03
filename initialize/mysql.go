package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"mtsw/global"
	"os"
)

func Mysql() {
	config := global.GVA_CONFIG.Mysql
	link := config.Username + ":" + config.Password + "@(" + config.Path + ")/" + config.Dbname + "?" + config.Config
	if db, err := gorm.Open(mysql.Open(link), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,                              // 使用单数表名，启用该选项后，`User` 表将是`user`
		},
		Logger: logger.Default.LogMode(logger.Info),
	}); err != nil {
		fmt.Println("mysql connect failed", err.Error())
		os.Exit(0)
	} else {

		global.GVA_DB = db
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(config.MaxIdleConns)
		sqlDb.SetMaxOpenConns(config.MaxOpenConns)
		//global.GVA_DB.SingularTable(true)
	}

}

func DBTables() {
	//global.GVA_DB.AutoMigrate(
	//	model.User{},
	//	)
}
