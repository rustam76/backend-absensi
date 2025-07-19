package service

import (
	"backend-absensi/model"
	"backend-absensi/repository"
)

type EmployeeService interface {
	GetAll() ([]repository.EmployeeResponse, error)
	GetByID(id string) (model.Employee, error)
	GetEmployeeByID(id string) (model.Employee, error)
	Create(employee model.Employee) (model.Employee, error)
	Update(id string, employee model.Employee) (model.Employee, error)
	Delete(id string) error
}

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeService{repo}
}

func (s *employeeService) GetAll() ([]repository.EmployeeResponse, error) {
	return s.repo.GetAll()
}

func (s *employeeService) GetByID(id string) (model.Employee, error) {
	return s.repo.GetByID(id)
}

func (s *employeeService) GetEmployeeByID(id string) (model.Employee, error) {
	return s.repo.GetEmployeeByID(id)
}

func (s *employeeService) Create(employee model.Employee) (model.Employee, error) {
	return s.repo.Create(employee)
}

func (s *employeeService) Update(id string, employee model.Employee) (model.Employee, error) {
	return s.repo.Update(id, employee)
}

func (s *employeeService) Delete(id string) error {
	return s.repo.Delete(id)
}
