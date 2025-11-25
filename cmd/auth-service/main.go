package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"github.com/tomioka/ldap-auth-service/internal/ldapauth"
)

type AuthRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	// โหลด .env (ถ้ารันใน docker ที่ตั้ง env แล้วจะไม่จำเป็น)
	_ = godotenv.Load()

	app := fiber.New()

	ldapSvc, err := ldapauth.NewService()
	if err != nil {
		log.Fatalf("failed to init ldap service: %v", err)
	}

	app.Post("/auth/domain-login", func(c *fiber.Ctx) error {
		var req AuthRequest
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
			// "token":   "xxx", // ใส่เพิ่มถ้าต้องการ
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	log.Printf("Auth service listening on :%s", port)
	log.Fatal(app.Listen(":" + port))
}
