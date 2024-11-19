package controllers

import (
	"beego_api/models"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
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
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		request.Data["json"] = map[string]interface{}{
			"status":  http.StatusBadRequest,
			"message": "Validation failed",
			"wrong":   err.Error(),
		}
		request.ServeJSON()
		return
	}
	insertedUser, err := models.CreateUser(u)
	if insertedUser != nil {
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

func (uc *UserController) UpdateUser() {
	var u models.Users
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	user := models.UpdateUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}

func (request *UserController) DeleteUser() {
	// Ambil ID dari parameter
	id := request.GetString(":id")

	// Panggil fungsi model untuk menghapus user
	err := models.DeleteUser(id)
	if err != nil {
		// Jika terjadi error, kirimkan respons error
		request.Data["json"] = map[string]interface{}{
			"status":  http.StatusInternalServerError,
			"result":  "error",
			"message": err.Error(),
		}
	} else {
		// Jika berhasil, kirimkan respons sukses
		request.Data["json"] = map[string]interface{}{
			"status": http.StatusOK,
			"result": "ok",
		}
	}
	request.ServeJSON()
}
