package controllers

import (
	"fmt"
	"myblog/models"
	// "fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.AddFilter("*", "BeforRouter", Logger)
	beego.AddFuncMap(`sprinttags`, SprintTags)
}

func Logger(ctx *context.Context) {
	if fip := ctx.Request.Header.Get(`X-Forwarded-For`); fip != `` {
		beego.Info(fip, ctx.Input.Method(), ctx.Input.Url())
	} else {
		beego.Info(ctx.Input.IP(), ctx.Input.Method(), ctx.Input.Url())
	}

}

func SprintTags(in []*models.Tag) string {
	out := ``
	for _, v := range in {
		out += (fmt.Sprint(v.Name) + ` `)
	}
	return out[:len(out)]
}
