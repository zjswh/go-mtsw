package model

import (
"gorm.io/gorm"

	"time"
)

type OrdersRefundRecord struct {
	Id         int       `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin        int       `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	OrderId    int       `gorm:"column:orderId;type:int(11);NOT NULL" json:"orderId"`
	Reason     string    `gorm:"column:reason;type:varchar(150);NOT NULL" json:"reason"` // 退款原因
	Type       int       `gorm:"column:type;type:int(5);NOT NULL" json:"type"`           // 退款类型
	CreateTime time.Time `gorm:"column:createTime;type:datetime;NOT NULL" json:"createTime"`
}

func (m *OrdersRefundRecord) TableName() string {
	return "orders_refund_record"
}

func (m *OrdersRefundRecord) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *OrdersRefundRecord) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *OrdersRefundRecord) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
