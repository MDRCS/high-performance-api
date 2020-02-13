package Models

type Employee struct {
	ID              int         `gorm:"column:employee_id" json:"employee_id"`
	Firstname       string      `gorm:"column:first_name" json:"first_name"`
	Lastname        string      `gorm:"column:last_name" json:"last_name"`
	Email			string      `gorm:"column:email" json:"email"`
	Gender  		string      `gorm:"column:gender" json:"gender"`
	Employeer       Employeer   `gorm:"column:employeer_id" json:"employeer_id"`
}

type Employeer struct {
	ID 		int    `gorm:"column:employeer_id" json:"employeer_id"`
	Company string `gorm:"column:company_name" json:"company_name"`
}
