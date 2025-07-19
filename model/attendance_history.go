package model

import "time"

type AttendanceHistory struct {
	ID             uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID     string    `gorm:"type:varchar(50);not null" json:"employee_id"`
	AttendanceID   string    `gorm:"type:varchar(100);not null" json:"attendance_id"`
	DateAttendance time.Time `gorm:"type:timestamp" json:"date_attendance"`
	AttendanceType uint8     `gorm:"type:tinyint" json:"attendance_type"`
	Description    string    `gorm:"type:text" json:"description"`
	CreatedAt      time.Time `gorm:"type:datetime" json:"created_at"`
	UpdatedAt      time.Time `gorm:"type:datetime" json:"updated_at"`

	Employee   Employee   `gorm:"-"`
	Attendance Attendance `gorm:"-"`
}

func (AttendanceHistory) TableName() string {
	return "attendance_history"
}
