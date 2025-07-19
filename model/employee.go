package model

import "time"

// Employee represents the employee table
type Employee struct {
	ID            uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID    string     `gorm:"type:varchar(50);uniqueIndex;not null" json:"employee_id"`
	DepartementID uint       `gorm:"type:int;not null" json:"departement_id"`
	Name          string     `gorm:"type:varchar(255);not null" json:"name"`
	Address       string     `gorm:"type:text" json:"address"`
	CreatedAt     *time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"type:datetime" json:"updated_at"`

	Departement Departement `gorm:"foreignKey:DepartementID" json:"departements"`
}

func (Employee) TableName() string {
	return "employee"
}
