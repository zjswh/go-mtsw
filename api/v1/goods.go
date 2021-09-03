package v1

import (
	"github.com/gin-gonic/gin"
	"mtsw/service"
	"mtsw/service/goodsService"
	"mtsw/types"
	"mtsw/types/response"
	"mtsw/utils"
)

func SetMailInfo(c *gin.Context) {
	var req types.SetMailInfoStruct
	err := c.ShouldBind(&req)
	if err != nil {
		response.ParamError(err.Error(), c)
		return
	}
	userInfo := service.GetUserInfo(c)
	err = goodsService.SetMailInfo(req, userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func GetMailInfo(c *gin.Context) {
	userInfo := service.GetUserInfo(c)
	info, err := goodsService.GetMailInfo(userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(info, c)
	return
}

func CreateCategories(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name == "" {
		response.ParamError("参数缺失", c)
		return
	}

	userInfo := service.GetUserInfo(c)
	id, err := goodsService.CreateCategory(name, userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(id, c)
	return
}

func ChangeCategories(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	categoryId := utils.DefaultIntFormValue("categoryId", 0, c)
	if name == "" || categoryId == 0 {
		response.ParamError("参数缺失", c)
		return
	}

	userInfo := service.GetUserInfo(c)
	err := goodsService.UpdateCategory(name, categoryId, userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("修改成功", c)
	return
}

func SaveCategories(c *gin.Context) {
	categoryIds := c.DefaultPostForm("categoryIds", "")
	if categoryIds == "" {
		response.ParamError("参数缺失", c)
		return
	}

	err := goodsService.SaveCategories(categoryIds)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("保存成功", c)
	return
}

func GetCategories(c *gin.Context) {
	userInfo := service.GetUserInfo(c)
	list, err := goodsService.GetCategories(userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(list, c)
	return
}

func AddGoods(c *gin.Context) {
	var req types.GoodsStruct
	err := c.ShouldBind(&req)
	if err != nil {
		response.ParamError(err.Error(), c)
		return
	}
	userInfo := service.GetUserInfo(c)
	id, err := goodsService.AddGoods(req, userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(id, c)
	return
}

func EditGoods(c *gin.Context) {
	var req types.GoodsStruct
	err := c.ShouldBind(&req)
	if err != nil {
		response.ParamError(err.Error(), c)
		return
	}
	if req.GoodsID == 0 {
		response.ParamError(err.Error(), c)
		return
	}

	userInfo := service.GetUserInfo(c)
	err = goodsService.EditGoods(req, userInfo.Uin, userInfo.Aid)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("success", c)
	return
}

func DealGoods(c *gin.Context) {
	goodsId := utils.DefaultIntFormValue("goodsId", 0, c)
	if goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	userInfo := service.GetUserInfo(c)
	err := goodsService.DeleteGoods(goodsId, userInfo.Uin)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("删除成功", c)
	return
}

func GetGoodsMsg(c *gin.Context) {
	goodsId := utils.DefaultIntParam("goodsId", 0, c)
	if goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	userInfo := service.GetUserInfo(c)
	info, err := goodsService.GetGoodsMsg(goodsId, userInfo.Uin)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(info, c)
	return
}

func GetGoodsList(c *gin.Context) {
	categoryId := utils.DefaultIntParam("categoryId", 0, c)
	goodsType := utils.DefaultIntParam("goodsType", 0, c)
	status := utils.DefaultIntParam("status", -1, c)
	goodsName := c.DefaultQuery("goodsName", "")
	page := utils.DefaultIntParam("page", 1, c)
	num := utils.DefaultIntParam("num", 10, c)
	sTime := utils.DefaultIntParam("sTime", 0, c)
	eTime := utils.DefaultIntParam("eTime", 0, c)

	userInfo := service.GetUserInfo(c)

	list, err := goodsService.GetGoodsList(goodsName, userInfo.Uin, userInfo.Aid, categoryId, goodsType, status, sTime, eTime, page, num)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success(list, c)
	return
}

func OnShelfGoods(c *gin.Context) {
	goodsId := utils.DefaultIntFormValue("id", 0, c)
	if goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := goodsService.SetGoodsStatus(goodsId, 1)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}

func OffShelfGoods(c *gin.Context) {
	goodsId := utils.DefaultIntFormValue("id", 0, c)
	if goodsId == 0 {
		response.ParamError("参数缺失", c)
		return
	}
	err := goodsService.SetGoodsStatus(goodsId, 0)
	if err != nil {
		response.DbError(err.Error(), c)
		return
	}
	response.Success("设置成功", c)
	return
}
