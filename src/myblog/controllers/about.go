package controllers

import (
	"github.com/astaxie/beego"
)

type about struct {
	beego.Controller
}

func init() {
	beego.Router(`/about.asp`, &about{})
}
func (this *about) Prepare() {
	this.Layout = `layout.html`
}
func (this *about) Get() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Data[`pos`] = `about`
	this.Data[`title`] = `42的小站-关于`
	this.TplNames = `about.html`
}
