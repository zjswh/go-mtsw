package v1

import (
	"github.com/gin-gonic/gin"
	"mtsw/service"
	"mtsw/service/caseService"
	"mtsw/types/response"
	"mtsw/utils"
	"unicode/utf8"
)

func CaseInfo(c *gin.Context) {
	id := utils.DefaultIntParam("id", 0, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	info, err := caseService.GetCaseInfo(id)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(info, c)
	return
}

func CaseList(c *gin.Context) {
	userInfo := service.GetUserInfo(c)
	list, err := caseService.GetCaseList(userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	count, err := caseService.GetCaseCount(userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(gin.H{
		"list":  list,
		"count": count,
	}, c)
	return
}

func CreateCase(c *gin.Context) {
	title := c.DefaultPostForm("title", "")
	if title == "" {
		response.ParamError("参数缺失", c)
		return
	}
	if utf8.RuneCountInString(title) > 10 {
		response.ParamError("标题过长", c)
		return
	}
	userInfo := service.GetUserInfo(c)
	id, err := caseService.CreateCase(userInfo.Uin, userInfo.Aid, title)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(id, c)
	return
}

func EditCase(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	title := c.DefaultPostForm("title", "")
	if id == 0 || title == "" {
		response.ParamError("参数缺失", c)
		return
	}
	if utf8.RuneCountInString(title) > 10 {
		response.ParamError("标题过长", c)
		return
	}
	err := caseService.UpdateCase(id, title)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("修改成功", c)
	return
}

func DeleteCase(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := caseService.DeleteCase(id)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("删除成功", c)
	return
}

func GetCaseGoodsList(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	page := utils.DefaultIntFormValue("page", 1, c)
	num := utils.DefaultIntFormValue("num", 10, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}

	userInfo := service.GetUserInfo(c)
	list, err := caseService.GetCaseGoodsList(id, userInfo.Uin, page, num)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	count, err := caseService.GetCaseGoodsCount(id)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(gin.H{
		"list":  list,
		"count": count,
	}, c)
	return
}

func GetGoodsListInCase(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	categoryId := utils.DefaultIntFormValue("categoryId", 0, c)
	page := utils.DefaultIntFormValue("page", 1, c)
	num := utils.DefaultIntFormValue("num", 10, c)
	if id == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	userInfo := service.GetUserInfo(c)
	list, err := caseService.GetGoodsList(id, userInfo.Uin, userInfo.Aid, categoryId, page, num)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	count, err := caseService.GetGoodsCount(userInfo.Uin, userInfo.Aid, categoryId)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(gin.H{
		"list":  list,
		"count": count,
	}, c)
	return
}

func SetTopGoods(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	goodsId := utils.DefaultIntFormValue("goodsId", 0, c)
	if id == 0 || goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := caseService.UpdateCaseGoodsTopStatus(id, goodsId, 1)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func UnSetTopGoods(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	goodsId := utils.DefaultIntFormValue("goodsId", 0, c)
	if id == 0 || goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := caseService.UpdateCaseGoodsTopStatus(id, goodsId, 0)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func SetExplain(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	goodsId := utils.DefaultIntFormValue("goodsId", 0, c)
	if id == 0 || goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := caseService.UpdateCaseGoodsExplainingStatus(id, goodsId, 1)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func UnSetExplain(c *gin.Context) {
	id := utils.DefaultIntFormValue("id", 0, c)
	goodsId := utils.DefaultIntFormValue("goodsId", 0, c)
	if id == 0 || goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := caseService.UpdateCaseGoodsExplainingStatus(id, goodsId, 0)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func AddGoodsInCase(c *gin.Context)  {
	id := utils.DefaultIntFormValue("id", 0, c)
	addGoodsId := c.DefaultPostForm("addGoodsId", "")
	deleteGoodsId := c.DefaultPostForm("deleteGoodsId", "")
	da, err := caseService.AddGoods(id, addGoodsId, deleteGoodsId)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(da, c)
	return
}
