package model

import "gorm.io/gorm"

type ShowcaseInclude struct {
	Id          int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Include     string `gorm:"column:include;type:varchar(30);NOT NULL" json:"include"`          // 引用id
	IncludeName string `gorm:"column:includeName;type:varchar(180);NOT NULL" json:"includeName"` // 引用名称
	CaseId      int    `gorm:"column:caseId;type:int(11);NOT NULL" json:"caseId"`                // 橱窗id
	Status      int    `gorm:"column:status;type:tinyint(2);default:0;NOT NULL" json:"status"`   // 是否开启 1开0关
	CreateTime  int    `gorm:"column:createTime;type:int(11);NOT NULL" json:"createTime"`
	UpdateTime  int    `gorm:"column:updateTime;type:int(11);default:0;NOT NULL" json:"updateTime"`
}

func (m *ShowcaseInclude) TableName() string {
	return "showcase_include"
}

func (m *ShowcaseInclude) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *ShowcaseInclude) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *ShowcaseInclude) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}
