package routes

import (
	"github.com/iralance/go-lottery/bootstrap"
	"github.com/iralance/go-lottery/services"
	"github.com/iralance/go-lottery/web/controllers"
	"github.com/iralance/go-lottery/web/middleware"
	"github.com/kataras/iris/v12/mvc"
)

func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	index := mvc.New(b.Party("/"))
	index.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	index.Handle(new(controllers.IndexController))

	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	admin.Handle(new(controllers.AdminController))

	adminUser := admin.Party("/user")
	adminUser.Register(userService)
	adminUser.Handle(new(controllers.AdminUserController))

	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(controllers.AdminGiftController))

	adminCode := admin.Party("/code")
	adminCode.Register(codeService)
	adminCode.Handle(new(controllers.AdminCodeController))

	adminResult := admin.Party("/result")
	adminResult.Register(resultService)
	adminResult.Handle(new(controllers.AdminResultController))

	adminBlackip := admin.Party("/blackip")
	adminBlackip.Register(blackipService)
	adminBlackip.Handle(new(controllers.AdminBlackipController))

}
