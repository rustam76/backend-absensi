package dto

type EmployeeDTO struct {
	ID            uint   `json:"id"`
	EmployeeID    string `json:"employee_id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	DepartementID uint   `json:"departement_id"`
}

type EmployeeResponse struct {
	ID              uint   `json:"id"`
	EmployeeID      string `json:"employee_id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	DepartementID   uint   `json:"departement_id"`
	DepartementName string `json:"departement_name"`
}

type EmployeeRequest struct {
	EmployeeID    string `json:"employee_id" example:"EMP001"`
	Name          string `json:"name" example:"John Doe"`
	Address       string `json:"address" example:"Jl. Sudirman No. 1"`
	DepartementID uint   `json:"departement_id" example:"2"`
}
