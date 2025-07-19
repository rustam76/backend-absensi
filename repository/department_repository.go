package repository

import (
	"backend-absensi/model"

	"gorm.io/gorm"
)

type DepartementRepository interface {
	GetAll() ([]model.Departement, error)
	GetByID(id uint) (model.Departement, error)
	Create(departement model.Departement) (model.Departement, error)
	Update(id uint, departement model.Departement) (model.Departement, error)
	Delete(id uint) error
}

type departementRepo struct {
	db *gorm.DB
}

func NewDepartementRepository(db *gorm.DB) DepartementRepository {
	return &departementRepo{db}
}

func (r *departementRepo) GetAll() ([]model.Departement, error) {
	var departements []model.Departement
	err := r.db.Find(&departements).Error
	return departements, err
}

func (r *departementRepo) GetByID(id uint) (model.Departement, error) {
	var departement model.Departement
	err := r.db.First(&departement, id).Error
	return departement, err
}

func (r *departementRepo) Create(departement model.Departement) (model.Departement, error) {
	err := r.db.Create(&departement).Error
	return departement, err
}

func (r *departementRepo) Update(id uint, departement model.Departement) (model.Departement, error) {
	//update departement by id
	err := r.db.Model(&model.Departement{}).Where("id = ?", id).Updates(departement).Error
	if err != nil {
		return model.Departement{}, err
	}
	return r.GetByID(id)
}

func (r *departementRepo) Delete(id uint) error {
	err := r.db.Delete(&model.Departement{}, id).Error
	return err
}
