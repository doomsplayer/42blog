package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

func init() {
	fmt.Print(``)
}

type Article struct {
	Id       int
	Name     string     `orm:"unique"`
	Content  string     `orm:"type(text)"`
	Time     time.Time  `orm:"auto_now"`
	Comments []*Comment `orm:"reverse(many)"`
	Tags     []*Tag     `orm:"rel(m2m)"`
	Category *Category  `orm:"rel(fk)"`
}

var ArticleCltn articleCltn

type articleCltn struct{}

func (articleCltn) ReadAllArticles() (articles []*Article, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadAllArticlesFailed: `, err)

		}
	}()
	beego.Trace(`ReadAllArticles`)
	_, err = O.QueryTable(`article`).All(&articles)
	e(err)
	for _, v := range articles {
		O.LoadRelated(v, `tags`)
		O.LoadRelated(v, `category`)
	}
	return
}

func (articleCltn) ReadArticlesByTimeRange(i, j int) (articles []*Article, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticlesByTimeRangeFailed: `, err)

		}
	}()
	beego.Trace(`ReadArticlesByTimeRange: `, i, `to`, j)
	if i > j || i < 1 || j < 1 {
		panic(`index number error`)
	}
	_, err = O.QueryTable(`article`).Offset(i - 1).Limit(j - i + 1).OrderBy(`-time`).All(&articles)
	e(err)
	for _, v := range articles {
		O.LoadRelated(v, `tags`)
		O.LoadRelated(v, `category`)
		O.LoadRelated(v, `comments`)
	}
	return
}
func (articleCltn) InsertArticle(article *Article) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`CreateArticleFailed: `, err)

		}
	}()
	beego.Trace(`InsertArticle: `, article.Name)
	_, err = O.Insert(article)
	e(err)
	m2m := O.QueryM2M(article, `tags`)
	_, err = m2m.Add(article.Tags)
	e(err)
	return
}

func (articleCltn) InsertArticleString(name, content, catagorie string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`InsertArticleString: `, name)
	beego.Trace(`InsertArticleString-ParseTags`)
	tag := []*Tag{}
	for _, v := range tags {
		t, err := TagCltn.ReadTagByName(v)
		e(err)
		tag = append(tag, t)
	}
	beego.Trace(`InsertArticleString-ParseCatagories`)
	cata, err := CategoryCltn.ReadCategoryByName(catagorie)
	e(err)

	a := new(Article)
	a.Name = name
	a.Content = content
	a.Category = cata
	beego.Trace(`InsertArticleString-InsertArticle`)
	_, err = O.Insert(a)
	e(err)
	m2m := O.QueryM2M(a, `tags`)
	_, err = m2m.Add(tag)
	e(err)
	return

}

func (articleCltn) UpsertArticleString(name, content, category string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`UpsertArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`UpsertArticleString: `, name)
	beego.Trace(`UpsertArticleString-ParseTags`)
	tag := []*Tag{}
	for _, v := range tags {
		t, err := TagCltn.ReadTagByName(v)
		if err == orm.ErrNoRows {
			err = TagCltn.InsertTagByName(v)
			e(err)
			t, err = TagCltn.ReadTagByName(v)
		}
		e(err)
		tag = append(tag, t)
	}
	beego.Trace(`UPsertArticleString-ParseCategory`)
	cata, err := CategoryCltn.ReadCategoryByName(category)
	if err == orm.ErrNoRows {
		err = CategoryCltn.InsertCategoryByName(category)
		e(err)
		cata, err = CategoryCltn.ReadCategoryByName(category)
	}
	e(err)

	a := new(Article)
	a.Name = name
	a.Content = content
	a.Category = cata
	a.Tags = tag
	beego.Trace(`UpsertArticleString-InsertArticle`)
	_, err = O.Insert(a)
	e(err)
	m2m := O.QueryM2M(a, `tags`)
	_, err = m2m.Add(tag)
	e(err)
	return

}

func (articleCltn) UpdateArticleString(id int, name, content, category string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`UpdateArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`UpdateArticleString: `, name)
	beego.Trace(`UpdateArticleString-ParseTags`)
	tag := []*Tag{}
	for _, v := range tags {
		t, err := TagCltn.ReadTagByName(v)
		if err == orm.ErrNoRows {
			err = TagCltn.InsertTagByName(v)
			e(err)
			t, err = TagCltn.ReadTagByName(v)
		}
		e(err)
		tag = append(tag, t)
	}
	beego.Trace(`UpdateArticleString-ParseCategory`)
	cata, err := CategoryCltn.ReadCategoryByName(category)
	if err == orm.ErrNoRows {
		err = CategoryCltn.InsertCategoryByName(category)
		e(err)
		cata, err = CategoryCltn.ReadCategoryByName(category)
	}
	e(err)

	a := new(Article)
	a.Id = id
	err = O.Read(a)
	e(err)
	a.Name = name
	a.Content = content
	a.Category = cata
	a.Tags = tag
	beego.Trace(`UpdateArticleString-InsertArticle`)
	_, err = O.Update(a)
	e(err)
	m2m := O.QueryM2M(a, `tags`)
	m2m.Clear()
	_, err = m2m.Add(tag)
	e(err)
	return
}

func (articleCltn) ReadArticleByName(name string) (article *Article, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticleByNameFailed: `, err)
		}
	}()
	beego.Trace(`ReadArticleByName: `, name)
	article = new(Article)
	article.Name = name
	err = O.Read(article, `name`)
	e(err)
	_, err = O.LoadRelated(article, `tags`)
	e(err)
	_, err = O.LoadRelated(article, `comments`)
	e(err)
	_, err = O.LoadRelated(article, `category`)
	e(err)
	return
}

func (articleCltn) ReadArticleById(id int) (article *Article, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticleByIdFailed: `, err)
		}
	}()
	beego.Trace(`ReadArticleById: `, id)
	article = new(Article)
	article.Id = id
	err = O.Read(article)
	e(err)
	_, err = O.LoadRelated(article, `tags`)
	e(err)
	_, err = O.LoadRelated(article, `comments`)
	e(err)
	_, err = O.LoadRelated(article, `category`)
	e(err)
	return
}

func (articleCltn) DeleteArticleById(id int) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticleByIdFailed: `, err)
		}
	}()
	a, err := ArticleCltn.ReadArticleById(id)
	e(err)
	m2m := O.QueryM2M(a, `tags`)
	m2m.Clear()
	_, err = O.Delete(a)
	e(err)
	return
}
