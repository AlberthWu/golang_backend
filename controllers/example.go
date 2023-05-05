package controllers

import (
	"encoding/json"
	"golang_backend/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (c *UserController) GetAll() {
	o := orm.NewOrm()

	var users []models.User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		logs.Error("Error retrieving users: %s", err)
		c.Abort("500")
	}

	c.Data["json"] = users
	c.ServeJSON()
}

func (c *UserController) Create() {
	o := orm.NewOrm()

	var user models.User
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	if err != nil {
		logs.Error("Invalid user data: %s", err)
		c.Abort("400")
	}

	id, err := o.Insert(&user)
	if err != nil {
		logs.Error("Error creating user: %s", err)
		c.Abort("500")
	}

	c.Data["json"] = map[string]int64{"id": id}
	c.ServeJSON()
}
