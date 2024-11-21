// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego_api/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/items/create/", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/items/:id", &controllers.UserController{}, "get:GetUserId")
	beego.Router("/items", &controllers.UserController{}, "get:GetUsers")
	beego.Router("/items/update/:id", &controllers.UserController{}, "put:UpdateUser")
	beego.Router("/items/delete/:id", &controllers.UserController{}, "delete:DeleteUser")

}
