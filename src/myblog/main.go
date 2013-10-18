package main

import (
	"github.com/astaxie/beego"
	_ "myblog/controllers"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/logs.log"}`)
	beego.SetLevel(beego.LevelTrace)
}
func main() {
	beego.Info(`main program start`)
	beego.Run()
}
