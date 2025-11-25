package ports

import "github.com/tomioka/ldap-auth-service/internal/core/domains"

type EmployeeRepository interface {

	// ====================== Employee View ===================================
	FindEmployeeByAccount(account string) (*domains.EmployeeView, error)
	GetEmployeeByEmpCode(empCode string) (*domains.EmployeeView, error)
}
