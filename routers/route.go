package routers

import (
	"github.com/Joyang0419/beego_note/controllers"
	"github.com/beego/beego/v2/server/web"
)

func InitRoute() {
	authController := &controllers.AuthController{}

	web.Router("/register", authController, "post:Register")
	web.Router("/signin", authController, "post:SignIn")

	sessionController := &controllers.SessionController{}
	web.Router("/sessioninfo", sessionController, "get:SessionInfo")
}
