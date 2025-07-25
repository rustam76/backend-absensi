package dto

import "time"

type AttendanceLogResponse struct {
	EmployeeName    string `json:"employee_name"`
	DepartementName string `json:"departement_name"`
	DateAttendance  string `json:"date_attendance"`
	ClockIn         string `json:"clock_in"`
	ClockOut        string `json:"clock_out,omitempty"`
	AttendanceID    string `json:"attendance_id"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
	StatusClockIn   string `json:"status_clock_in"`
	StatusClockOut  string `json:"status_clock_out"`
	IsLate          string `json:"is_late"`
	IsLeaveEarly    string `json:"is_leave_early"`
}

type AttendanceClockInOutRequest struct {
	EmployeeID string `json:"employee_id"`
}

type AttendanceClockInOutResponse struct {
	EmployeeID   string     `json:"employee_id"`
	AttendanceID string     `json:"attendance_id"`
	ClockIn      time.Time  `json:"clock_in"`
	ClockOut     *time.Time `json:"clock_out,omitempty"`
}
