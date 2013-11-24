package models

import (
	"fmt"
	"labix.org/v2/mgo"
)

var db *mgo.Database
var ArticleCollection *mgo.Collection
var TsukkomiCollection *mgo.Collection
var ViewCountCollection *mgo.Collection

func init() {
	fmt.Print(``)
	session, err := mgo.Dial("localhost")
	e(err)
	db = session.DB(`42blog`)
	ArticleCollection = db.C(`article`)
	func() {
		found := false
		ids, err := ArticleCollection.Indexes()
		e(err)
	O:
		for _, v := range ids {

			for _, u := range v.Key {

				if u == "name" {
					found = true
					break O
				}

			}

		}
		if !found {
			e(ArticleCollection.EnsureIndex(mgo.Index{Key: []string{"name"}, Unique: true}))
		}
	}()
	ViewCountCollection = db.C("meta")
	TsukkomiCollection = db.C(`tsukkomi`)
}
func e(err error) {
	if err != nil {
		panic(err)
	}
}
