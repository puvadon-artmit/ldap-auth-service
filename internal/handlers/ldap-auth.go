package handlers

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/tomioka/ldap-auth-service/internal/core/models"
	"github.com/tomioka/ldap-auth-service/internal/ldapauth"
)

func LdapAuthHandler() fiber.Handler {

	ldapSvc, err := ldapauth.NewService()
	if err != nil {
		log.Fatalf("failed to init ldap service: %v", err)
	}

	return func(c *fiber.Ctx) error {
		var req models.AuthRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "invalid request body",
			})
		}

		if req.Username == "" || req.Password == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"error":   "username and password are required",
			})
		}

		ok, msg := ldapSvc.Authenticate(req.Username, req.Password)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"success": false,
				"error":   msg,
			})
		}

		return c.JSON(fiber.Map{
			"success": true,
			"message": "authenticated",
		})
	}
}
