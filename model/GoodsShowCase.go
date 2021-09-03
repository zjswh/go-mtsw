package model

import (
	"gorm.io/gorm"
	"time"
)

type GoodsShowCase struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin        int    `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Title      string `gorm:"column:title;type:varchar(10);NOT NULL" json:"title"`            // 橱窗标题
	Aid        int    `gorm:"column:aid;type:int(11);default:0;NOT NULL" json:"aid"`          // 频道标识
	Status     int    `gorm:"column:status;type:tinyint(2);default:1;NOT NULL" json:"status"` // 1正常0删除
	CreateTime int64  `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime int64  `gorm:"column:updateTime;type:int(11);default:0;NOT NULL" json:"updateTime"`
}

func (m *GoodsShowCase) TableName() string {
	return "goods_show_case"
}

func (m *GoodsShowCase) Create(Db *gorm.DB) error {
	m.CreateTime = time.Now().Unix()
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *GoodsShowCase) Update(Db *gorm.DB, field ...string) error {
	m.UpdateTime = time.Now().Unix()
	sql := Db.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	err := sql.Where("id", m.Id).Updates(m).Error
	return err
}

func (m *GoodsShowCase) GetInfo(Db *gorm.DB) error {
	sql := Db.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func (m *GoodsShowCase) GetList(Db *gorm.DB) ([]GoodsShowCase, error) {
	list := []GoodsShowCase{}
	err := Db.Model(m).Where("uin = ? AND aid = ? AND status = 1", m.Uin, m.Aid).Order("id desc").Find(&list).Error
	return list, err
}

func (m *GoodsShowCase) GetCount(Db *gorm.DB) (int64, error) {
	var count int64
	err := Db.Model(m).Where("uin = ? AND aid = ? AND status = 1", m.Uin, m.Aid).Count(&count).Error
	return count, err
}
