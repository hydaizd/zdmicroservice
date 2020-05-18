package main

import (
	"fmt"
	"github.com/go-chassis/go-chassis"
	"github.com/go-chassis/go-chassis/core/server"
	"github.com/hydaizd/zdmicroservice/user-service/pb"
	"github.com/hydaizd/zdmicroservice/user-service/server/service"
)

// 如果conf文件夹没有在工作目录下，则运行时需要配置环境变量CHASSIS_CONF_DIR=/path/to/conf_folder，否则运行时会提示找不到配置文件
func main() {
	// chassis.RegisterSchema("rest", &service.Server{})
	chassis.RegisterSchema("grpc", &service.Server{}, server.WithRPCServiceDesc(&pb.UserService_serviceDesc))
	if err := chassis.Init(); err != nil {
		fmt.Println("Init failed.")
		return
	}
	_ = chassis.Run()
}
