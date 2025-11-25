package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	routes "github.com/tomioka/ldap-auth-service/app/api/routes"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	api := app.Group("/api", logger.New())
	api.Mount("/auth", routes.RoutesAuth())
	api.Mount("/employee", routes.RoutesEmployee(db))
}
