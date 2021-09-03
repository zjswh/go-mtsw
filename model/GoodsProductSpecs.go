package model

import (
	"gorm.io/gorm"
	"time"
)

type GoodsProductSpecs struct {
	Id         int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId    int     `gorm:"column:goodsId;type:int(11);NOT NULL" json:"goodsId"`
	Specs      string  `gorm:"column:specs;type:varchar(150)" json:"specs"`           // 规格
	Stock      int     `gorm:"column:stock;type:int(11);NOT NULL" json:"stock"`       // 数量
	Sum        int     `gorm:"column:sum;type:int(11);default:0;NOT NULL" json:"sum"` // 总数
	Price      float64 `gorm:"column:price;type:decimal(11,2)" json:"price"`          // 价格
	Img        string  `gorm:"column:img;type:varchar(200)" json:"img"`               // 图片
	CreateTime int64     `gorm:"column:createTime;type:int(10);NOT NULL" json:"createTime"`
	UpdateTime int64     `gorm:"column:updateTime;type:int(10)" json:"updateTime"`
}

func CreateProduceSpecsMulti(Db *gorm.DB, list []GoodsProductSpecs) error {
    err := Db.Model(&GoodsProductSpecs{}).Create(&list).Error
    return err
}

func(m *GoodsProductSpecs) Create(Db *gorm.DB) error {
	m.CreateTime = time.Now().Unix()
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *GoodsProductSpecs) Update(Db *gorm.DB, field ...string) error {
	m.UpdateTime = time.Now().Unix()
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *GoodsProductSpecs) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetProductList(Db *gorm.DB, goodsId int) ([]GoodsProductSpecs, error) {
	list := []GoodsProductSpecs{}
	err := Db.Model(&GoodsProductSpecs{}).Where("goodsId", goodsId).Find(&list).Error
	return list, err
}

func DeleteProductSpec(Db *gorm.DB, idArr []int) error {
	err := Db.Where("id IN ?", idArr).Delete(&GoodsProductSpecs{}).Error
	return err
}

func DeleteProductByGoodsId(Db *gorm.DB, goodsId int) error {
	err := Db.Where("goodsId", goodsId).Delete(&GoodsProductSpecs{}).Error
	return err
}
