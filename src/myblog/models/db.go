package models

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

func init() {
	fmt.Print(``)
	cfg, err := config.NewConfig(`ini`, `conf/custom.conf`)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(Article), new(Comment), new(Tag), new(Category), new(Tsukkomi))
	orm.RegisterDataBase("default", "mysql", cfg.String(`dbusername`)+`:`+cfg.String(`dbpasswd`)+`@tcp(localhost:`+cfg.String(`dbport`)+`)/`+cfg.String(`dbname`)+`?charset=utf8`, 30)
	orm.RunCommand()
	orm.Debug = true
	O = orm.NewOrm()
}
func e(err error) {
	if err != nil {
		panic(err)
	}
}

var Meta cache.Cache

func init() {
	var err error
	Meta, err = cache.NewCache(`redis`, `{"conn":"127.0.0.1:6379"}`)
	if err != nil {
		panic(err)
	}
}
