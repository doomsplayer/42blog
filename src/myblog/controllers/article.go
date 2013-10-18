package controllers

import (
	"fmt"
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
	beego.Router(`/article.php`, &article{})
}

func (this *article) Prepare() {
	v := this.GetSession(`admin_logined`)
	if v != nil {
		this.Data[`logined`] = true
	}
	this.Layout = `layout.html`
}

func (this *article) Get() {

	article_id, err := this.GetInt(`id`)
	if err != nil {
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

	article, err := models.ArticleCltn.ReadArticleById(int(article_id))
	if err != nil {
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
	this.Data[`title`] = `42的小站-文章`
	this.Data[`Content`] = template.HTML(blackfriday.MarkdownCommon([]byte(article.Content)))
	this.Data[`Article`] = article
	this.TplNames = `article-sub.html`

}

func (this *article) Post() {
	switch this.GetString(`type`) {
	case `add_comment`:
		{
			c := struct {
				Article_id int    `valid:"Required"form:"article_id,text"`
				Author     string `valid:"Required;MinSize(1);MaxSize(128)"form:"author,text"`
				Comment    string `valid:"Required;MinSize(1);MaxSize(1024)"form:"comment,text"`
				Email      string `valid:"Email"form:"email,text"`
			}{}
			err := this.ParseForm(&c)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			valid := validation.Validation{}
			b, err := valid.Valid(&c)
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
			err = models.CommentCltn.InsertCommentToArticle(c.Author, c.Email, c.Comment, c.Article_id)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(this.Ctx.Request.URL.String()+`?id=`+fmt.Sprint(c.Article_id), 302)
		}
	case `add_article`:
		{
			a := struct {
				Title    string `valid:"MinSize(1);MaxSize(255)"form:"title,text"`
				Content  string `valid:"Required"form:"content,text"`
				TagS     string `valid:"Match(/^[^ ]*(?: .+)* *?$/)"form:"tags,text"`
				Category string `valid:""form:"category,text"`
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
			tags := strings.Fields(a.TagS)
			err = models.ArticleCltn.UpsertArticleString(a.Title, a.Content, a.Category, tags)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(this.Ctx.Request.Referer(), 302)

		}
	case `del_article`:
		{
			id, err := this.GetInt(`id`)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			err = models.ArticleCltn.DeleteArticleById(int(id))
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(`/article.php`, 302)
		}
	case `edit_article`:
		{
			a := struct {
				Id       int    `valid:"Required"form:"id,text"`
				Title    string `valid:"MinSize(1);MaxSize(255)"form:"title,text"`
				Content  string `valid:"Required"form:"content,text"`
				Tags     string `valid:"Match(/^[^ ]*(?: .+)* *?$/)"form:"tags,text"`
				Category string `valid:""form:"category,text"`
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
			err = models.ArticleCltn.UpdateArticleString(a.Id, a.Title, a.Content, a.Category, tags)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			this.Redirect(`/article.php?id=`+fmt.Sprint(a.Id), 302)
		}
	case `del_comment`:
		{
			id, err := this.GetInt(`id`)
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}
			err = models.CommentCltn.DelCommentById(int(id))
			if err != nil {
				this.TplNames = `error.html`
				this.Data[`error`] = err
				return
			}

			this.Redirect(this.Ctx.Request.Referer(), 302)
		}
	default:
		{
			this.Redirect(this.Ctx.Request.URL.String(), 302)
		}
	}
}
