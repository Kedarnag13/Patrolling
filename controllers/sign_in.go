package controllers

import (
	"encoding/json"
	"github.com/Kedarnag13/Patrolling/models"
	"github.com/astaxie/beego"
)

type SessionController struct {
	beego.Controller
}

func (s *SessionController) Post() {
	var user models.Sessions
	json.Unmarshal(r.Ctx.Input.RequestBody, &user)

}
