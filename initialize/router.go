package initialize

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mtsw/config"
	"mtsw/middleware"
	"mtsw/router"
	response2 "mtsw/types/response"
	"net/http"
	"time"
)

func RunServer() {
	//加载路由
	_router := Routers()

	address := fmt.Sprintf(":%d", config.GVA_CONFIG.System.Addr)

	//http
	s := &http.Server{
		Addr:           address,
		Handler:        _router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}


func Routers() *gin.Engine {
	Router := gin.Default()

	//处理跨域
	//Router.Use(cors())

	//日记记录
	Router.Use(middleware.RequestId(), middleware.LoggerToFile()) //, middleware.Exception()

	ApiGroup := Router.Group("v1/mtsw")
	router.InitRouter(ApiGroup)

	//处理404
	Router.NoMethod(HandleNotFind)
	Router.NoRoute(HandleNotFind)

	return Router
}

//处理404
func HandleNotFind(c *gin.Context)  {
	response2.Result(4, "api不存在", "", c)
}

//跨域
func cors() gin.HandlerFunc  {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-ca-stage")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//过滤options请求
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}
