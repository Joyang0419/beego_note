package routers

import (
	"github.com/Joyang0419/beego_note/controllers"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

func InitAuth(ormer orm.Ormer) {
	authController := &controllers.AuthController{
		Ormer: ormer,
	}

	web.Router("/register", authController, "post:Register")
}
