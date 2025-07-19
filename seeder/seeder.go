package seeder

import (
	"backend-absensi/model"
	"time"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	// Seeder Departement
	departementMap := map[string]model.Departement{
		"HRD":     {DepartementName: "HRD", MaxClockInTime: "09:00:00", MaxClockOutTime: "17:00:00"},
		"IT":      {DepartementName: "IT", MaxClockInTime: "10:00:00", MaxClockOutTime: "18:00:00"},
		"Finance": {DepartementName: "Finance", MaxClockInTime: "08:30:00", MaxClockOutTime: "16:30:00"},
	}

	for name, dept := range departementMap {
		var existing model.Departement
		err := db.Where("departement_name = ?", dept.DepartementName).First(&existing).Error
		if err != nil && err == gorm.ErrRecordNotFound {
			if err := db.Create(&dept).Error; err == nil {
				departementMap[name] = dept
			}
		} else {
			departementMap[name] = existing
		}
	}

	// Seeder Employee
	now := time.Now()
	employees := []model.Employee{
		{
			EmployeeID:    "EMP0011",
			Name:          "Budi Santoso",
			Address:       "Jl. Merdeka No. 1",
			DepartementID: departementMap["HRD"].ID,
			CreatedAt:     &now,
			UpdatedAt:     &now,
		},
		{
			EmployeeID:    "EMP0022",
			Name:          "Sari Dewi",
			Address:       "Jl. Sudirman No. 5",
			DepartementID: departementMap["IT"].ID,
			CreatedAt:     &now,
			UpdatedAt:     &now,
		},
		{
			EmployeeID:    "EMP0033",
			Name:          "Joko Pranoto",
			Address:       "Jl. Diponegoro No. 9",
			DepartementID: departementMap["Finance"].ID,
			CreatedAt:     &now,
			UpdatedAt:     &now,
		},
	}

	for _, e := range employees {
		var existing model.Employee
		if err := db.Where("employee_id = ?", e.EmployeeID).First(&existing).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				db.Create(&e)
			}
		}
	}

	return nil
}
