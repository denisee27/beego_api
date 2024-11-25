package routers

import (
	"beego_api/controllers"
	"beego_api/middleware"

	"github.com/astaxie/beego"
)

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, middleware.ValidateToken)

	beego.Router("/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/items/create/", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/items/:id", &controllers.UserController{}, "get:GetUserId")
	beego.Router("/items", &controllers.UserController{}, "get:GetUsers")
	beego.Router("/items/update/:id", &controllers.UserController{}, "put:UpdateUser")
	beego.Router("/items/delete/:id", &controllers.UserController{}, "delete:DeleteUser")

}
