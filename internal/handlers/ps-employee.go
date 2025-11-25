package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	services "github.com/tomioka/ldap-auth-service/internal/core/ports/services"
)

type EmployeeHandler struct {
	EmployeeSrv services.EmployeeService
}

func NewEmployeeHandler(insSrv services.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{EmployeeSrv: insSrv}
}

func (h *EmployeeHandler) FindEmployeeByAccount() fiber.Handler {
	return func(c *fiber.Ctx) error {
		account := c.Query("account")
		if account == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "account parameter is required",
			})
		}

		employee, err := h.EmployeeSrv.FindEmployeeByAccount(account)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "failed to fetch employee",
				"detail": err.Error(),
			})
		}

		if employee == nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"error": "employee not found",
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data": employee,
		})
	}
}

func (h *EmployeeHandler) GetEmployeeByEmpCode() fiber.Handler {
	return func(c *fiber.Ctx) error {
		empCode := c.Query("empCode")
		fmt.Println("empCode : ", empCode)
		if empCode == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "empCode parameter is required",
			})
		}

		employee, err := h.EmployeeSrv.GetEmployeeByEmpCodeService(empCode)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error":  "failed to fetch employee",
				"detail": err.Error(),
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"data": employee,
		})
	}
}
