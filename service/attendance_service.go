package service

import (
	"backend-absensi/model"
	"backend-absensi/repository"
	"fmt"
	"time"
)

type AttendanceService interface {
	ClockIn(attendance model.Attendance) (model.Attendance, error)
	ClockOut(attendanceID string, employeeID string) (model.Attendance, error)
	GetAttendanceLogs(start_date string, end_date string, departementID string, employeeID string) ([]repository.AttendanceLogResponse, error)
}

type attendanceService struct {
	repo repository.AttendanceRepository
}

func NewAttendanceService(repo repository.AttendanceRepository) AttendanceService {
	return &attendanceService{repo}
}

func (s *attendanceService) ClockIn(attendance model.Attendance) (model.Attendance, error) {
	now := time.Now()
	attendance.ClockIn = now
	attendance.AttendanceID = fmt.Sprintf("%s-%d", attendance.EmployeeID, now.Unix())
	attendance.CreatedAt = now
	attendance.UpdatedAt = now
	attendance.ClockOut = nil

	created, err := s.repo.Create(attendance)
	if err != nil {
		return created, err
	}

	history := model.AttendanceHistory{
		EmployeeID:     attendance.EmployeeID,
		AttendanceID:   attendance.AttendanceID,
		DateAttendance: attendance.ClockIn,
		AttendanceType: 1,
		Description:    "Clock In",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	_, _ = s.repo.CreateHistory(history)

	return created, nil
}
func (s *attendanceService) ClockOut(attendanceID string, employeeID string) (model.Attendance, error) {
	attendance, err := s.repo.GetByAttendanceID(attendanceID)
	if err != nil {
		return attendance, err
	}

	now := time.Now()
	attendance.ClockOut = &now

	attendance.UpdatedAt = time.Now()

	updated, err := s.repo.Update(attendanceID, attendance)
	if err != nil {
		return updated, err
	}

	history := model.AttendanceHistory{
		EmployeeID:     employeeID,
		AttendanceID:   attendanceID,
		DateAttendance: now,
		AttendanceType: 2,
		Description:    "Clock Out",
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}
	_, _ = s.repo.CreateHistory(history)

	return updated, nil
}

func (s *attendanceService) GetAttendanceLogs(start_date string, end_date string, departementID string, employeeID string) ([]repository.AttendanceLogResponse, error) {
	return s.repo.GetAttendanceLogs(start_date, end_date, departementID, employeeID)
}
