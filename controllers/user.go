package controllers

import (
	"beego_api/models"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/beego/beego/v2/adapter/orm"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) CreateItem() {
	var item models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &item)
	o := orm.NewOrm()
	_, err := o.Insert(&item)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"message": "Item created successfully", "item": item}
	} else {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}
	c.ServeJSON()
}

func (uc *UserController) GetUserById() {
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))
	user := models.GetUserById(id)
	uc.Data["json"] = user
	uc.ServeJSON()
}

func (c *UserController) GetItems() {
	c.Data["json"] = models.GetAllUsers()
	c.ServeJSON()
}
func (uc *UserController) AddNewUser() {
	var u models.Users
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	user := models.InsertOneUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}
func (uc *UserController) UpdateUser() {
	var u models.Users
	json.Unmarshal(uc.Ctx.Input.RequestBody, &u)
	user := models.UpdateUser(u)
	uc.Data["json"] = user
	uc.ServeJSON()
}
func (c *UserController) UpdateItem() {
	id, _ := c.GetInt(":id")
	var updatedItem models.Users
	json.Unmarshal(c.Ctx.Input.RequestBody, &updatedItem)
	updatedItem.Id = id
	o := orm.NewOrm()
	_, err := o.Update(&updatedItem)
	if err == nil {
		c.Data["json"] = map[string]interface{}{"message": "Item updated successfully", "item": updatedItem}
	} else {
		c.Data["json"] = map[string]interface{}{"error": err.Error()}
	}
	c.ServeJSON()
}
func (uc *UserController) DeleteUser() {
	id, _ := strconv.Atoi(uc.Ctx.Input.Param(":id"))
	deleted := models.DeleteUser(id)
	uc.Data["json"] = map[string]bool{"deleted": deleted}
	uc.ServeJSON()
}
