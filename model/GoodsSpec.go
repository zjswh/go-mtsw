package model

import (
	"gorm.io/gorm"
)

type GoodsSpec struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId    int    `gorm:"column:goodsId;type:int(11);NOT NULL" json:"goodsId"` // 商品id
	Name       string `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	DeleteTime int    `gorm:"column:deleteTime;type:int(11);default:(-)" json:"deleteTime"`
}

type GoodsSpecList struct {
	Id        int    `gorm:"column:id" json:"id"`
	GoodsId   int    `gorm:"column:goodsId" json:"goodsId"`
	Name      string `gorm:"column:name" json:"name"`
	SpecId    int    `gorm:"column:specId" json:"specId"`
	ValueName string `gorm:"column:valueName" json:"valueName"`
	ValueId   int    `gorm:"column:valueId" json:"valueId"`
}

func (m *GoodsSpec) TableName() string {
	return "goods_spec"
}

func (m *GoodsSpec) Create(Db *gorm.DB) error {
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *GoodsSpec) Update(Db *gorm.DB, field ...string) error {
	sql := Db.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	err := sql.Where("id", m.Id).Updates(m).Error
	return err
}

func (m *GoodsSpec) GetInfo(Db *gorm.DB) error {
	sql := Db.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func DeleteSpec(Db *gorm.DB, goodsId int) error  {
	err := Db.Where("goodsId", goodsId).Delete(&GoodsSpec{}).Error
	return err
}

func GetSpecList(Db *gorm.DB, goodsId int) ([]GoodsSpecList, error) {
	list := []GoodsSpecList{}
	goodsSpecModel := GoodsSpec{}
	goodsSpecValueModel := GoodsSpecValue{}

	goodsSpecTable := goodsSpecModel.TableName() + " l"
	goodsSpecValueTable := goodsSpecValueModel.TableName() + " c"
	err := Db.Table(goodsSpecTable).
		Joins("LEFT JOIN "+goodsSpecValueTable+" on c.specId=l.id").
		Where("l.goodsId = ? AND l.deleteTime is NULL AND c.deleteTime is NULL", goodsId).
		Select("l.id,l.name,l.goodsId,c.name as valueName, c.id as valueId, c.specId").
		Scan(&list).Error
	return list, err
}
