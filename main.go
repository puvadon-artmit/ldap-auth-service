package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	routes "github.com/tomioka/ldap-auth-service/app/api"
	database "github.com/tomioka/ldap-auth-service/app/external/db"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/tomioka/ldap-auth-service/app/configs"
	"github.com/tomioka/ldap-auth-service/app/middlewares"
	"github.com/tomioka/ldap-auth-service/internal/pkgs/logs"
)

var (
	db *gorm.DB
)

func init() {
	configs.Init()
	logs.LogInit()
	db = database.InitDataBase()
}

func main() {
	app := fiber.New(fiber.Config{
		AppName:   "atelnord",
		BodyLimit: 50 * 1024 * 1024,
	})

	// go func() {
	// 	log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	// }()

	app.Use(
		middlewares.NewLoggerMiddleware,
		middlewares.NewCorsMiddleware,
	)

	routes.SetupRoutes(app, db)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		serv := <-c
		if serv.String() == "interrupt" {
			fmt.Println("Gracefully shutting down...")
			app.Shutdown()
		}
	}()

	err := app.Listen("0.0.0.0:" + os.Getenv("SERVER_PORT"))
	if err != nil {
		log.Fatal(err)
	}
}
