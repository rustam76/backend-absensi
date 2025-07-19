package service

import (
	"backend-absensi/model"
	"backend-absensi/repository"
)

type DepartmentService interface {
	GetAll() ([]model.Departement, error)
	GetByID(id uint) (model.Departement, error)
	Create(departement model.Departement) (model.Departement, error)
	Update(id uint, departement model.Departement) (model.Departement, error)
	Delete(id uint) error
}

type departmentService struct {
	repo repository.DepartementRepository
}

func NewDepartmentService(repo repository.DepartementRepository) DepartmentService {
	return &departmentService{repo}
}

func (s *departmentService) GetAll() ([]model.Departement, error) {
	return s.repo.GetAll()
}

func (s *departmentService) GetByID(id uint) (model.Departement, error) {
	return s.repo.GetByID(id)
}

func (s *departmentService) Create(departement model.Departement) (model.Departement, error) {
	return s.repo.Create(departement)
}

func (s *departmentService) Update(id uint, departement model.Departement) (model.Departement, error) {
	return s.repo.Update(id, departement)
}

func (s *departmentService) Delete(id uint) error {
	return s.repo.Delete(id)
}
