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
	query := `
        SELECT *
        FROM employee_view
        WHERE UHR_FullNameEn = ?
        LIMIT 1;
    `
	if err := r.db.Raw(query, fullNameEn).Scan(&emp).Error; err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepositoryDB) GetEmployeeByEmpCode(empCode string) (*domains.EmployeeView, error) {
	var emp domains.EmployeeView
	query := `
        SELECT *
        FROM employee_view
        WHERE UHR_EmpCode = ?
        LIMIT 1;
    `
	if err := r.db.Debug().Raw(query, empCode).Scan(&emp).Error; err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepositoryDB) FindEmployeeByAccount(account string) (*domains.EmployeeView, error) {
	fmt.Println("account:", account)
	var user domains.EmployeeView

	query := `
        SELECT *
        FROM employee_view
        WHERE AD_UserLogon = ?
        LIMIT 1;
    `

	tx := r.db.Raw(query, account).Scan(&user)

	if tx.Error != nil {
		return nil, tx.Error
	}

	if tx.RowsAffected == 0 {
		return nil, nil
	}

	return &user, nil
}
