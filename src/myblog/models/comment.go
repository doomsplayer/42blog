package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Comment struct {
	Id      int
	Article *Article `orm:"rel(fk)"`
	Content string   `orm:"type(text)"`
	Author  string
	Email   string
	Time    time.Time `orm:"auto_now_add"`
}

type commentCltn struct{}

var CommentCltn commentCltn

func (commentCltn) ReadAllCommentsByArticleId(id int) (comments []*Comment, err error) {
	return nil, fmt.Errorf(`todo`)
}
func (commentCltn) InsertComment(comment *Comment) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertCommentFailed`, err)
		}
	}()
	_, err = O.Insert(comment)
	e(err)
	return
}

func (commentCltn) DelCommentById(id int) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`DelCommentByIdFailed`, err)
		}
	}()
	beego.Trace(`DelCommentById: `, id)
	c := &Comment{Id: id}
	_, err = O.Delete(c)
	e(err)
	return
}

func (commentCltn) DelCommentByAuthor(author string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`DelCommentByAuthorFailed`, err)
		}
	}()
	beego.Trace(`DelCommentByAuthor: `, author)
	_, err = O.QueryTable(`comment`).Filter(`author`, author).Delete()
	e(err)
	return
}

func (commentCltn) InsertCommentToArticle(author, email, comment string, id int) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertCommentFailed`, err)
		}
	}()
	c := Comment{}
	c.Article = &Article{Id: id}
	c.Author = author
	c.Content = comment
	c.Email = email
	_, err = O.Insert(&c)
	e(err)
	return
}
