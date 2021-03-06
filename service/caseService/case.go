package caseService

import (
	"errors"
	"github.com/zjswh/go-tool/ossService"
	"mtsw/config"
	"mtsw/model"
)

const CaseNumLimit = 20
const module = "program"

type CaseListStruct struct {
	CoverImg   string `json:"coverImg"`
	Title      string `json:"title"`
	Id         int    `json:"id"`
	CreateTime int64  `json:"createTime"`
}

type CaseGoodsListStruct struct {
	model.CaseGoods
	CoverImg string `json:"coverImg"`
}

type GoodsList struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	CoverImg string `json:"coverImg"`
	Selected int    `json:"selected"`
}

func GetCaseList(uin, aid int) (interface{}, error) {
	list := []CaseListStruct{}
	showCase := model.GoodsShowCase{
		Uin: uin,
		Aid: aid,
	}
	tmpList, err := showCase.GetList(config.GVA_DB)
	if err != nil {
		return tmpList, err
	}

	if len(tmpList) > 0 {
		caseIdArr := []int{}
		goodsIdArr := []int{}
		imgMap := map[int]string{}
		caseMap := map[int]int{}
		for _, v := range tmpList {
			caseIdArr = append(caseIdArr, v.Id)
		}
		goodsList, _ := model.GetCaseGoodsByCaseIdArr(config.GVA_DB, caseIdArr)
		if len(goodsList) > 0 {
			for _, v := range goodsList {
				goodsIdArr = append(goodsIdArr, v.GoodsId)
				caseMap[v.CaseId] = v.GoodsId
			}
			imgList, _ := model.GetImgListByGoodsIdArr(config.GVA_DB, goodsIdArr)
			for _, v := range imgList {
				imgMap[v.GoodsId] = ossService.Get(v.Img, uin, module)
			}
		}

		for _, v := range tmpList {
			img := ""
			if _, ok := imgMap[caseMap[v.Id]]; ok {
				img = imgMap[caseMap[v.Id]]
			}
			list = append(list, CaseListStruct{
				CoverImg:   img,
				Title:      v.Title,
				Id:         v.Id,
				CreateTime: v.CreateTime,
			})
		}
	}

	return list, err
}

func GetCaseCount(uin, aid int) (int64, error) {
	showCase := model.GoodsShowCase{
		Uin: uin,
		Aid: aid,
	}
	count, err := showCase.GetCount(config.GVA_DB)
	if err != nil {
		return 0, err
	}

	return count, err
}

func GetCaseInfo(id int) (interface{}, error) {
	showCase := model.GoodsShowCase{
		Id: id,
	}

	err := showCase.GetInfo(config.GVA_DB)
	if err != nil {
		return showCase, err
	}
	info := CaseListStruct{
		Title:      showCase.Title,
		Id:         showCase.Id,
		CreateTime: showCase.CreateTime,
	}
	return info, err
}

func CreateCase(uin, aid int, title string) (int, error) {
	showCase := model.GoodsShowCase{
		Uin:   uin,
		Aid:   aid,
		Title: title,
	}

	count, err := showCase.GetCount(config.GVA_DB)
	if err != nil {
		return 0, err
	}
	if count >= CaseNumLimit {
		return 0, errors.New("??????????????????")

	}

	err = showCase.Create(config.GVA_DB)
	if err != nil {
		return 0, err
	}
	return showCase.Id, err
}

func UpdateCase(id int, title string) error {
	showCase := model.GoodsShowCase{
		Id:    id,
		Title: title,
	}

	err := showCase.Update(config.GVA_DB, "title", "updateTime")
	if err != nil {
		return err
	}
	return nil
}

func DeleteCase(id int) error {
	showCase := model.GoodsShowCase{
		Id:     id,
		Status: 0,
	}

	err := showCase.Update(config.GVA_DB, "status", "updateTime")
	if err != nil {
		return err
	}

	//todo ??????????????????????????? ???????????????????????????
	return nil
}

func GetCaseGoodsList(id, uin, page, num int) (interface{}, error) {
	list := []CaseGoodsListStruct{}
	caseBindModel := model.ShowcaseBindGoods{
		CaseId: id,
	}
	tmpList, err := caseBindModel.GetGoodsList(config.GVA_DB, page, num)
	if len(tmpList) > 0 {
		imgMap := map[int]string{}
		goodsIdArr := []int{}
		for _, v := range tmpList {
			goodsIdArr = append(goodsIdArr, v.GoodsId)
		}
		imgList, _ := model.GetImgListByGoodsIdArr(config.GVA_DB, goodsIdArr)
		for _, v := range imgList {
			imgMap[v.GoodsId] = ossService.Get(v.Img, uin, module)
		}

		for _, v := range tmpList {
			if _, ok := imgMap[v.GoodsId]; ok {
				list = append(list, CaseGoodsListStruct{
					v,
					imgMap[v.GoodsId],
				})
			}
		}

	}
	return list, err
}

func GetCaseGoodsCount(id int) (int64, error) {
	caseBindModel := model.ShowcaseBindGoods{
		CaseId: id,
	}
	count, err := caseBindModel.GetGoodsCount(config.GVA_DB)

	return count, err
}

func GetGoodsList(id, uin, aid, categoryId, page, num int) (interface{}, error) {
	list := []GoodsList{}
	goodsModel := model.Goods{
		Uin:        uin,
		Aid:        aid,
		CategoryId: categoryId,
	}
	tmpList, err := goodsModel.GetListWithCase(config.GVA_DB, id, page, num)
	if len(tmpList) > 0 {
		imgMap := map[int]string{}
		goodsIdArr := []int{}
		for _, v := range tmpList {
			goodsIdArr = append(goodsIdArr, v.Id)
		}

		imgList, _ := model.GetImgListByGoodsIdArr(config.GVA_DB, goodsIdArr)
		for _, v := range imgList {
			imgMap[v.GoodsId] = ossService.Get(v.Img, uin, module)
		}

		for _, v := range tmpList {
			img := ""
			if _, ok := imgMap[v.Id]; ok {
				img = imgMap[v.Id]
			}
			selected := 0
			if v.CaseId > 0 {
				selected = 1
			}
			list = append(list, GoodsList{
				CoverImg: img,
				Id:       v.Id,
				Name:     v.Name,
				Selected: selected,
			})
		}
	}
	return list, err
}

func GetGoodsCount(uin, aid, categoryId int) (int64, error) {
	goodsModel := model.Goods{
		Uin:        uin,
		Aid:        aid,
		CategoryId: categoryId,
		Status:     -1,
	}
	count, err := goodsModel.GetCountWithCase(config.GVA_DB)
	return count, err
}

func UpdateCaseGoodsTopStatus(id, goodsId, isTop int) error {
	caseGoodsModel := model.ShowcaseBindGoods{
		CaseId:  id,
		GoodsId: goodsId,
	}
	//?????????????????????
	if isTop == 1 {
		caseGoodsModel.ClearTopStatus(config.GVA_DB)
	}

	caseGoodsModel.IsTop = isTop
	err := caseGoodsModel.Update(config.GVA_DB, "updateTime", "isTop")

	//??????DMS??????

	//??????key???

	return err
}

func UpdateCaseGoodsExplainingStatus(id, goodsId, isExplaining int) error {
	caseGoodsModel := model.ShowcaseBindGoods{
		CaseId:       id,
		GoodsId:      goodsId,
		IsExplaining: isExplaining,
	}
	//?????????????????????
	if isExplaining == 1 {
		caseGoodsModel.ClearExplainingStatus(config.GVA_DB)
	}

	caseGoodsModel.IsExplaining = isExplaining
	err := caseGoodsModel.Update(config.GVA_DB, "updateTime", "isExplaining")

	//??????DMS??????

	//??????key???
	return err
}

func AddGoods(id int, addGoodsId, deleteGoodsId string) (interface{}, error) {
	//????????????
	caseBindModel := model.ShowcaseBindGoods{
		CaseId: id,
	}
	tmpList, err := caseBindModel.GetGoodsList(config.GVA_DB, 0, 0)
	return tmpList, err
}
