package models

import (
//"github.com/astaxie/beego"
//"labix.org/mgo/bson"
)

type Tag struct {
	Id       int
	Name     string     `orm:"unique"`
	Articles []*Article `orm:"reverse(many)"`
}
type tagCltn struct{}

var TagCltn tagCltn

//func (tagCltn) InsertTag(tag *Tag) (err error) {
//	defer func() {
//		if ret := recover(); ret != nil {
//			err = ret.(error)
//			beego.Error(`InsertTagFailed`, err)
//		}
//	}()
//	beego.Trace(`InsertTag: `, tag.Name)
//	_, err = O.Insert(tag)

//	e(err)
//	m2m := O.QueryM2M(tag, `articles`)
//	_, err = m2m.Add(tag.Articles)
//	e(err)

//	return

//}

//func (tagCltn) InsertTagByName(name string) (err error) {
//	defer func() {
//		if ret := recover(); ret != nil {
//			err = ret.(error)
//			beego.Error(`InsertTagFailed`, err)
//		}
//	}()
//	beego.Trace(`InsertTagByName: `, name)
//	_, err = O.Insert(&Tag{Name: name})
//	e(err)
//	return

//}

//func (tagCltn) ReadAllTags() (tags []*Tag, err error) {
//	defer func() {
//		if ret := recover(); ret != nil {
//			err = ret.(error)
//			beego.Error(`ReadAllTagsFailed`, err)
//		}
//	}()
//	beego.Trace(`ReadAllTags`)
//	_, err = O.QueryTable(`tag`).All(&tags)
//	e(err)
//	for _, v := range tags {
//		_, err = O.LoadRelated(v, `articles`)
//		e(err)

//	}
//	return
//}

//func (tagCltn) ReadTagByName(name string) (tag *Tag, err error) {
//	defer func() {
//		if ret := recover(); ret != nil {
//			err = ret.(error)
//			beego.Error(`ReadTagByNameFailed`, err)
//		}
//	}()
//	beego.Trace(`ReadTagByName: `, name)
//	tag = new(Tag)
//	tag.Name = name
//	err = O.Read(tag, `name`)
//	e(err)
//	_, err = O.LoadRelated(tag, `articles`)
//	e(err)
//	return
//}
