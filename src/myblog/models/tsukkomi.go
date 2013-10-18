package models

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Tsukkomi struct {
	Id      int
	Content string
	Time    time.Time `orm:"auto_now_add"`
}

type tsukkomiCltn struct{}

var TsukkomiCltn tsukkomiCltn

func (tsukkomiCltn) ReadAllTsukkomi() (tsukkomi []*Tsukkomi, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadAllTsukkomiFailed`, err)
		}
	}()
	beego.Trace(`ReadAllTsukkomis`)
	_, err = O.QueryTable(`tsukkomi`).All(&tsukkomi)
	return
}

func (tsukkomiCltn) ReadTsukkomiByTimeRange(i, j int) (tsukkomis []*Tsukkomi, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadTsukkomisByTimeRangeFailed: `, err)

		}
	}()
	beego.Trace(`ReadTsukkomiByTimeRange: `, i, `to`, j)
	if i > j || i < 1 || j < 1 {
		panic(`index number error`)
	}
	_, err = O.QueryTable(`tsukkomi`).Offset(i - 1).Limit(j - i + 1).OrderBy(`-time`).All(&tsukkomis)
	e(err)
	return
}

func (tsukkomiCltn) InsertTsukkomi(tsukkomi *Tsukkomi) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertTsukkomiFailed`, err)
		}
	}()
	beego.Trace(`InsertTsukkomi: `, tsukkomi.Content)
	_, err = O.Insert(tsukkomi)
	e(err)
	return

}

func (tsukkomiCltn) InsertTsukkomiWithContent(content string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`InsertTsukkomiFailed:`, err)
		}
	}()
	beego.Trace(`InsertTsukkomiWithContent: `, content)
	if content == `` {
		panic(fmt.Errorf("Content is Wmpty"))
	}
	_, err = O.Insert(&Tsukkomi{Content: content})
	e(err)
	return

}

func (tsukkomiCltn) DeleteTsukkomiById(id int) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`DeleteTsukkomiFailed:`, err)
		}
	}()
	beego.Trace(`DeleteTsukkomiById: `, id)
	_, err = O.QueryTable(`tsukkomi`).Filter(`id`, id).Delete()
	e(err)
	return
}
