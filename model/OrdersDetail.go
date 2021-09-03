package model

import "gorm.io/gorm"

type OrdersDetail struct {
	Id          int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	GoodsId     int     `gorm:"column:goodsId;type:int(11)" json:"goodsId"`
	GoodsImg    string  `gorm:"column:goodsImg;type:varchar(100)" json:"goodsImg"`
	GoodsName   string  `gorm:"column:goodsName;type:varchar(100)" json:"goodsName"`
	GoodsPrice  float64 `gorm:"column:goodsPrice;type:decimal(11,2)" json:"goodsPrice"`
	Freight     float64 `gorm:"column:freight;type:decimal(10,2);default:0.00" json:"freight"` // 运费
	OrderId     int     `gorm:"column:orderId;type:int(11)" json:"orderId"`
	OrderMsg    string  `gorm:"column:orderMsg;type:varchar(255)" json:"orderMsg"` // 买家留言
	SpecIds     string  `gorm:"column:specIds;type:varchar(40)" json:"specIds"`
	NewSpecId   int     `gorm:"column:newSpecId;type:int(11)" json:"newSpecId"`
	SpecValues  string  `gorm:"column:specValues;type:varchar(500)" json:"specValues"`
	Name        string  `gorm:"column:name;type:varchar(200)" json:"name"`
	Mobile      string  `gorm:"column:mobile;type:varchar(11)" json:"mobile"`
	Province    string  `gorm:"column:province;type:varchar(100)" json:"province"`
	City        string  `gorm:"column:city;type:varchar(100)" json:"city"`
	Area        string  `gorm:"column:area;type:varchar(100)" json:"area"`
	Detail      string  `gorm:"column:detail;type:varchar(400)" json:"detail"`
	FreightType int     `gorm:"column:freightType;type:int(4);default:1" json:"freightType"` // 配送方式
}

func (m *OrdersDetail) TableName() string {
	return "orders_detail"
}

func (m *OrdersDetail) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *OrdersDetail) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *OrdersDetail) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
