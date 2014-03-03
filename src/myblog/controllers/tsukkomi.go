package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	// "myblog/models"
)

type tsukkomi struct {
	beego.Controller
}

func init() {
	t := new(tsukkomi)
	beego.Router(`/tsukkomi/add`, t, `post:AddTsukkomi`)
	beego.Router(`/tsukkomi/:id/del`, t, `post:DelTsukkomi`)
	beego.Router(`/tsukkomi`, t)
}

func (this *tsukkomi) Prepare() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Data[`viewtoday`] = models.ViewCountCltn.GetTodayView()
	this.Data[`viewall`] = models.ViewCountCltn.GetAllView()
	this.Layout = `layout.html`
}

func (this *tsukkomi) Get() {

	this.TplNames = `tsukkomi.html`
	// this.TplNames = `tsukkomi-sub.html`
}
func (this *tsukkomi) AddTsukkomi() {
	v := this.GetSession(`admin_logined`)
	if v == nil {
		this.Redirect(`/`, 302)
	}
	err := models.TsukkomiCltn.InsertTsukkomiWithContent(this.GetString(`content`))
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	this.Redirect(this.Ctx.Request.Referer(), 302)
}
func (this *tsukkomi) DelTsukkomi() {
	v := this.GetSession(`admin_logined`)
	if v == nil {
		this.Redirect(`/`, 302)
	}
	id := this.Ctx.Input.Param(`:id`)
	err := models.TsukkomiCltn.DeleteTsukkomiById(id)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	this.Redirect(this.Ctx.Request.Referer(), 302)
}
