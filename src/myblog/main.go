package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	_ "myblog/controllers"
	"os"
)

func init() {
	beego.SetLogger("file", `{"filename":"logs/logs.log"}`)
	conf, _ := config.NewConfig(`ini`, `conf/custom.conf`)
	switch conf.String(`loglevel`) {
	case `trace`:
		{
			beego.SetLevel(beego.LevelTrace)
		}
	case `debug`:
		{
			beego.SetLevel(beego.LevelDebug)
		}
	case `info`:
		{
			beego.SetLevel(beego.LevelInfo)
		}
	case `warn`:
		{
			beego.SetLevel(beego.LevelWarning)
		}
	case `error`:
		{
			beego.SetLevel(beego.LevelError)
		}
	case `critical`:
		{
			beego.SetLevel(beego.LevelCritical)
		}
	default:
		{
			fmt.Println(`LoadConfigError`)
			os.Exit(1)
		}
	}
}

func main() {
	beego.Info(`main program start`)
	beego.Run()
}
