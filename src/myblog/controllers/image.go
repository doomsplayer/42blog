package controllers

import (
	"github.com/astaxie/beego"
	"myblog/models"
	// "myblog/models"
)

type image struct {
	beego.Controller
}

func init() {
	beego.Router(`/image.php`, &image{})
}

func (this *image) Prepare() {
	this.Layout = `layout.html`
}

func (this *image) Get() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Data[`viewtoday`] = models.ViewCountCltn.GetTodayView()
	this.Data[`viewall`] = models.ViewCountCltn.GetAllView()
	this.Data[`title`] = `42的小站-照片`
	this.TplNames = `image.html`
}
