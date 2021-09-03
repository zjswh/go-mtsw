package model

import "gorm.io/gorm"

type MenuGoods struct {
	Id        int `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	IncludeId int `gorm:"column:include_id;type:int(11);default:0" json:"include_id"`
	MenuId    int `gorm:"column:menu_id;type:int(11);NOT NULL" json:"menu_id"`
	GoodsId   int `gorm:"column:goods_id;type:int(11);NOT NULL" json:"goods_id"`
	Sort      int `gorm:"column:sort;type:int(5);default:99;NOT NULL" json:"sort"`      // 排序
	Type      int `gorm:"column:type;type:int(3);default:1;NOT NULL" json:"type"`       // 1为菜单下 2为购物袋
	IsTop     int `gorm:"column:isTop;type:tinyint(1);default:0;NOT NULL" json:"isTop"` // 是否置顶
	OldSort   int `gorm:"column:oldSort;type:int(5);default:0" json:"oldSort"`
}

func (m *MenuGoods) TableName() string {
	return "menu_goods"
}

func (m *MenuGoods) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *MenuGoods) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *MenuGoods) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
