package controllers

import (
	"encoding/json"
	"github.com/Kedarnag13/Patrolling/models"
	"github.com/astaxie/beego"
)

type SignInController struct {
	beego.Controller
}

func (s *SignInController) Post() {
	var user models.Sessions
	json.Unmarshal(s.Ctx.Input.RequestBody, &user)

}
