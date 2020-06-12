package main

import (
	"adminframe/application/model"
	"adminframe/framework"
	"adminframe/framework/config"
	"fmt"
	"net/http"
)

func init(){
	// #### 初始化配置文件
	config.InitConfig()
	// #### 初始化数据库连接
	model.InitModel()
}

func main(){
	app := framework.NewApp()
	port := fmt.Sprintf(":%d", config.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           port,
		Handler:        app,
		ReadTimeout:    config.ServerSetting.ReadTimeout,
		WriteTimeout:   config.ServerSetting.WriteTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("[Err] Start Server Error : %s",err))
	}
}
