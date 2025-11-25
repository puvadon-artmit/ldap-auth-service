package services

import (
	"github.com/tomioka/ldap-auth-service/internal/core/models"
	ports "github.com/tomioka/ldap-auth-service/internal/core/ports/repositories"
	servicesports "github.com/tomioka/ldap-auth-service/internal/core/ports/services"
)

type EmployeeService struct {
	EmployeeRepo ports.EmployeeRepository
}

func NewEmployeeService(EmployeeRepo ports.EmployeeRepository) servicesports.EmployeeService {
	return &EmployeeService{EmployeeRepo: EmployeeRepo}
}

func (s *EmployeeService) FindEmployeeByAccount(account string) (*models.EmployeeViewResp, error) {
	employee, err := s.EmployeeRepo.FindEmployeeByAccount(account)
	if err != nil {
		return nil, err
	}
	if employee == nil {
		return nil, nil
	}

	resp := &models.EmployeeViewResp{
		UHR_EmpCode:      employee.UHR_EmpCode,
		UHR_FirstName_en: employee.UHR_FirstName_en,
		UHR_LastName_en:  employee.UHR_LastName_en,
		UHR_Department:   employee.UHR_Department,
		AD_UserLogon:     employee.AD_UserLogon,
		AD_Mail:          employee.AD_Mail,
		AD_AccountStatus: employee.AD_AccountStatus,
	}

	return resp, nil
}

func (s *EmployeeService) GetEmployeeByEmpCodeService(empCode string) (models.EmployeeViewResp, error) {
	user, err := s.EmployeeRepo.GetEmployeeByEmpCode(empCode)
	if err != nil {
		return models.EmployeeViewResp{}, err
	}

	userReq := models.EmployeeViewResp{
		UHR_EmpCode:      user.UHR_EmpCode,
		UHR_FirstName_en: user.UHR_FirstName_en,
		UHR_LastName_en:  user.UHR_LastName_en,
		UHR_Department:   user.UHR_Department,
		AD_UserLogon:     user.AD_UserLogon,
		AD_Mail:          user.AD_Mail,
		AD_AccountStatus: user.AD_AccountStatus,
	}

	return userReq, nil
}
