package main

import (
	_ "github.com/Piasy/BeegoBootStrap/docs"
	_ "github.com/Piasy/BeegoBootStrap/routers"

	"github.com/astaxie/beego"
)

func main() {
	if beego.RunMode == "dev" {
		beego.DirectoryIndex = true
		beego.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
