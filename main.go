package main

import (
	_ "beego_api/routers"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	orm.RegisterDriver("postgres", orm.DRPostgres)
	dbUser := beego.AppConfig.String("db_user")
	dbPass := beego.AppConfig.String("db_pass")
	dbHost := beego.AppConfig.String("db_host")
	dbPort := beego.AppConfig.String("db_port")
	dbName := beego.AppConfig.String("db_name")
	conn := "user=" + dbUser + " password=" + dbPass + " host=" + dbHost + " port=" + dbPort + " dbname=" + dbName + " sslmode=disable"
	beego.Info("Database Connection String:", conn)

	regDB := orm.RegisterDataBase("default", "postgres", conn)
	if regDB != nil {
		print("Failed to register database:", regDB)
	}
	name := "default"
	force := false
	verbose := true
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	beego.Run()
}
