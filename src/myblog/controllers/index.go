package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/doomsplayer/sinaIp2Geo"
	"github.com/doomsplayer/weatherCN"
	"myblog/models"
)

func init() {
	cfg, err := config.NewConfig(`ini`, `conf/custom.conf`)
	if err != nil {
		panic(err)
	}
	key = cfg.String(`siteadminpasswd`)
}

var key = `12345`

type Root struct {
	beego.Controller
}

func init() { beego.Router(`/`, &Root{}) }

func (this *Root) Get() {
	this.Redirect(`/index.asp`, 302)
}

func (this *Root) Post() {
	if this.GetString(`passwd`) == key {
		this.SetSession(`admin_logined`, true)
	}
	this.Redirect(`/`, 302)
}

type exit struct {
	beego.Controller
}

func init() {
	beego.Router(`/exit`, &exit{})
}
func (this *exit) Get() {
	this.DelSession(`admin_logined`)
	this.Redirect(`/`, 302)
}

type IndexController struct {
	beego.Controller
}

func init() {
	beego.Router("/index.asp", &IndexController{})
}

func (this *IndexController) Prepare() {
	this.Layout = `layout.html`
}

func (this *IndexController) Get() {
	//models.UpsertArticleString(`LifetimeExplanationForRust`, ``, `Rust`, []string{})
	//models.InsertTsukkomiWithContent(`CSS真是史上最烂发明`)
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	articles, err := models.ArticleCltn.ReadArticlesByTimeRange(1, 8)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}

	tsukkomis, err := models.TsukkomiCltn.ReadTsukkomiByTimeRange(1, 5)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}

	var ip string
	if ip = this.Ctx.Request.Header.Get(`X-Forwarded-For`); ip == `` {
		ip = this.Ctx.Input.IP()
	}

	ig, err := sinaIp2Geo.New(ip)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	err = ig.Parse()
	if err != nil {
		ig.RetJson.City = `成都`
	}

	geocode, err := weather.Geo2Code(ig.RetJson.City)
	if err != nil {
		geocode = `101270101`
	}

	wg := weather.New()
	wg.SetACode(geocode)

	ret, _ := wg.GetInfo()

	this.Data[`weather`] = ret
	this.Data[`Tsukkomis`] = tsukkomis
	this.Data[`Articles`] = articles
	this.Data[`title`] = `42的小站-主页`
	this.Data[`pos`] = `home`
	this.TplNames = "index.html"
}
