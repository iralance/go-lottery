package main

import (
	"fmt"
	"github.com/iralance/go-lottery/bootstrap"
	"github.com/iralance/go-lottery/web/middleware/identity"
	"github.com/iralance/go-lottery/web/routes"
)

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	// 初始化应用
	app := bootstrap.New("抽奖系统", "iralance")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}

func main() {
	app := newApp()
	app.Listen(fmt.Sprintf(":%d", port))
}
