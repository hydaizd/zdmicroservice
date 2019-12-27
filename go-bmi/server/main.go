package main

import (
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/lager"
	"github.com/hydaizd/zdmicroservice/go-bmi/server/app"
)

/**
服务注册与配置
完成业务逻辑代码的编写之后，需要将业务逻辑注册到Go-chassis框架，注册时可以同时指定微服务的名称、ID 等属性
除了在代码中指定的部分属性外，更多的属性是通过配置文件来进行配置。配置文件包括chassis.yaml和microservice.yaml，放置于代码目录下的conf文件夹内。
*/
func main() {
	chassis.RegisterSchema("rest", &app.CalculateBmi{})

	// 初始化框架
	if err := chassis.Init(); err != nil {
		lager.Logger.Errorf("Init FAILED %s", err.Error())
		return
	}
	// 运行微服务
	chassis.Run()
}
