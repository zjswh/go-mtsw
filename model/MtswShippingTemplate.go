package model

import (
"gorm.io/gorm"

	"time"
)

type MtswShippingTemplate struct {
	Id          int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin         int       `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Aid         int       `gorm:"column:aid;type:int(11);default:0;NOT NULL" json:"aid"`
	Title       string    `gorm:"column:title;type:varchar(10);NOT NULL" json:"title"`
	Type        int       `gorm:"column:type;type:tinyint(1);NOT NULL" json:"type"` // 1包邮 0不包邮
	BaseConfig  string    `gorm:"column:baseConfig;type:text" json:"baseConfig"`    // 配置
	ExtraConfig string    `gorm:"column:extraConfig;type:text" json:"extraConfig"`
	LimitArea   string    `gorm:"column:limitArea;type:text" json:"limitArea"` // 限制购买地区
	Deleted     int       `gorm:"column:deleted;type:tinyint(1);default:0;NOT NULL" json:"deleted"`
	CreateTime  time.Time `gorm:"column:createTime;type:datetime" json:"createTime"`
	UpdateTime  time.Time `gorm:"column:updateTime;type:datetime;NOT NULL" json:"updateTime"`
}

func (m *MtswShippingTemplate) TableName() string {
	return "mtsw_shipping_template"
}

func (m *MtswShippingTemplate) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *MtswShippingTemplate) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *MtswShippingTemplate) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
