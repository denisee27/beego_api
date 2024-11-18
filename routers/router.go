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
	beego.Router("/items/create", &controllers.UserController{}, "post:CreateItem")
	beego.Router("/items/:id", &controllers.UserController{}, "get:GetItemByID")
	beego.Router("/items", &controllers.UserController{}, "get:GetItems")
	beego.Router("/items/update/:id", &controllers.UserController{}, "put:UpdateItem")
	beego.Router("/items/delete/:id", &controllers.UserController{}, "delete:DeleteItem")

}
