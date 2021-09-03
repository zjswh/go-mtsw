package model

import "gorm.io/gorm"

type GoodsBindTag struct {
	Id      int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId int `gorm:"column:goodsId;type:int(11);NOT NULL" json:"goodsId"` // 商品id
	TagId   int `gorm:"column:tagId;type:int(11);NOT NULL" json:"tagId"`     // 标签id
}

func (m *GoodsBindTag) TableName() string {
	return "goods_bind_tag"
}

func (m *GoodsBindTag) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *GoodsBindTag) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *GoodsBindTag) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
