package model

import "time"

type Attendance struct {
	ID           uint64     `gorm:"primaryKey;autoIncrement" json:"id"`
	EmployeeID   string     `gorm:"type:varchar(50);not null" json:"employee_id"`
	AttendanceID string     `gorm:"type:varchar(100);uniqueIndex;not null" json:"attendance_id"`
	ClockIn      time.Time  `gorm:"type:timestamp" json:"clock_in"`
	ClockOut     *time.Time `gorm:"type:timestamp" json:"clock_out"`
	CreatedAt    time.Time  `gorm:"type:datetime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"type:datetime" json:"updated_at"`

	Employee Employee `gorm:"-"`
}

func (Attendance) TableName() string {
	return "attendance"
}
