package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"github.com/russross/blackfriday"
	"html/template"
	"myblog/models"
	"strings"
)

type article struct {
	beego.Controller
}

func init() {
	a := new(article)
	beego.Router(`/article/add`, a, `post:AddArticle`)
	beego.Router(`/article/:id/del`, a, "post:DelArticle")
	beego.Router(`/article/:id/edit`, a, `post:EditArticle`)
	beego.Router(`/article/:id`, a)
	beego.Router(`/article`, a, "get:GetList")

}

func (this *article) Prepare() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Data[`viewtoday`] = models.ViewCountCltn.GetTodayView()
	this.Data[`viewall`] = models.ViewCountCltn.GetAllView()
	this.Layout = `layout.html`
}

func (this *article) GetList() {
	articles, err := models.ArticleCltn.ReadArticlesByTimeRange(1, 10)
	if err != nil {
		this.TplNames = "error.html"
		this.Data[`error`] = err
		return
	}
	this.Data[`title`] = `42的小站-文章列表`
	this.Data[`Articles`] = articles
	this.TplNames = "article.html"
	return
}

func (this *article) Get() {
	article_id := this.Ctx.Input.Param(`:id`)
	if article_id == `` {
		this.Redirect(`/article`, 302)
	}
	article, err := models.ArticleCltn.ReadArticleById(article_id)
	if err != nil {
		this.Redirect(`/article`, 302)
	}
	this.Data[`title`] = `42的小站-文章`
	this.Data[`Content`] = template.HTML(blackfriday.MarkdownCommon([]byte(article.Content)))
	this.Data[`Article`] = article
	this.TplNames = `article-sub.html`

}

func (this *article) AddArticle() {
	v := this.GetSession(`admin_logined`)
	if v == nil {
		this.Redirect(`/`, 302)
	}
	a := struct {
		Title      string `valid:"MinSize(1);MaxSize(255)"form:"title,text"`
		Content    string `valid:"Required"form:"content,text"`
		Tags       string `valid:"Match(/^[^ ]*(?: .+)* *?$/)"form:"tags,text"`
		Categories string `valid:""form:"categories,text"`
	}{}
	err := this.ParseForm(&a)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&a)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	if b != true {
		this.TplNames = `error.html`
		this.Data[`error`] = `输入有误`
		return
	}
	tags := strings.Fields(a.Tags)
	categories := strings.Fields(a.Categories)
	err = models.ArticleCltn.InsertArticleString(a.Title, a.Content, categories, tags)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	this.Redirect(this.Ctx.Request.Referer(), 302)

}

func (this *article) DelArticle() {
	v := this.GetSession(`admin_logined`)
	if v == nil {
		this.Redirect(`/`, 302)
	}
	id := this.Ctx.Input.Param(`:id`)
	err := models.ArticleCltn.DeleteArticleById(id)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	this.Redirect(`/article`, 302)

}

func (this *article) EditArticle() {
	v := this.GetSession(`admin_logined`)
	if v == nil {
		this.Redirect(`/`, 302)
	}
	id := this.Ctx.Input.Param(`:id`)
	a := struct {
		Title      string `valid:"MinSize(1);MaxSize(255)"form:"title,text"`
		Content    string `valid:"Required"form:"content,text"`
		Tags       string `valid:"Match(/^[^ ]*(?: .+)* *?$/)"form:"tags,text"`
		Categories string `valid:""form:"category,text"`
	}{}

	err := this.ParseForm(&a)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	valid := validation.Validation{}
	b, err := valid.Valid(&a)

	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	if b != true {
		this.TplNames = `error.html`
		this.Data[`error`] = `输入有误`
		return
	}
	tags := strings.Fields(a.Tags)
	categories := strings.Fields(a.Categories)
	err = models.ArticleCltn.UpdateArticleStringById(id, a.Title, a.Content, categories, tags)
	if err != nil {
		this.TplNames = `error.html`
		this.Data[`error`] = err
		return
	}
	this.Redirect(`/article/`+id, 302)
}
