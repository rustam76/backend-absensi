package dto

type LoginResponseDTO struct {
	Token string                   `json:"token" example:"eyJhbGciOiJIUzI1NiIs..."`
	User  LoginUserInfoResponseDTO `json:"user"`
}

type LoginUserInfoResponseDTO struct {
	EmployeeID      string `json:"employee_id"`
	Name            string `json:"name"`
	Address         string `json:"address"`
	Departement     string `json:"departement"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
}
