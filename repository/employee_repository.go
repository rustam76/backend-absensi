package repository

import (
	"backend-absensi/model"

	"gorm.io/gorm"
)

type EmployeeRepository interface {
	GetAll() ([]EmployeeResponse, error)
	GetByID(id string) (model.Employee, error)
	GetEmployeeByID(id string) (model.Employee, error)
	Create(employee model.Employee) (model.Employee, error)
	Update(id string, employee model.Employee) (model.Employee, error)
	Delete(id string) error
}

type EmployeeResponse struct {
	ID              uint   `json:"id"`
	EmployeeID      string `json:"employee_id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	DepartementID   uint   `json:"departement_id"`
	DepartementName string `json:"departement_name"`
}

type employeeRepo struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) EmployeeRepository {
	return &employeeRepo{db}
}

func (r *employeeRepo) GetAll() ([]EmployeeResponse, error) {
	var employees []model.Employee
	err := r.db.Preload("Departement").Find(&employees).Error

	if err != nil {
		return nil, err
	}

	var res []EmployeeResponse

	for _, employee := range employees {
		res = append(res, EmployeeResponse{
			ID:              employee.ID,
			EmployeeID:      employee.EmployeeID,
			Name:            employee.Name,
			Address:         employee.Address,
			DepartementID:   employee.Departement.ID,
			DepartementName: employee.Departement.DepartementName,
		})
	}

	return res, nil
}

func (r *employeeRepo) GetByID(id string) (model.Employee, error) {
	var employee model.Employee
	err := r.db.Preload("Departement").First(&employee, "employee_id = ?", id).Error
	return employee, err
}

func (r *employeeRepo) GetEmployeeByID(id string) (model.Employee, error) {
	var employee model.Employee
	err := r.db.Preload("Departement").First(&employee, "employee_id = ?", id).Error
	return employee, err
}

func (r *employeeRepo) Create(employee model.Employee) (model.Employee, error) {
	err := r.db.Create(&employee).Error
	return employee, err
}

func (r *employeeRepo) Update(id string, employee model.Employee) (model.Employee, error) {
	err := r.db.Model(&model.Employee{}).Where("employee_id = ?", id).Updates(employee).Error
	return employee, err
}

func (r *employeeRepo) Delete(id string) error {
	err := r.db.Delete(&model.Employee{}, "employee_id = ?", id).Error
	return err
}
