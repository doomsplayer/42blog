package controllers

import (
	"github.com/astaxie/beego"
)

type AuthController struct {
	beego.Controller
}

func (this *AuthController) Prepare() {
	if this.GetSession(`auth`) != nil {
		this.Data[`auth`] = true
	}
}
