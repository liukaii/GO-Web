package main

import (
	_ "GO-Web/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.StaticDir["/static"] = "static"
	beego.Run()
}

