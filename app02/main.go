package main

import (
	"app02/controllers"
	"app02/db"
	"app02/log"
	repoimpl "app02/reponsitory/repo_impl"
	"app02/router"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

func init() {
	fmt.Println("init package main")
	os.Setenv("APP_NAME", "github")
	log.InitLogger(false)
}
func main() {
	sql := &db.Sql{
		Host:     "localhost",
		Port:     5432,
		UserName: "tusha",
		Password: "postgres",
		DbName:   "app02",
	}
	sql.Connect()
	defer sql.Close()
	e := echo.New()
	loginout := controllers.Loginout{
		UserRepo: repoimpl.NewUserRepo(sql),
	}
	app := router.API{
		Echo:     e,
		Loginout: loginout,
	}
	app.SetUpRouter()
	e.Logger.Fatal(e.Start(":3000"))
}
