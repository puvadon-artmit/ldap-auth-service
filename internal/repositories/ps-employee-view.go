package repositories

import (
	"fmt"

	"github.com/tomioka/ldap-auth-service/internal/core/domains"
	ports "github.com/tomioka/ldap-auth-service/internal/core/ports/repositories"

	"gorm.io/gorm"
)

type EmployeeRepositoryDB struct {
	db *gorm.DB
}

func NewEmployeeRepositoryDB(db *gorm.DB) ports.EmployeeRepository {
	return &EmployeeRepositoryDB{db: db}
}

func (r *EmployeeRepositoryDB) GetEmployeeByFullNameEn(fullNameEn string) (*domains.EmployeeView, error) {
	var emp domains.EmployeeView
	if err := r.db.
		Where("UHR_FullNameEn = ?", fullNameEn).
		First(&emp).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepositoryDB) GetEmployeeByEmpCode(empCode string) (*domains.EmployeeView, error) {
	fmt.Println("empCode : ", empCode)
	var emp domains.EmployeeView
	if err := r.db.Debug().
		Where("UHR_EmpCode = ?", empCode).
		First(&emp).Error; err != nil {
		return nil, err
	}
	return &emp, nil
}

func (r *EmployeeRepositoryDB) FindEmployeeByAccount(account string) (*domains.EmployeeView, error) {
	fmt.Println("account : ", account)
	var user domains.EmployeeView

	if err := r.db.Where("AD_UserLogon = ?", account).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
