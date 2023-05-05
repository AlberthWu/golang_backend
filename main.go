// package main

// import (
// 	_ "golang_backend/models"
// 	_ "golang_backend/routers"

// 	// "github.com/beego/beego/v2/core/logs"

// 	// "github.com/beego/beego/v2/client/orm"
// 	beego "github.com/beego/beego/v2/server/web"
// )

// func main() {
// 	// dbType, _ := beego.AppConfig.String("db_type")
// 	// dbAlias, _ := beego.AppConfig.String(dbType + "::db_alias")
// 	// dbName, _ := beego.AppConfig.String(dbType + "::db_name")
// 	// dbUser, _ := beego.AppConfig.String(dbType + "::db_user")
// 	// dbPwd, _ := beego.AppConfig.String(dbType + "::db_pwd")
// 	// dbHost, _ := beego.AppConfig.String(dbType + "::db_host")
// 	// dbPort, _ := beego.AppConfig.String(dbType + "::db_port")
// 	// dbCharset, _ := beego.AppConfig.String(dbType + "::db_charset")

// 	// logs.Info("Check :", dbAlias, "Check 2 :", dbType, "Check 3 :", beego.BConfig.RunMode)

// 	// orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
// 	// 	dbPort+")/"+dbName+"?charset="+dbCharset)

//		if beego.BConfig.RunMode == "dev" {
//			beego.BConfig.WebConfig.DirectoryIndex = true
//			beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
//		}
//		beego.Run()
//	}
package main

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"

	_ "golang_backend/routers"
)

func main() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(localhost:3306)/mydb?charset=utf8mb4&parseTime=true&loc=Local")
	orm.RunSyncdb("default", false, true)

	// Start the Beego application
	web.Run()
}
