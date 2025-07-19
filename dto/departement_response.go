package dto

type DepartmentRequest struct {
	DepartementName string `json:"departement_name"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
}

type DepartmentResponse struct {
	ID              uint   `json:"id"`
	DepartementName string `json:"departement_name"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
}
