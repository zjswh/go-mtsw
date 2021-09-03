package model

import "gorm.io/gorm"

type Orders struct {
	Id              int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin             int     `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Aid             int     `gorm:"column:aid;type:int(11);default:0" json:"aid"`
	TransactionId   string  `gorm:"column:transactionId;type:varchar(40)" json:"transactionId"`
	OrderNo         string  `gorm:"column:orderNo;type:varchar(20)" json:"orderNo"`
	TrackingCompany string  `gorm:"column:trackingCompany;type:varchar(30)" json:"trackingCompany"` // 物流公司
	TrackingNo      string  `gorm:"column:trackingNo;type:varchar(30)" json:"trackingNo"`
	RefundId        string  `gorm:"column:refundId;type:varchar(20)" json:"refundId"`                   // 退款单号
	RefundReason    string  `gorm:"column:refundReason;type:varchar(150);NOT NULL" json:"refundReason"` // 退款原因
	RefundType      int     `gorm:"column:refundType;type:int(5);default:0;NOT NULL" json:"refundType"` // 退款类型
	UserId          int     `gorm:"column:userId;type:int(11)" json:"userId"`
	Status          int     `gorm:"column:status;type:int(11);default:0;NOT NULL" json:"status"` // 0->未支付 10->已支付 11->支付成功但是订单失败
	Num             int     `gorm:"column:num;type:int(11);default:0;NOT NULL" json:"num"`
	Price           float64 `gorm:"column:price;type:decimal(11,2)" json:"price"`
	PayType         int     `gorm:"column:payType;type:int(11);default:0;NOT NULL" json:"payType"` // 支付方式 1=>微信支付
	Code            string  `gorm:"column:code;type:varchar(30)" json:"code"`                      // 核销码
	CreateTime      int     `gorm:"column:createTime;type:int(11)" json:"createTime"`
	UpdateTime      int     `gorm:"column:updateTime;type:int(11)" json:"updateTime"`
	DeleteTime      int     `gorm:"column:deleteTime;type:int(11)" json:"deleteTime"`
}

func (m *Orders) TableName() string {
	return "orders"
}

func (m *Orders) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *Orders) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *Orders) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
