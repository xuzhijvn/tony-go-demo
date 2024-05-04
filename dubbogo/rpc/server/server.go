package main

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports" // dubbogo 框架依赖，所有dubbogo进程都需要隐式引入一次
	"dubbogo3-rpc/api"
	"github.com/dubbogo/gost/log/logger"
	"strconv"
	"time"
)

type UserProvider struct {
}

// 实现接口方法
func (u *UserProvider) GetUser(ctx context.Context, req int32) (*api.User, error) {
	var err error
	logger.Infof("req:%#v", req)
	user := &api.User{}
	user.ID = strconv.Itoa(int(req))
	user.Name = "tonyxu"
	user.Age = 22
	user.Time = time.Now()
	return user, err
}

//// MethodMapper 定义方法名映射，从 Go 的方法名映射到 Java 小写方法名，只有 dubbo 协议服务接口才需要使用
//// go -> go 互通无需使用
//func (s *UserProvider) MethodMapper() map[string]string {
//	return map[string]string{
//		"GetUser": "getUser",
//	}
//}

func init() {
	config.SetProviderService(&UserProvider{}) // 注册服务提供者类，类名与配置文件中的 service 对应
}

// export DUBBO_GO_CONFIG_PATH=dubbogo.yml 运行前需要设置环境变量，指定配置文件位置
func main() {
	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}
