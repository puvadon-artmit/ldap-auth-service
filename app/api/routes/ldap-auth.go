package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tomioka/ldap-auth-service/internal/handlers"
)

func RoutesAuth() *fiber.App {

	app := fiber.New()

	app.Post("/domain-login", handlers.LdapAuthHandler())

	return app
}
