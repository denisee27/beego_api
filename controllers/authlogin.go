package controllers

import (
	"beego_api/models"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (request *AuthController) Login() {
	var loginStruct models.LoginRequest
	if err := json.Unmarshal(request.Ctx.Input.RequestBody, &loginStruct); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Invalid format data",
		}
		request.ServeJSON()
		return
	}
	//Generate JWT Secret Key
	// err := models.GenerateJwtKey()
	// if err != nil {
	// 	request.Data["json"] = map[string]interface{}{
	// 		"status": http.StatusInternalServerError,
	// 		"wrong":  "Something Wrong!",
	// 	}
	// }

	result, err := models.AuthLogin(&loginStruct)
	if err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  err.Error(),
		}
	} else {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": result,
		}
	}
	request.ServeJSON()
}
