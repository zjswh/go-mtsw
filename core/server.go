package core

import (
	"fmt"
	"mtsw/global"
	"mtsw/initialize"
	"net/http"
	"time"
)

func RunServer() {
	//加载路由
	router := initialize.Routers()


	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)

	//http
	s := &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
