package model

import "gorm.io/gorm"

type GoodsTag struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `gorm:"column:name;type:varchar(10);default:0;NOT NULL" json:"name"`
	Alias      string `gorm:"column:alias;type:varchar(10)" json:"alias"`
	Type       int    `gorm:"column:type;type:tinyint(3);default:1;NOT NULL" json:"type"` // 0为系统预设 1为客户自定义
	CreateTime int    `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime int    `gorm:"column:updateTime;type:int(11);default:0" json:"updateTime"`
}

func (m *GoodsTag) TableName() string {
	return "goods_tag"
}

func (m *GoodsTag) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *GoodsTag) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *GoodsTag) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
