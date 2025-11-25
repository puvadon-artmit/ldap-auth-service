package models

type EmployeeViewResp struct {
	UHR_EmpCode      string `gorm:"column:user_id"`
	UHR_FirstName_en string `gorm:"column:firstname"`
	UHR_LastName_en  string `gorm:"column:lastname"`
	UHR_Department   string `gorm:"column:role_name"`
	AD_UserLogon     string `gorm:"column:username"`
	AD_Mail          string `gorm:"column:email"`
	AD_AccountStatus string `gorm:"column:status"`
}
