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
	var user models.User
	json.Unmarshal(r.Ctx.Input.RequestBody, &user)
	models.Save(user)
	// r.Data["json"] = map[string]string{"saved_user": saved_user}
	// r.ServeJSON()
}
