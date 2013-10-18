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
	beego.Router(`/tsukkomi.php`, &tsukkomi{})
}

func (this *tsukkomi) Prepare() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Layout = `layout.html`
}

func (this *tsukkomi) Get() {

	this.TplNames = `tsukkomi.html`
	// this.TplNames = `tsukkomi-sub.html`
}
func (this *tsukkomi) Post() {
	switch this.GetString(`type`) {
	case `add_tsukkomi`:
		{
			err := models.TsukkomiCltn.InsertTsukkomiWithContent(this.GetString(`content`))
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(this.Ctx.Request.Referer(), 302)
		}
	case `del_tsukkomi`:
		{
			id, err := this.GetInt(`id`)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			err = models.TsukkomiCltn.DeleteTsukkomiById(int(id))
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(this.Ctx.Request.Referer(), 302)
		}
	default:
		{
			this.Redirect(this.Ctx.Request.Referer(), 302)
		}
	}
}
