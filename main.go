package main

import (
	"fmt"
	"github.com/Kedarnag13/Patrolling/models"
	_ "github.com/Kedarnag13/Patrolling/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func init() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=postgres dbname=patrolling_development sslmode=disable")
	orm.RegisterModel(new(models.Users), new(models.Sessions), new(models.Devises))
}

func main() {
	name := "default"

	// Drop table and re-create.
	force := true

	// Print log.
	verbose := true

	// Error.
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
