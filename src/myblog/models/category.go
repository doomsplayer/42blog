package models

import (
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id       int
	Name     string     `orm:"unique"`
	Articles []*Article `orm:"reverse(many)"`
}

type categoryCltn struct{}

var CategoryCltn categoryCltn

func (categoryCltn) InsertCategory(category *Category) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertCategoryFailed`, err)
		}
	}()
	beego.Trace(`InsertCategory: `, category.Name)
	_, err = O.Insert(category)
	e(err)
	return

}

func (categoryCltn) InsertCategoryByName(name string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertCategoryFailed`, err)
		}
	}()
	beego.Trace(`InsertCategoryByName: `, name)
	_, err = O.Insert(&Category{Name: name})
	e(err)
	return

}

func (categoryCltn) ReadAllCatagories() (catagories []*Category, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadAllCatagoriesFailed`, err)
		}
	}()
	beego.Trace(`ReadAllCatagories`)
	_, err = O.QueryTable(`category`).All(&catagories)
	e(err)
	return
}

func (categoryCltn) ReadCategoryByName(name string) (category *Category, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadCategoryByNameFailed`, err)
		}
	}()
	beego.Trace(`ReadCategoryByName: `, name)
	category = new(Category)
	category.Name = name
	err = O.Read(category, `name`)
	e(err)
	return
}
