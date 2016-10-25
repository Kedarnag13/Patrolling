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
	var session models.Sessions
	json.Unmarshal(s.Ctx.Input.RequestBody, &session)
	s.Data["json"] = models.CreateSession(session)
	s.ServeJSON()
}
