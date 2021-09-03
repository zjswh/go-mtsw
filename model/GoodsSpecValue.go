package model

import (
	"gorm.io/gorm"
)

type GoodsSpecValue struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId    int    `gorm:"column:goodsId;type:int(11)" json:"goodsId"`
	SpecId     int    `gorm:"column:specId;type:int(11);NOT NULL" json:"specId"`
	Name       string `gorm:"column:name;type:varchar(500);NOT NULL" json:"name"`
	DeleteTime int    `gorm:"column:deleteTime;type:int(11);default:(-)" json:"deleteTime"`
}

func (m *GoodsSpecValue) TableName() string {
	return "goods_spec_value"
}

func CreateMultiSpecValue(Db *gorm.DB,list []GoodsSpecValue) error {
    err := Db.Model(&GoodsSpecValue{}).Create(&list).Error
    return err
}

func (m *GoodsSpecValue) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func DeleteSpecValue(Db *gorm.DB, goodsId int) error  {
	err := Db.Where("goodsId", goodsId).Delete(&GoodsSpecValue{}).Error
	return err
}

func (m *GoodsSpecValue) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
