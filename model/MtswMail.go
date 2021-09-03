package model

import (
	"mtsw/global"
	"time"
)

type MtswMail struct {
	Id                   int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Title                string `gorm:"column:title;type:varchar(20);NOT NULL" json:"title"` // 商城名称
	Aid                  int    `gorm:"column:aid;type:int(11);default:0;NOT NULL" json:"aid"`
	Uin                  int    `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Intro                string `gorm:"column:intro;type:varchar(100);NOT NULL" json:"intro"`                             // 商城简介
	ShareImg             string `gorm:"column:shareImg;type:varchar(200);NOT NULL" json:"shareImg"`                       // 分享图片
	ContactNumber        string `gorm:"column:contactNumber;type:varchar(16);NOT NULL" json:"contactNumber"`              // 联系电话
	ContactNumberDefault int    `gorm:"column:contactNumberDefault;type:tinyint(1);NOT NULL" json:"contactNumberDefault"` // 联系电话是否设为默认
	QrCode               string `gorm:"column:qrCode;type:varchar(100);NOT NULL" json:"qrCode"`                           // 客户二维码
	QrCodeDefault        int    `gorm:"column:qrCodeDefault;type:tinyint(1);NOT NULL" json:"qrCodeDefault"`               // 二维码是否设为默认
	CreateTime           int64    `gorm:"column:createTime;type:int(11);default:0;NOT NULL" json:"createTime"`
	UpdateTime           int64    `gorm:"column:updateTime;type:int(11);default:0;NOT NULL" json:"updateTime"`
}

func (m *MtswMail) TableName() string {
	return "mtsw_mail"
}

func (m *MtswMail) Create() error {
	m.CreateTime = time.Now().Unix()
	err := global.GVA_DB.Model(&m).Create(&m).Error
	return err
}

func (m *MtswMail) Update(field ...string) error {
	m.UpdateTime = time.Now().Unix()
	sql := global.GVA_DB.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	err := sql.Where("id", m.Id).Updates(m).Error
	return err
}

func (m *MtswMail) GetInfo() error {
	sql := global.GVA_DB.Model(m).Where("id = ? ", m.Id)
	err := sql.First(&m).Error
	return err
}

func (m *MtswMail) UpdateMail() error {
	err := m.Update("title", "intro", "shareImg", "contactNumber", "contactNumberDefault", "qrCode", "qrCodeDefault")
	return err
}

func GetMailInfo(uin, aid int) (MtswMail, error) {
	info := MtswMail{}
	sql := global.GVA_DB.Model(MtswMail{}).Where("uin = ? AND aid = ? ", uin, aid)
	err := sql.First(&info).Error
	return info, err
}
