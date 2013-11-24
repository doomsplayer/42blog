package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
	"time"
)

func init() {
	fmt.Print(``)

}

type Article struct {
	OId          bson.ObjectId `bson:"_id"`
	Id           string        `bson:"-"`
	Name         string
	Content      string
	CreateTime   time.Time
	ModifiedTime time.Time
	Tags         []string
	Categories   []string
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
	err = ArticleCollection.Find(nil).All(&articles)
	e(err)
	for _, v := range articles {
		v.Id = fmt.Sprintf("%x", string(v.OId))
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
	err = ArticleCollection.Find(bson.M{}).Sort(`$nature`).Skip(i - 1).Limit(j - i + 1).All(&articles)

	e(err)
	for _, v := range articles {
		v.Id = fmt.Sprintf("%x", string(v.OId))
	}
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
	e(ArticleCollection.Find(bson.M{"name": name}).One(&article))
	article.Id = fmt.Sprintf("%x", string(article.OId))
	return
}

func (articleCltn) ReadArticleById(id string) (article *Article, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticleByIdFailed: `, err)
		}
	}()
	beego.Trace(`ReadArticleById: `, id)
	article = new(Article)
	e(ArticleCollection.Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&article))
	article.Id = fmt.Sprintf("%x", string(article.OId))
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
	switch {
	case article.OId == `` && article.Id != ``:
		{
			article.OId = bson.ObjectIdHex(article.Id)
		}
	case article.OId != ``:
		{
			break
		}
	case article.OId == `` && article.Id == ``:
		{
			article.OId = bson.NewObjectId()
		}
	}

	err = ArticleCollection.Insert(article)
	e(err)
	return
}

func (articleCltn) InsertArticleString(name, content string, catagories []string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`InsertArticleString: `, name)
	a := new(Article)
	a.OId = bson.NewObjectId()
	a.Name = name
	a.Content = content
	a.Categories = catagories
	a.Tags = tags
	a.CreateTime = time.Now()
	a.ModifiedTime = a.CreateTime
	beego.Trace(`InsertArticleString-InsertArticle`)
	err = ArticleCollection.Insert(a)
	e(err)
	return

}

func (articleCltn) UpsertArticleStringByName(name, content string, categories []string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`UpsertArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`UpsertArticleString: `, name)

	a := new(Article)
	a.Name = name
	a.Content = content
	a.Categories = categories
	a.Tags = tags
	c, er := ArticleCollection.Find(bson.M{"name": name}).Count()
	e(er)
	if c != 0 {
		a.ModifiedTime = time.Now()
	} else {
		a.CreateTime = time.Now()
		a.ModifiedTime = a.CreateTime
	}

	beego.Trace(`UpsertArticleString-InsertArticle`)
	_, err = ArticleCollection.Upsert(bson.M{"name": name}, a)
	e(err)
	return
}

func (articleCltn) UpsertArticleStringById(id, name, content string, categories []string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`UpsertArticleByStringFailed: `, err)
		}
	}()
	beego.Trace(`UpsertArticleStringById: `, id)

	a := new(Article)
	a.OId = bson.ObjectIdHex(id)
	a.Name = name
	a.Content = content
	a.Categories = categories
	a.Tags = tags
	c, er := ArticleCollection.Find(bson.M{"name": name}).Count()
	e(er)
	if c != 0 {
		a.ModifiedTime = time.Now()
	} else {
		a.CreateTime = time.Now()
		a.ModifiedTime = a.CreateTime
	}
	beego.Trace(`UpsertArticleString-InsertArticle`)
	_, err = ArticleCollection.Upsert(bson.M{"_id": bson.ObjectIdHex(id)}, a)
	e(err)
	return
}

func (articleCltn) UpdateArticleStringById(id, name, content string, categories []string, tags []string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`UpdateArticleByStringFailed: `, err)
		}
	}()

	beego.Trace(`UpdateArticleString: `, name)

	a := new(Article)
	a.Id = id
	a.OId = bson.ObjectIdHex(id)
	a.Name = name
	a.Content = content
	a.Categories = categories
	a.Tags = tags
	a.ModifiedTime = time.Now()
	beego.Trace(`UpdateArticleString-InsertArticle`)
	e(ArticleCollection.Update(bson.M{"_id": id}, a))
	return
}

func (articleCltn) DeleteArticleById(id string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadArticleByIdFailed: `, err)
		}
	}()
	e(ArticleCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)}))

	return
}
