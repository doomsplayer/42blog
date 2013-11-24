package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
)

type about struct {
	beego.Controller
}

func init() {
	a := new(about)
	beego.Router(`/about`, a)
}
func (this *about) Prepare() {
	this.Layout = `layout.html`
}
func (this *about) Get() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Data[`viewtoday`] = models.ViewCountCltn.GetTodayView()
	this.Data[`viewall`] = models.ViewCountCltn.GetAllView()
	this.Data[`pos`] = `about`
	this.Data[`title`] = `42的小站-关于`
	this.TplNames = `about.html`
}
