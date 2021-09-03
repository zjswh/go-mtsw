package model

import "gorm.io/gorm"

type MtswCategory struct {
	Id         int    `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin        int    `gorm:"column:uin;type:int(11);NOT NULL" json:"uin"`
	Aid        int    `gorm:"column:aid;type:int(11);default:0;NOT NULL" json:"aid"`
	Name       string `gorm:"column:name;type:varchar(10);NOT NULL" json:"name"`       // 分类名称
	Sort       int    `gorm:"column:sort;type:int(7);default:99;NOT NULL" json:"sort"` // 排序
	CreateTime int    `gorm:"column:createTime;type:int(11);default:0;NOT NULL" json:"createTime"`
	UpdateTime int    `gorm:"column:updateTime;type:int(11);default:0;NOT NULL" json:"updateTime"`
}

func (m *MtswCategory) TableName() string {
	return "mtsw_category"
}

func (m *MtswCategory) Create(Db *gorm.DB) error {
    err := Db.Model(&m).Create(&m).Error
    return err
}

func (m *MtswCategory) Update(Db *gorm.DB, field ...string) error {
    sql := Db.Model(&m)
    if len(field) > 0 {
        sql = sql.Select(field)
    }
    err := sql.Where("id", m.Id).Updates(m).Error
    return err
}

func (m *MtswCategory) GetInfo(Db *gorm.DB) error {
    sql := Db.Model(m).Where("id = ? ", m.Id)
    err := sql.First(&m).Error
    return err
}

func GetCategoryList(Db *gorm.DB, uin, aid int) ([]MtswCategory, error) {
	list := []MtswCategory{}
	err := Db.Model(&MtswCategory{}).Where("uin = ? AND aid = ?", uin, aid).Find(&list).Order("sort asc").Error
	return list, err
}
