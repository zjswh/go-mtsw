package goodsService

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"html"
	"mtsw/config"
	"mtsw/model"
	"mtsw/types"
	"mtsw/utils"
	"mtsw/utils/appConst"
	"github.com/zjswh/go-tool/ossService"
	"strconv"
	"strings"
)

const module = "program"

type CategoryList struct {
	Id   int    `json:"id"`
	Uin  int    `json:"uin"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
}

type GoodsMsg struct {
	ID                int                           `json:"id"`
	Uin               int                           `json:"uin"`
	Aid               int                           `json:"aid"`
	Name              string                        `json:"name"`
	Desc              string                        `json:"desc"`
	Price             interface{}                   `json:"price"`
	ShopPrice         string                        `json:"shop_price"`
	Stock             int                           `json:"stock"`
	Sum               int                           `json:"sum"`
	Source            int                           `json:"source"`
	SourceContent     string                        `json:"sourceContent"`
	Context           string                        `json:"context"`
	LimitNum          interface{}                   `json:"limitNum"`
	OrderNum          int                           `json:"orderNum"`
	Type              int                           `json:"type"`
	Freight           string                        `json:"freight"`
	FreightTemplateID int                           `json:"freightTemplateId"`
	FreightType       int                           `json:"freightType"`
	PickUpAddress     string                        `json:"pickUpAddress"`
	PickUpDeadline    int                           `json:"pickUpDeadline"`
	ContactNumber     string                        `json:"contactNumber"`
	QrCode            string                        `json:"qrCode"`
	IsTop             int                           `json:"isTop"`
	Status            int                           `json:"status"`
	CreateTime        int64                         `json:"createTime"`
	UpdateTime        int64                         `json:"updateTime"`
	CategoryID        int                           `json:"categoryId"`
	LaunchTime        int64                         `json:"launchTime"`
	Imgs              []string                      `json:"imgs"`
	Spec              []types.GoodsSpec             `json:"spec"`
	SpecInfo          []types.ResponseGoodsSpecInfo `json:"specInfo"`
}

func GetMailInfo(uin, aid int) (model.MtswMail, error) {
	info, err := model.GetMailInfo(config.GVA_DB, uin, aid)
	return info, err
}

func SetMailInfo(req types.SetMailInfoStruct, uin, aid int) error {
	info, _ := GetMailInfo(uin, aid)
	isNew := true
	if info.Id != 0 {
		isNew = false
	}

	contactNumberDefault := 0
	if req.ContactNumberDefault != 0 {
		contactNumberDefault = 1
	}

	qrCodeDefault := 0
	if req.QrCodeDefault != 0 {
		qrCodeDefault = 1
	}

	mail := model.MtswMail{
		Title:                req.Title,
		Intro:                req.Intro,
		ShareImg:             req.ShareImg,
		ContactNumber:        req.ContactNumber,
		ContactNumberDefault: contactNumberDefault,
		QrCode:               req.QrCode,
		QrCodeDefault:        qrCodeDefault,
	}

	var err error
	if isNew == true {
		mail.Uin = uin
		mail.Aid = aid
		err = mail.Create(config.GVA_DB)
	} else {
		mail.Id = info.Id
		err = mail.UpdateMail(config.GVA_DB)
	}

	//????????????
	config.GVA_REDIS.Del(appConst.GetMailInfoKey(uin, aid))
	return err
}

func CreateCategory(name string, uin, aid int) (int, error) {
	category := model.MtswCategory{
		Name: name,
		Uin:  uin,
		Aid:  aid,
		Sort: 99,
	}
	err := category.Create(config.GVA_DB)
	return category.Id, err
}

func UpdateCategory(name string, categoryId, uin, aid int) error {
	category := model.MtswCategory{
		Id:   categoryId,
		Uin:  uin,
		Aid:  aid,
		Name: name,
	}
	err := category.GetInfo(config.GVA_DB)
	if err != nil {
		return err
	}

	category = model.MtswCategory{
		Id:   categoryId,
		Name: name,
	}
	err = category.Update(config.GVA_DB, "name")
	return err
}

func SaveCategories(categoryIds string) error {
	idArr := strings.Split(categoryIds, ",")
	for k, v := range idArr {
		id, _ := strconv.Atoi(v)
		info := model.MtswCategory{
			Id:   id,
			Sort: k + 1,
		}
		info.Update(config.GVA_DB, "sort")
	}
	return nil
}

func GetCategories(uin, aid int) ([]CategoryList, error) {
	list := []CategoryList{}
	var err error
	tmpList, err := model.GetCategoryList(config.GVA_DB, uin, aid)
	for _, v := range tmpList {
		list = append(list, CategoryList{
			Id:   v.Id,
			Uin:  v.Uin,
			Name: v.Name,
			Sort: v.Sort,
		})
	}
	return list, err
}

func AddGoods(req types.GoodsStruct, uin, aid int) (int, error) {
	goods, specInfo, goodsSpec, err := dealGoods(req, uin, aid)
	if err != nil {
		return 0, err
	}
	//????????????
	db := config.GVA_DB.Begin()
	err = goods.Create(db)
	if err != nil {
		db.Rollback()
		return 0, err
	}
	//???????????????
	err = model.CreateImgMulti(db, dealImg(goods.Id, uin, req.GoodsImgs, []model.GoodsImgs{}))
	if err != nil {
		db.Rollback()
		return 0, err
	}

	//????????????
	err = addSpec(db, goodsSpec, goods.Id)
	if err != nil {
		db.Rollback()
		return 0, err
	}

	//????????????
	err = addProductSpec(db, specInfo, uin, goods.Id, []model.GoodsProductSpecs{})
	if err != nil {
		db.Rollback()
		return 0, err
	}
	db.Commit()
	return goods.Id, nil
}

func GetGoodsMsg(goodsId, uin int) (GoodsMsg, error) {
	goods := model.Goods{
		Id: goodsId,
	}
	err := goods.GetInfo(config.GVA_DB)
	goodsInfo := GoodsMsg{
		ID:                goods.Id,
		Uin:               goods.Uin,
		Aid:               goods.Aid,
		Name:              goods.Name,
		Desc:              goods.Desc,
		Source:            goods.Source,
		SourceContent:     html.UnescapeString(goods.SourceContent),
		Context:           goods.Context,
		OrderNum:          goods.OrderNum,
		Type:              goods.Type,
		Freight:           fmt.Sprintf("%.2f", goods.Freight),
		FreightTemplateID: goods.FreightTemplateId,
		FreightType:       goods.FreightType,
		PickUpAddress:     goods.PickUpAddress,
		PickUpDeadline:    goods.PickUpDeadline,
		ContactNumber:     goods.ContactNumber,
		QrCode:            ossService.Get(goods.QrCode, uin, module),
		IsTop:             goods.IsTop,
		Status:            goods.Status,
		CreateTime:        goods.CreateTime,
		UpdateTime:        goods.UpdateTime,
		CategoryID:        goods.CategoryId,
		LaunchTime:        goods.LaunchTime,
	}

	if goods.LimitNum != 0 {
		goodsInfo.LimitNum = goods.LimitNum
	}

	if goods.ShopPrice != 0 {
		goodsInfo.ShopPrice = ""
	} else {
		goodsInfo.ShopPrice = fmt.Sprintf("%.2f", goods.ShopPrice)
	}

	imgList, err := getImgList(goodsId, uin)
	if err != nil {
		return goodsInfo, err
	}
	goodsInfo.Imgs = imgList
	spec, err := getSpec(goodsId)
	if err != nil {
		return goodsInfo, err
	}
	goodsInfo.Spec = spec

	productList, minPrice, maxPrice, sum, stock, err := getProductInfo(goodsId, uin)
	if err != nil {
		return goodsInfo, err
	}
	goodsInfo.Sum = sum
	goodsInfo.Stock = stock
	if minPrice == maxPrice {
		goodsInfo.Price = fmt.Sprintf("%.2f", minPrice)
	} else {
		goodsInfo.Price = []string{fmt.Sprintf("%.2f", minPrice), fmt.Sprintf("%.2f", maxPrice)}
	}
	goodsInfo.SpecInfo = productList
	return goodsInfo, err
}

func EditGoods(req types.GoodsStruct, uin, aid int) error {
	//???????????????
	goods, specInfo, goodsSpec, err := dealGoods(req, uin, aid)
	if err != nil {
		return err
	}
	//
	oldProductList, err := model.GetProductList(config.GVA_DB, req.GoodsID)
	if err != nil {
		return err
	}
	oldProductMap := map[int]model.GoodsProductSpecs{}
	for _, v := range oldProductList {
		oldProductMap[v.Id] = v
	}

	//????????????????????????
	numFlag := true
	for k, v := range specInfo {
		if _, ok := oldProductMap[v.ID]; ok {
			if v.Sum < oldProductMap[v.ID].Sum {
				numFlag = false
				break
			} else {
				specInfo[k].Stock = v.Sum - oldProductMap[v.ID].Sum + oldProductMap[v.ID].Stock
			}
		} else {
			specInfo[k].Stock = v.Sum
		}
	}

	if numFlag == false && req.GoodsType != 2 {
		return errors.New("???????????????????????????????????????")
	}

	oldImgList, err := model.GetImgList(config.GVA_DB, req.GoodsID)
	if err != nil {
		return err
	}

	oldProductSpec, err := model.GetProductList(config.GVA_DB, req.GoodsID)
	if err != nil {
		return err
	}

	//????????????
	db := config.GVA_DB.Begin()
	goods.Id = req.GoodsID
	err = goods.UpdateGoods(db)
	if err != nil {
		db.Rollback()
		return err
	}

	//??????????????????
	err = model.DeleteImg(db, oldImgList)
	if err != nil {
		db.Rollback()
		return err
	}
	//???????????????
	err = model.CreateImgMulti(db, dealImg(goods.Id, uin, req.GoodsImgs, oldImgList))
	if err != nil {
		db.Rollback()
		return err
	}

	//???????????????
	err = model.DeleteSpec(db, goods.Id)
	if err != nil {
		db.Rollback()
		return err
	}

	err = model.DeleteSpecValue(db, goods.Id)
	if err != nil {
		db.Rollback()
		return err
	}

	//????????????
	err = addSpec(db, goodsSpec, goods.Id)
	if err != nil {
		db.Rollback()
		return err
	}

	//????????????
	err = addProductSpec(db, specInfo, uin, goods.Id, oldProductSpec)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func DeleteGoods(goodsId, uin int) error {
	db := config.GVA_DB.Begin()
	goods := model.Goods{
		Id:  goodsId,
		Uin: uin,
	}

	err := goods.Delete(db)
	if err != nil {
		db.Rollback()
		return err
	}
	oldImgList, err := model.GetImgList(config.GVA_DB, goodsId)
	if err != nil {
		db.Rollback()
		return err
	}
	//??????????????????
	err = model.DeleteImg(db, oldImgList)
	if err != nil {
		db.Rollback()
		return err
	}

	//???????????????
	err = model.DeleteSpec(db, goods.Id)
	if err != nil {
		db.Rollback()
		return err
	}

	//??????product
	err = model.DeleteProductByGoodsId(db, goods.Id)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

func SetGoodsStatus(goodsId, status int)  error {
	goods:= model.Goods{
		Id: goodsId,
		Status: status,
	}
	err := goods.UpdateStatus(config.GVA_DB)
	return err
}

func GetGoodsList(name string, uin, aid, categoryId, goodsType, status, sTime, eTime, page, num int) (interface{}, error) {
	goods := model.Goods{
		Uin: uin,
		Aid: aid,
	}
	tmpList, err := goods.GetList(config.GVA_DB, name, categoryId, goodsType, status, sTime, eTime, page, num)
	return tmpList, err
}

func getImgList(goodsId, uin int) ([]string, error) {
	imgList := []string{}
	tmpList, err := model.GetImgList(config.GVA_DB, goodsId)
	if err != nil {
		return imgList, err
	}
	for _, v := range tmpList {
		img := ossService.Get(v.Img, uin, module)
		imgList = append(imgList, img)
	}
	return imgList, err
}

func getSpec(goodsId int) ([]types.GoodsSpec, error) {
	specList := []types.GoodsSpec{}
	tmpList, err := model.GetSpecList(config.GVA_DB, goodsId)
	if err != nil {
		return specList, err
	}

	valueArr := map[int][]types.SpecValue{}
	specFlag := map[int]bool{}
	for _, v := range tmpList {
		valueArr[v.SpecId] = append(valueArr[v.SpecId], types.SpecValue{
			ID:   v.ValueId,
			Name: v.ValueName,
		})
		if specFlag[v.Id] == false {
			specList = append(specList, types.GoodsSpec{
				ID:      v.Id,
				GoodsID: v.GoodsId,
				Name:    v.Name,
			})
		}
		specFlag[v.Id] = true
	}
	for k, v := range specList {
		if _, ok := valueArr[v.ID]; ok {
			specList[k].SpecValue = valueArr[v.ID]
		}
	}
	return specList, err
}

func getProductInfo(goodsId, uin int) ([]types.ResponseGoodsSpecInfo, float64, float64, int, int, error) {
	productList := []types.ResponseGoodsSpecInfo{}
	price := 0.00
	sum := 0
	stock := 0
	minPrice, maxPrice := price, price
	tmpList, err := model.GetProductList(config.GVA_DB, goodsId)
	if err != nil {
		return productList, minPrice, maxPrice, sum, stock, err
	}
	price = tmpList[0].Price
	minPrice, maxPrice = price, price

	for _, v := range tmpList {
		productList = append(productList, types.ResponseGoodsSpecInfo{
			Specs:   v.Specs,
			ID:      v.Id,
			GoodsID: v.GoodsId,
			Stock:   v.Stock,
			Sum:     v.Sum,
			Price:   fmt.Sprintf("%.2f", v.Price),
			Img:     ossService.Get(v.Img, uin, module),
		})
		sum += v.Sum
		stock += v.Stock
		if v.Price < minPrice {
			minPrice = v.Price
		}
		if v.Price > maxPrice {
			maxPrice = v.Price
		}
	}
	return productList, minPrice, maxPrice, sum, stock, err
}

func addSpec(db *gorm.DB, goodsSpec []types.GoodsSpec, goodId int) error {
	for _, v := range goodsSpec {
		specModel := model.GoodsSpec{
			GoodsId: goodId,
			Name:    v.Name,
		}
		err := specModel.Create(db)
		if err != nil {
			db.Rollback()
			return err
		}
		specValueList := []model.GoodsSpecValue{}
		for _, v1 := range v.SpecValue {
			specValueList = append(specValueList, model.GoodsSpecValue{
				SpecId:  specModel.Id,
				Name:    v1.Name,
				GoodsId: goodId,
			})
		}
		if len(specValueList) > 0 {
			err = model.CreateMultiSpecValue(db, specValueList)
			if err != nil {
				db.Rollback()
				return err
			}
		}
	}
	return nil
}

func dealGoods(req types.GoodsStruct, uin, aid int) (goods model.Goods, specInfo []types.GoodsSpecInfo, goodsSpec []types.GoodsSpec, err error) {
	if req.GoodsType == 2 { //2???????????????
		if req.GoodsSource == 0 || req.GoodsSourceContent == "" {
			err = errors.New("????????????")
		}
	}
	paramFlag := true
	json.Unmarshal([]byte(req.GoodsSpecInfo), &specInfo)

	json.Unmarshal([]byte(req.GoodsSpec), &goodsSpec)

	//??????key????????? ?????????????????????map????????????
	mapStr := []string{}
	for _, v := range goodsSpec {
		mapStr = append(mapStr, v.Name)
	}

	if len(specInfo) == 0 {
		if req.GoodsPrice < 0 || req.GoodsStock < 0 {
			paramFlag = false
		}
		specInfo = append(specInfo, types.GoodsSpecInfo{
			SpecsStr: "default",
			Price:    req.GoodsPrice,
			Stock:    req.GoodsStock,
			Sum:      req.GoodsStock,
		})
	} else {
		for k, v := range specInfo {
			if v.Price < 0 || v.Sum < 0 {
				paramFlag = false
				break
			}
			//specInfo ???map?????????
			specInfo[k].SpecsStr = utils.MapToJson(v.Specs, mapStr...)
			specInfo[k].Stock = v.Sum
			specInfo[k].Img = v.Img
		}
	}
	if paramFlag == false {
		err = errors.New("????????????????????????")
		return
	}

	qrCode, _, err := ossService.Copy(req.QrCode, uin, module, false)
	if err != nil {
		return
	}
	pickUpDeadline := req.PickUpDeadline
	if pickUpDeadline > 7 {
		pickUpDeadline = 7
	}

	goods = model.Goods{
		Type:              req.GoodsType,
		Name:              req.GoodsName,
		Desc:              req.GoodsDesc,
		ShopPrice:         req.GoodsShopPrice,
		Context:           req.GoodsContext,
		Uin:               uin,
		Aid:               aid,
		Freight:           req.Freight,
		FreightType:       req.FreightType,
		FreightTemplateId: req.FreightTemplateID,
		LimitNum:          req.LimitNum,
		ContactNumber:     req.ContactNumber,
		QrCode:            qrCode,
		PickUpDeadline:    pickUpDeadline,
		CategoryId:        req.CategoryID,
		PickUpAddress:     req.PickUpAddress,
		Source:            req.GoodsSource,
		SourceContent:     html.EscapeString(req.GoodsSourceContent),
	}

	return
}

func dealImg(goodsId, uin int, goodsImg string, oldImgList []model.GoodsImgs) []model.GoodsImgs {
	var imgArr []string
	oldImgMap := map[string]int{}
	for _, v := range oldImgList {
		img := ossService.Get(v.Img, uin, module)
		oldImgMap[img] = 0
	}
	json.Unmarshal([]byte(goodsImg), &imgArr)
	imgList := []model.GoodsImgs{}
	//????????????
	for _, v := range imgArr {
		img, _, _ := ossService.Copy(v, uin, module, false)
		if _, ok := oldImgMap[v]; ok { //??????????????????
			oldImgMap[v] = 1
		}
		imgList = append(imgList, model.GoodsImgs{
			GoodsId: goodsId,
			Img:     img,
		})
	}
	//????????????????????????????????????
	for k, v := range oldImgMap {
		if v == 0 {
			ossService.Delete(k, uin, module)
		}
	}
	return imgList
}

func addProductSpec(db *gorm.DB, specInfo []types.GoodsSpecInfo, uin, goodsId int, oldProductSpec []model.GoodsProductSpecs) error {
	oldSpecInfoList := map[int]model.GoodsProductSpecs{}
	for _, v := range oldProductSpec {
		oldSpecInfoList[v.Id] = v
	}

	var err error
	errFlag := true
	for _, v := range specInfo {
		img, _, _ := ossService.Copy(v.Img, uin, module, true)
		product := model.GoodsProductSpecs{
			GoodsId: goodsId,
			Specs:   v.SpecsStr,
			Stock:   v.Stock,
			Sum:     v.Sum,
			Price:   v.Price,
			Img:     img,
		}
		if _, ok := oldSpecInfoList[v.ID]; ok {
			product.Id = v.ID
			err = product.Update(db)
			delete(oldSpecInfoList, v.ID)
		} else {
			err = product.Create(db)
		}
		if err != nil {
			errFlag = false
			break
		}
	}
	if errFlag == false {
		return err
	}

	deleteIdArr := []int{}
	for _, v := range oldSpecInfoList {
		deleteIdArr = append(deleteIdArr, v.Id)
	}

	//????????????
	if len(deleteIdArr) > 0 {
		err = model.DeleteProductSpec(db, deleteIdArr)
	}
	return err
}
