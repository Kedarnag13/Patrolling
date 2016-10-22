package controllers

import (
	"encoding/json"
	"github.com/Kedarnag13/Patrolling/models"
	"github.com/astaxie/beego"
)

// Operations about Users
type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post() {
	var user models.Users
	json.Unmarshal(r.Ctx.Input.RequestBody, &user)
	created_user := models.Save(user)
	r.Data["json"] = created_user
	r.ServeJSON()
}
