package model

import (
	"gorm.io/gorm"
	"time"
)

type Goods struct {
	Id                int     `gorm:"column:id;type:int(11);primary_key;AUTO_INCREMENT" json:"id"`
	Uin               int     `gorm:"column:uin;type:int(11)" json:"uin"`
	Aid               int     `gorm:"column:aid;type:int(11)" json:"aid"`
	Name              string  `gorm:"column:name;type:varchar(100);NOT NULL" json:"name"`
	Desc              string  `gorm:"column:desc;type:varchar(500)" json:"desc"` // 商品描述
	Price             float64 `gorm:"column:price;type:decimal(6,2)" json:"price"`
	ShopPrice         float64 `gorm:"column:shop_price;type:decimal(11,2)" json:"shop_price"` // 划线价格
	Stock             int     `gorm:"column:stock;type:int(11);default:0" json:"stock"`       // 库存
	Sum               int     `gorm:"column:sum;type:int(11);default:0" json:"sum"`           // 总数
	Source            int     `gorm:"column:source;type:int(4);default:1" json:"source"`      // 商品来源 1链接2口令
	SourceContent     string  `gorm:"column:sourceContent;type:text" json:"sourceContent"`    // 来源内容
	Context           string  `gorm:"column:context;type:text" json:"context"`
	LimitNum          int     `gorm:"column:limitNum;type:int(11);default:0" json:"limitNum"`                   // 限购数
	OrderNum          int     `gorm:"column:orderNum;type:int(11)" json:"orderNum"`                             // 排序
	Type              int     `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                        // 1 =>基础类型2外链
	Freight           float64 `gorm:"column:freight;type:decimal(10,2);default:0.00" json:"freight"`            // 运费
	FreightTemplateId int     `gorm:"column:freightTemplateId;type:int(11);default:0" json:"freightTemplateId"` // 运费模板
	FreightType       int     `gorm:"column:freightType;type:tinyint(255);default:3" json:"freightType"`        // 1直付2模板3包邮
	PickUpAddress     string  `gorm:"column:pickUpAddress;type:varchar(60)" json:"pickUpAddress"`               // 自提地址
	PickUpDeadline    int     `gorm:"column:pickUpDeadline;type:int(4);default:1" json:"pickUpDeadline"`        // 自提期限
	ContactNumber     string  `gorm:"column:contactNumber;type:varchar(20)" json:"contactNumber"`               // 联系电话
	QrCode            string  `gorm:"column:qrCode;type:varchar(150)" json:"qrCode"`                            // 客服二维码
	IsTop             int     `gorm:"column:isTop;type:tinyint(1);default:0" json:"isTop"`                      // 1->推荐商品
	Status            int     `gorm:"column:status;type:int(11);default:0" json:"status"`                       // 1->正常上架
	CreateTime        int64   `gorm:"column:createTime;type:int(11)" json:"createTime"`
	UpdateTime        int64   `gorm:"column:updateTime;type:int(11)" json:"updateTime"`
	DeleteTime        int64   `gorm:"column:deleteTime;type:int(11);default:(-)" json:"deleteTime"`
	CategoryId        int     `gorm:"column:categoryId;type:int(11);default:0;NOT NULL" json:"categoryId"` // 分类id
	LaunchTime        int64   `gorm:"column:launchTime;type:int(11);default:0;NOT NULL" json:"launchTime"` // 上架时间
}
type GoodsWithCaseId struct {
	Id     int    `gorm:"column:id" json:"id"`
	Name   string `gorm:"column:name" json:"name"`
	CaseId int    `gorm:"column:caseId" json:"caseId"`
}

func (m *Goods) TableName() string {
	return "goods"
}

func (m *Goods) Create(Db *gorm.DB) error {
	m.CreateTime = time.Now().Unix()
	err := Db.Model(&m).Create(&m).Error
	return err
}

func (m *Goods) UpdateGoods(Db *gorm.DB) error {
	m.UpdateTime = time.Now().Unix()
	field := []string{"name", "type", "desc", "shopPrice", "context", "freight", "freightType", "freightTemplateId",
		"limitNum", "contactNumber", "qrCode", "pickUpDeadline", "categoryId", "pickUpAddress", "source", "sourceContent", "updateTime"}
	return m.Update(Db, field)
}

func (m *Goods) Update(Db *gorm.DB, field []string) error {
	sql := Db.Model(&m)
	if len(field) > 0 {
		sql = sql.Select(field)
	}
	if m.Uin > 0 {
		sql = sql.Where("uin", m.Uin)
	}

	err := sql.Where("id", m.Id).Updates(m).Error
	return err
}

func (m *Goods) GetInfo(Db *gorm.DB) error {
	err := Db.Model(m).Where("id = ? AND deleteTime is NULL", m.Id).First(&m).Error
	return err
}

func (m *Goods) GetList(Db *gorm.DB, goodsName string, categoryId, goodsType, status, sTime, eTime, page, num int) ([]Goods, error) {
	list := []Goods{}
	sql := Db.Model(&m).Where("uin = ? AND aid = ? AND deleteTime is NULL", m.Uin, m.Aid)
	if goodsName != "" {
		sql = sql.Where("name LIKE ?", "%"+goodsName+"%")
	}
	if categoryId > 0 {
		sql = sql.Where("categoryId", categoryId)
	}
	if status != -1 {
		sql = sql.Where("status", status)
	}
	if goodsType > 0 {
		sql = sql.Where("type", goodsType)
	}
	if sTime > 0 && eTime > 0 {
		sql = sql.Where("createTime BETWEEN ? AND ?", sTime, eTime)
	}
	err := sql.Order("id desc").Limit(num).Offset((page - 1) * num).Find(&list).Error
	return list, err
}

func (m *Goods) GetCount(Db *gorm.DB, goodsName string, categoryId, goodsType, status, sTime, eTime int) (int64, error) {
	var count int64
	sql := Db.Model(&m).Where("uin = ?  AND deleteTime is NULL", m.Uin)
	if goodsName != "" {
		sql = sql.Where("name LIKE ?", "%"+goodsName+"%")
	}
	if categoryId > 0 {
		sql = sql.Where("categoryId", categoryId)
	}
	if status != -1 {
		sql = sql.Where("status", status)
	}
	if goodsType > 0 {
		sql = sql.Where("type", goodsType)
	}
	if sTime > 0 && eTime > 0 {
		sql = sql.Where("createTime BETWEEN ? AND ?", sTime, eTime)
	}
	err := sql.Order("id desc").Count(&count).Error
	return count, err
}

func (m *Goods) GetCountWithCase(Db *gorm.DB) (int64, error) {
	return m.GetCount(Db, "", m.CategoryId, 0, m.Status, 0,0)
}

func (m *Goods) GetListWithCase(Db *gorm.DB, caseId, page, num int) ([]GoodsWithCaseId, error) {
	list := []GoodsWithCaseId{}
	goodsModel := Goods{}
	caseModel := ShowcaseBindGoods{}
	caseTable := caseModel.TableName() + " c"
	goodsTable := goodsModel.TableName() + " g"
	sql := Db.Table(goodsTable).
		Joins("LEFT JOIN "+caseTable+" ON c.goodsId = g.id AND c.deleted = 0 AND c.caseId = ? ", caseId).
		Where("g.uin = ? AND g.status = 1 AND g.deleteTime is NULL", m.Uin)
	if m.CategoryId > 0 {
		sql = sql.Where("categoryId", m.CategoryId)
	}

	err := sql.Limit(num).
		Offset((page - 1) * num).
		Select("g.id,g.name,c.caseId").
		Order("c.sort desc, c.id desc").
		Scan(&list).
		Error
	return list, err
}

func (m *Goods) Delete(Db *gorm.DB) error {
	m.DeleteTime = time.Now().Unix()
	field := []string{"deleteTime"}
	return m.Update(Db, field)
}

func (m *Goods) UpdateStatus(Db *gorm.DB) error {
	m.UpdateTime = time.Now().Unix()
	field := []string{"updateTime", "status"}
	return m.Update(Db, field)
}
