package model

import (
	"gorm.io/gorm"
)

type GoodsImgs struct {
	Id      int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId int    `gorm:"column:goodsId;type:int(11)" json:"goodsId"`
	Img     string `gorm:"column:img;type:varchar(250)" json:"img"`
}

func (m *GoodsImgs) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func CreateImgMulti(Db *gorm.DB, list []GoodsImgs) error {
	err := Db.Model(&GoodsImgs{}).Create(&list).Error
	return err
}


func (m *GoodsImgs) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *GoodsImgs) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetImgList(Db *gorm.DB, goodsId int) ([]GoodsImgs, error) {
	list := []GoodsImgs{}
	err := Db.Model(&GoodsImgs{}).Where("goodsId = ? ", goodsId).Find(&list).Order("id desc").Error
	return list, err
}

func GetImgListByGoodsIdArr(Db *gorm.DB, goodsIdArr []int) ([]GoodsImgs, error) {
	list := []GoodsImgs{}
	err := Db.Model(&GoodsImgs{}).Where("goodsId IN ? ", goodsIdArr).Group("goodsId").
		Order("id asc").Find(&list).Error
	return list, err
}

func DeleteImg(Db *gorm.DB, list []GoodsImgs) error {
	idArr := []int{}
	for _ , v := range list {
		idArr = append(idArr, v.Id)
	}
	err := Db.Where("id IN ?", idArr).Delete(&GoodsImgs{}).Error
	return err
}
