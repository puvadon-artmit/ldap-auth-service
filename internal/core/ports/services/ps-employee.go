package ports

import "github.com/tomioka/ldap-auth-service/internal/core/models"

type EmployeeService interface {

	// ====================== Employee View ===================================

	FindEmployeeByAccount(account string) (*models.EmployeeViewResp, error)
	GetEmployeeByEmpCodeService(empCode string) (models.EmployeeViewResp, error)
}
