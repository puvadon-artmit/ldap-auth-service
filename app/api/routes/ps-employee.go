package routes

import (
	"github.com/tomioka/ldap-auth-service/internal/core/services"
	"github.com/tomioka/ldap-auth-service/internal/handlers"
	"github.com/tomioka/ldap-auth-service/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RoutesEmployee(db *gorm.DB) *fiber.App {
	if db == nil {
		panic("Database connection is nil")
	}

	app := fiber.New()

	EmployeeRepository := repositories.NewEmployeeRepositoryDB(db)
	EmployeeService := services.NewEmployeeService(EmployeeRepository)
	EmployeeHandler := handlers.NewEmployeeHandler(EmployeeService)

	app.Get("/find-employee-by-account", EmployeeHandler.FindEmployeeByAccount())
	app.Get("/get-employee-by-emp-code", EmployeeHandler.GetEmployeeByEmpCode())

	return app
}
