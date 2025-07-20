package repository

import (
	"backend-absensi/model"
	"strings"
	"time"

	"gorm.io/gorm"
)

type AttendanceRepository interface {
	Create(attendance model.Attendance) (model.Attendance, error)
	Update(id string, attendance model.Attendance) (model.Attendance, error)
	CreateHistory(history model.AttendanceHistory) (model.AttendanceHistory, error)
	GetByEmployeeIDAndDate(employeeID string, start, end time.Time) (model.Attendance, error)
	GetByAttendanceID(attendanceID string) (model.Attendance, error)
	GetAttendanceLogs(start_date string, end_date string, departementID string, employeeID string) ([]AttendanceLogResponse, error)
}

type attendanceRepo struct {
	db *gorm.DB
}

type AttendanceLogResponse struct {
	EmployeeName    string `json:"employee_name"`
	DepartementName string `json:"departement_name"`
	DateAttendance  string `json:"date_attendance"`
	ClockIn         string `json:"clock_in"`
	ClockOut        string `json:"clock_out"`
	AttendanceID    string `json:"attendance_id"`
	MaxClockInTime  string `json:"max_clock_in_time"`
	MaxClockOutTime string `json:"max_clock_out_time"`
	StatusClockIn   string `json:"status_clock_in"`
	StatusClockOut  string `json:"status_clock_out"`
	IsLate          string `json:"is_late"`        // "true" or "false"
	IsLeaveEarly    string `json:"is_leave_early"` // "true" or "false"
}

func NewAttendanceRepository(db *gorm.DB) AttendanceRepository {
	return &attendanceRepo{db}
}

func (r *attendanceRepo) CreateHistory(history model.AttendanceHistory) (model.AttendanceHistory, error) {
	err := r.db.Create(&history).Error
	return history, err
}

func (r *attendanceRepo) GetByAttendanceID(attendanceID string) (model.Attendance, error) {
	var attendance model.Attendance
	err := r.db.Where("attendance_id = ?", attendanceID).First(&attendance).Error
	return attendance, err
}

func (r *attendanceRepo) Create(attendance model.Attendance) (model.Attendance, error) {
	err := r.db.Create(&attendance).Error
	return attendance, err
}

func (r *attendanceRepo) Update(id string, attendance model.Attendance) (model.Attendance, error) {
	err := r.db.Save(&attendance).Error
	return attendance, err
}

func (r *attendanceRepo) GetByEmployeeIDAndDate(employeeID string, start, end time.Time) (model.Attendance, error) {
	var attendance model.Attendance
	err := r.db.Where("employee_id = ? AND clock_in >= ? AND clock_in < ?", employeeID, start, end).
		First(&attendance).Error
	return attendance, err
}

func (r *attendanceRepo) GetAttendanceLogs(start_date string, end_date string, departementID string, employeeID string) ([]AttendanceLogResponse, error) {
	var logs []AttendanceLogResponse

	// Subquery untuk ambil status check-in dan check-out
	subQuery := r.db.Table("attendance_history AS ah").
		Select(`
			ah.attendance_id,
			MAX(CASE WHEN ah.attendance_type = 1 THEN ah.date_attendance END) AS date_attendance,
			MAX(CASE WHEN ah.attendance_type = 1 THEN ah.attendance_type END) AS status_clock_in,
			MAX(CASE WHEN ah.attendance_type = 2 THEN ah.attendance_type END) AS status_clock_out
		`).
		Group("ah.attendance_id")

	// Query utama
	query := r.db.Table("attendance AS a").
		Select(`
			e.name AS employee_name, 
			d.departement_name,
			ah.date_attendance,
			a.clock_in, 
			a.clock_out, 
			a.attendance_id,
			d.max_clock_in_time, 
			d.max_clock_out_time,
			ah.status_clock_in,
			ah.status_clock_out
		`).
		Joins("LEFT JOIN (?) AS ah ON a.attendance_id = ah.attendance_id", subQuery).
		Joins("LEFT JOIN employee AS e ON a.employee_id = e.employee_id").
		Joins("LEFT JOIN departements AS d ON e.departement_id = d.id").
		Order("ah.date_attendance DESC")

	// Filter opsional
	if start_date != "" && end_date != "" {
		query = query.Where("ah.date_attendance BETWEEN ? AND ?", start_date, end_date)
	}
	if departementID != "" {
		query = query.Where("e.departement_id = ?", departementID)
	}
	if employeeID != "" {
		query = query.Where("a.employee_id = ?", employeeID)
	}

	// Eksekusi query
	err := query.Scan(&logs).Error
	if err != nil {
		return nil, err
	}

	// Tambahkan pengecekan is_late dan is_leave_early di Go
	layout := "15:04:05"

	for i := range logs {
		// Check is_late
		if logs[i].ClockIn != "" && logs[i].MaxClockInTime != "" {
			clockInParts := strings.Split(logs[i].ClockIn, "T")
			if len(clockInParts) == 2 {
				clockInStr := clockInParts[1][:8] // HH:mm:ss
				clockIn, err1 := time.Parse(layout, clockInStr)
				maxIn, err2 := time.Parse(layout, logs[i].MaxClockInTime)
				if err1 == nil && err2 == nil && clockIn.After(maxIn) {
					logs[i].IsLate = "true"
				} else {
					logs[i].IsLate = "false"
				}
			}
		} else {
			logs[i].IsLate = "false"
		}

		// Check is_leave_early
		if logs[i].ClockOut != "" && logs[i].MaxClockOutTime != "" {
			clockOutParts := strings.Split(logs[i].ClockOut, "T")
			if len(clockOutParts) == 2 {
				clockOutStr := clockOutParts[1][:8] // HH:mm:ss
				clockOut, err1 := time.Parse(layout, clockOutStr)
				maxOut, err2 := time.Parse(layout, logs[i].MaxClockOutTime)
				if err1 == nil && err2 == nil && clockOut.Before(maxOut) {
					logs[i].IsLeaveEarly = "true"
				} else {
					logs[i].IsLeaveEarly = "false"
				}
			}
		} else {
			logs[i].IsLeaveEarly = "false"
		}
	}

	return logs, nil
}
