package api

import (
	"context"
	"dubbo.apache.org/dubbo-go/v3/config"
	hessian "github.com/apache/dubbo-go-hessian2"
	"time"
)

// 1. 定义传输结构， 如需 Java 互通，字段需要与 Java 侧对应，首字母大写
type User struct {
	ID   string
	Name string
	Age  int32
	Time time.Time
}

func (u *User) JavaClassName() string {
	return "org.apache.dubbo.User" // 如果与 Java 互通，需要与 Java 侧 User class全名对应,
}

var (
	UserProviderClient = &UserProvider{} // 客户端指针
)

// 2。 定义客户端存根类：UserProvider
type UserProvider struct {
	// dubbo标签，用于适配go侧客户端大写方法名 -> java侧小写方法名，只有 dubbo 协议客户端才需要使用
	GetUser func(ctx context.Context, req int32) (*User, error) //`dubbo:"getUser"`
}

func init() {
	hessian.RegisterPOJO(&User{}) // 注册传输结构到 hessian 库
	// 注册客户端存根类到框架，实例化客户端接口指针 userProvider
	config.SetConsumerService(UserProviderClient)
}
