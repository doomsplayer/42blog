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

type Index struct {
	beego.Controller
}

func init() {
	beego.Router("/index", &Index{})
	beego.Router(`/exit`, &Index{}, `get,post:Exit`)
	beego.Router(`/`, &Index{}, `get:Index;post:Login`)
	beego.Router(`/weather`, &Index{}, `get:Weather`)
}

func (this *Index) Prepare() {
	this.Data[`viewtoday`] = models.ViewCountCltn.GetTodayView()
	this.Data[`viewall`] = models.ViewCountCltn.GetAllView()
	this.Layout = `layout.html`
}

func (this *Index) Index() {
	this.Redirect(`/index`, 302)
}

func (this *Index) Login() {
	if this.GetString(`passwd`) == key {
		this.SetSession(`admin_logined`, true)
	}
	this.Redirect(this.Ctx.Request.Referer(), 302)
}

func (this *Index) Exit() {
	this.DelSession(`admin_logined`)
	this.Redirect(this.Ctx.Request.Referer(), 302)
}

func (this *Index) Get() {
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

	geocode := `101270101`
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

func (this *Index) Weather() {
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
	this.Layout = ``
	this.TplNames = `weatherpanel.html`
}
