package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (request *AuthController) Login() {
	var login LoginRequest

	if err := json.Unmarshal(request.Ctx.Input.RequestBody, &login); err != nil {
		request.Ctx.Output.SetStatus(400)
		request.Data["json"] = map[string]string{"error": "Invalid input"}
		request.ServeJSON()
		return

		// request.Data["json"] = map[string]interface{}{
		// 	"status": http.StatusBadRequest,
		// 	"wrong":  "invalid input",
		// }
		// request.ServeJSON()
		// return
	}

}
