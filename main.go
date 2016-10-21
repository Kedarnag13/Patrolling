package main

import (
	_ "github.com/Kedarnag13/Patrolling/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	o := orm.NewOrm()
	o.Using("default")
	o.RegisterDriver("postgres", o.DR_Postgres)
	o.RegisterDataBase("default", "postgres", "postgres:postgres/patrolling_development?charset=utf8")

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
