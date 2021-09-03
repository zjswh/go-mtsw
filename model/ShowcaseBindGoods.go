package model

import (
	"gorm.io/gorm"
	"time"
)

type ShowcaseBindGoods struct {
	Id           int   `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	CaseId       int   `gorm:"column:caseId;type:int(11);NOT NULL" json:"caseId"`                          // 橱窗id
	GoodsId      int   `gorm:"column:goodsId;type:int(11);NOT NULL" json:"goodsId"`                        // 商品id
	Deleted      int   `gorm:"column:deleted;type:tinyint(2);default:0;NOT NULL" json:"deleted"`           // 是否移除 1移除
	IsTop        int   `gorm:"column:isTop;type:tinyint(2);default:0;NOT NULL" json:"isTop"`               // 是否置顶 1置顶
	IsExplaining int   `gorm:"column:isExplaining;type:tinyint(2);default:0;NOT NULL" json:"isExplaining"` // 是否讲解中 1是
	Sort         int   `gorm:"column:sort;type:int(10);NOT NULL" json:"sort"`                              // 排序
	CreateTime   int64 `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime   int64 `gorm:"column:updateTime;type:int(11);default:0;NOT NULL" json:"updateTime"`
}

type CaseGoods struct {
	ShowcaseBindGoods
	Name string `json:"name"`
}

func (m *ShowcaseBindGoods) TableName() string {
	return "showcase_bind_goods"
}

func (m *ShowcaseBindGoods) Create(Db *gorm.DB) error {
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *ShowcaseBindGoods) Update(Db *gorm.DB, field ...string) error {
	m.UpdateTime = time.Now().Unix()
	sql := Db.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	err := sql.Where("caseId = ? AND goodsId = ?", m.CaseId, m.GoodsId).Updates(m).Error
	return err
}

func (m *ShowcaseBindGoods) ClearTopStatus(Db *gorm.DB) error {
	m.IsTop = 0
	err := Db.Model(&m).Select("isTop").Where("caseId = ?", m.CaseId).Updates(m).Error
	return err
}

func (m *ShowcaseBindGoods) ClearExplainingStatus(Db *gorm.DB) error {
	m.IsExplaining = 0
	err := Db.Model(&m).Select("isExplaining").Where("caseId = ?", m.CaseId).Updates(m).Error
	return err
}

func (m *ShowcaseBindGoods) GetInfo(Db *gorm.DB) error {
	sql := Db.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func (m *ShowcaseBindGoods) GetGoodsList(Db *gorm.DB, page, num int) ([]CaseGoods, error) {
	list := []CaseGoods{}
	caseModel := ShowcaseBindGoods{}
	goodsModel := Goods{}
	caseTable := caseModel.TableName() + " c"
	goodsTable := goodsModel.TableName() + " g"
	err := Db.Table(caseTable).
		Joins("LEFT JOIN "+goodsTable+" ON c.goodsId = g.id").
		Where("c.caseId = ? AND c.deleted = 0 AND g.deleteTime is NULL", m.CaseId).
		Limit(num).
		Offset((page - 1) * num).
		Select("c.*, g.name").
		Order("c.sort desc, c.id desc").
		Scan(&list).
		Error
	return list, err
}

func GetCaseGoodsByCaseIdArr(Db *gorm.DB, caseIdArr []int) ([]ShowcaseBindGoods, error) {
	list := []ShowcaseBindGoods{}
	err := Db.Model(ShowcaseBindGoods{}).Where("id IN (?) ", Db.Model(ShowcaseBindGoods{}).Where("caseId IN ? AND deleted = 0 ", caseIdArr).
		Group("caseId").Select("max(id)")).Find(&list).Error
	return list, err
}

func (m *ShowcaseBindGoods) GetGoodsCount(Db *gorm.DB) (int64, error) {
	var count int64
	caseModel := ShowcaseBindGoods{}
	goodsModel := Goods{}
	caseTable := caseModel.TableName() + " c"
	goodsTable := goodsModel.TableName() + " g"
	err := Db.Table(caseTable).
		Joins("LEFT JOIN "+goodsTable+" ON c.goodsId = g.id").
		Where("c.caseId = ? AND c.deleted = 0 AND g.deleteTime is NULL", m.CaseId).
		Count(&count).Error
	return count, err
}
