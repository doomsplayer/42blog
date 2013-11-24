package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"labix.org/v2/mgo/bson"
	"time"
)

func init() {

}

type Tsukkomi struct {
	Id      string        `bson:"-"`
	OId     bson.ObjectId `bson:"_id"`
	Content string
	Time    time.Time
}

type tsukkomiCltn struct{}

var TsukkomiCltn tsukkomiCltn

func (tsukkomiCltn) ReadAllTsukkomi() (tsukkomis []*Tsukkomi, err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`ReadAllTsukkomiFailed`, err)
		}
	}()
	beego.Trace(`ReadAllTsukkomis`)
	e(TsukkomiCollection.Find(nil).All(&tsukkomis))
	for _, v := range tsukkomis {
		v.Id = fmt.Sprintf("%x", string(v.OId))
	}
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
	e(TsukkomiCollection.Find(nil).Sort("$nature").Skip(i - 1).Limit(j - i + 1).All(&tsukkomis))
	for _, v := range tsukkomis {
		v.Id = fmt.Sprintf("%x", string(v.OId))
	}
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
	switch {
	case tsukkomi.OId == `` && tsukkomi.Id != ``:
		{
			tsukkomi.OId = bson.ObjectIdHex(tsukkomi.Id)
		}
	case tsukkomi.OId != ``:
		{
			break
		}
	case tsukkomi.OId == `` && tsukkomi.Id == ``:
		{
			tsukkomi.OId = bson.NewObjectId()
		}
	}
	e(TsukkomiCollection.Insert(tsukkomi))
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
	e(TsukkomiCollection.Insert(&Tsukkomi{Content: content, OId: bson.NewObjectId(), Time: time.Now()}))
	return

}

func (tsukkomiCltn) DeleteTsukkomiById(id string) (err error) {
	defer func() {
		if ret := recover(); ret != nil {
			err = ret.(error)
			beego.Error(`DeleteTsukkomiFailed:`, err)
		}
	}()
	beego.Trace(`DeleteTsukkomiById: `, id)
	e(TsukkomiCollection.Remove(bson.M{"_id": bson.ObjectIdHex(id)}))
	return
}
