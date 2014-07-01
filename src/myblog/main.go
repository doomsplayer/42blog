package main

import (
	"flag"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "myblog/controllers"
)

var loglevel = flag.Int("ll", 0, "loglevel")
var filelog = flag.String("fl", "", "filelog directory")

func init() {

}

func main() {
	flag.Parse()
	if *filelog != `` {
		beego.SetLogger("file", `{"filename":"`+*filelog+`"}`)
	}
	beego.SetLevel(*loglevel)

	beego.Info(`Main Program Start`)

	beego.Run()
}
