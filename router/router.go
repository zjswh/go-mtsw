package router

import (
	"github.com/gin-gonic/gin"
	v1 "mtsw/api/v1"
	"mtsw/service"
)

func InitRouter(Router *gin.RouterGroup) {
	goodsRouter := Router.Group("Goods").
		Use(service.CheckLogin())
	{
		goodsRouter.POST("setMailInfo", v1.SetMailInfo)           //设置商城信息
		goodsRouter.GET("getMailInfo", v1.GetMailInfo)            //获取商城信息
		goodsRouter.POST("createCategories", v1.CreateCategories) //创建分类
		goodsRouter.POST("changeCategories", v1.ChangeCategories) //修改分类信息
		goodsRouter.POST("saveCategories", v1.SaveCategories)     //保存分类顺序
		goodsRouter.GET("getCategories", v1.GetCategories)        //获取分类列表
		goodsRouter.POST("addGoods", v1.AddGoods)                 //添加商品
		goodsRouter.POST("editGoods", v1.EditGoods)               //编辑商品
		goodsRouter.GET("getGoodsList", v1.GetGoodsList)          //获取商品列表
		goodsRouter.POST("delGoods", v1.DealGoods)                //删除商品
		goodsRouter.GET("getGoodsMsg", v1.GetGoodsMsg)            //获取商品详情
		goodsRouter.POST("onShelfGoods", v1.OnShelfGoods)         //上架商品 新
		goodsRouter.POST("offShelfGoods", v1.OffShelfGoods)       //下架商品 新
		//goodsRouter.POST("setTopGoods", v1.AuditReport) //设置推荐商品
		//goodsRouter.POST("cancelTopGoods", v1.AuditReport) //取消置顶商品
		//goodsRouter.POST("bindGoodsToMenu", v1.AuditReport) //添加商品与菜单的绑定关系
		//goodsRouter.POST("unBindGoodsToMenu", v1.AuditReport) //取消置顶商品
		//goodsRouter.GET("getSoldGoods", v1.AuditReport) //获取上架商品列表 旧
		//goodsRouter.POST("setSoldGoods", v1.AuditReport) //上架商品修改权重 旧
		//goodsRouter.POST("setTrackingNo", v1.AuditReport) //设置运单号
		//goodsRouter.GET("getTrackingCompany", v1.AuditReport) //获取支持显示物流公司
		//goodsRouter.POST("importOrder", v1.AuditReport) //导入订单 批量修改运单号
	}

	showCaseRouter := Router.Group("ShowCase").
		Use(service.CheckLogin())
	{
		showCaseRouter.GET("info", v1.CaseInfo)//获取橱窗信息
		showCaseRouter.GET("list", v1.CaseList) //获取橱窗列表
		showCaseRouter.POST("create", v1.CreateCase) //创建橱窗
		showCaseRouter.POST("edit", v1.EditCase) //修改橱窗
		showCaseRouter.POST("delete", v1.DeleteCase) //删除橱窗
		showCaseRouter.GET("getCaseGoodsList", v1.GetCaseGoodsList) //获取橱窗商品列表
		showCaseRouter.GET("getGoodsList", v1.GetGoodsListInCase)
		showCaseRouter.POST("addGoods", v1.AddGoodsInCase) //橱窗添加商品
		//showCaseRouter.POST("setGoodsSort", v1.Publish) //保存商品排序
		//showCaseRouter.POST("removeGoods", v1.AuditReport) //橱窗移除商品
		showCaseRouter.POST("setTopGoods", v1.SetTopGoods) //置顶商品
		showCaseRouter.POST("unSetTopGoods", v1.UnSetTopGoods) //取消置顶商品
		showCaseRouter.POST("setExplain", v1.SetExplain) //设置讲解状态
		showCaseRouter.POST("unSetExplain", v1.UnSetExplain) //取消讲解状态
		//showCaseRouter.GET("getIncludeList", v1.AuditReport) //获取橱窗
		//showCaseRouter.POST("saveConfig", v1.AuditReport) //配置橱窗
		//showCaseRouter.POST("getConfig", v1.AuditReport) //获取配置信息
	}
	//
	//shippingRouter := Router.Group("Shipping").
	//	Use(service.CheckLogin())
	//{
	//	shippingRouter.GET("info", v1.SetGuestSwitch)//获取运费模板信息
	//	shippingRouter.GET("list", v1.SetGuestSwitch)//获取运费模板列表
	//	shippingRouter.POST("add", v1.GetLiveBaseSetting) //创建运费模板
	//	shippingRouter.POST("update", v1.GetLiveBaseSetting) //更新运费模板
	//	shippingRouter.POST("delete", v1.GetLiveBaseSetting) //删除运费模板
	//}
	//
	//orderRouter := Router.Group("Order").
	//	Use(service.CheckLogin())
	//{
	//	orderRouter.GET("getMtswOrderList", v1.SetGuestSwitch)//获取订单列表
	//	orderRouter.GET("getOrderStream", v1.SetGuestSwitch)//商品购买流水
	//	orderRouter.GET("getRefundReason", v1.SetGuestSwitch)//获取退款理由
	//	orderRouter.GET("getPayRefundList", v1.SetGuestSwitch)//退款申请列表
	//	orderRouter.POST("dealRefundOrder", v1.SetGuestSwitch)//处理退款申请
	//}
	//
	//crontabRouter := Router.Group("Crontab").
	//	Use(service.CheckLogin())
	//{
	//	crontabRouter.GET("dealOverDaysOrderToComplete", v1.SetGuestSwitch)//超过七天的订单的状态为已完成
	//}
}
