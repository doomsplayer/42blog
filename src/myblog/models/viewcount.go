package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type ViewCounter struct {
	Name      string
	ViewToday int
	ViewAll   int
	Today     time.Time
}

func init() {

}

type viewCountCltn struct{}

var ViewCountCltn viewCountCltn

func (viewCountCltn) GetTodayView() int {
	c := &ViewCounter{}
	err := ViewCountCollection.Find(bson.M{"name": "viewCount"}).One(c)
	if err != nil {
		return 0
	}
	return c.ViewToday

}

func (viewCountCltn) GetAllView() int {
	c := &ViewCounter{}
	err := ViewCountCollection.Find(bson.M{"name": "viewCount"}).One(c)
	if err != nil {
		return 0
	}
	return c.ViewAll

}

func (viewCountCltn) IncrView() {
	defer func() {
		//recover()
	}()
	if n, _ := ViewCountCollection.Find(bson.M{"name": "viewCount"}).Count(); n == 0 {
		e(ViewCountCollection.Insert(ViewCounter{"viewCount", 0, 0, time.Now()}))
	}
	c := &ViewCounter{}
	e(ViewCountCollection.Find(bson.M{"name": "viewCount"}).One(c))
	if time.Now().Day() != c.Today.Day() || time.Now().Month() != c.Today.Month() || time.Now().Year() != c.Today.Year() {
		ViewCountCollection.Update(bson.M{"name": "viewCount"}, bson.M{"viewtoday": 0, "today": time.Now()})

	}
	e(ViewCountCollection.Update(bson.M{"name": "viewCount"}, bson.M{"$inc": bson.M{"viewtoday": 1}}))
	e(ViewCountCollection.Update(bson.M{"name": "viewCount"}, bson.M{"$inc": bson.M{"viewall": 1}}))
}
