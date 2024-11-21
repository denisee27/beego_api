package controllers

import (
	"beego_api/models"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (request *UserController) GetUserId() {
	id := request.GetString(":id")
	user := models.GetUserById(string(id))
	if user == nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": map[string]interface{}{},
		}
	} else {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": user,
		}
		request.ServeJSON()
	}
}

func (request *UserController) GetUsers() {
	request.Data["json"] = map[string]interface{}{
		"status": http.StatusOK,
		"result": models.GetAllUsers(),
	}
	request.ServeJSON()
}

func (request *UserController) CreateUser() {
	var requestBody map[string]json.RawMessage
	if err := json.Unmarshal(request.Ctx.Input.RequestBody, &requestBody); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Invalid JSON format",
		}
		request.ServeJSON()
		return
	}
	var u models.Users
	if err := json.Unmarshal(requestBody["data"], &u); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Invalid 'data' format",
		}
		request.ServeJSON()
		return
	}
	if models.IsEmailExists(u.Email) {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Email already exists",
		}
		request.ServeJSON()
		return
	}
	err := models.CreateUser(u)
	if err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusInternalServerError,
			"wrong":  err.Error(),
		}
	} else {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": "ok",
		}
	}
	request.ServeJSON()
}

func (request *UserController) UpdateUser() {
	var requestBody map[string]json.RawMessage
	if err := json.Unmarshal(request.Ctx.Input.RequestBody, &requestBody); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Invalid JSON format",
		}
		request.ServeJSON()
		return
	}
	var u models.Users
	if err := json.Unmarshal(requestBody["data"], &u); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusBadRequest,
			"wrong":  "Invalid 'data' format",
		}
		request.ServeJSON()
		return
	}
	if err := models.UpdateUser(u); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusInternalServerError,
			"wrong":  err.Error(),
		}
		request.ServeJSON()
		return
	}
	request.Data["json"] = map[string]interface{}{
		"status": http.StatusOK,
		"result": "ok",
	}
	request.ServeJSON()
}

func (request *UserController) DeleteUser() {
	id := request.GetString(":id")
	err := models.DeleteUser(id)
	if err != nil {
		request.Data["json"] = map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"result":  "error",
			"message": err.Error(),
		}
	} else {
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": "ok",
		}
	}
	request.ServeJSON()
}
