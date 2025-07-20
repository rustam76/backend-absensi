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
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)

	// Cek apakah sudah absen hari ini
	existing, err := s.repo.GetByEmployeeIDAndDate(attendance.EmployeeID, startOfDay, endOfDay)
	if err == nil && existing.AttendanceID != "" {
		return existing, fmt.Errorf("anda sudah melakukan absen hari ini")
	}

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
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	_, _ = s.repo.CreateHistory(history)

	return created, nil
}

func (s *attendanceService) ClockOut(attendanceID string, employeeID string) (model.Attendance, error) {
	attendance, err := s.repo.GetByAttendanceID(attendanceID)
	if err != nil {
		return attendance, fmt.Errorf("data absen tidak ditemukan, silakan clock in terlebih dahulu")
	}

	if attendance.ClockOut != nil {
		return attendance, fmt.Errorf("anda sudah melakukan clock out hari ini")
	}

	now := time.Now()
	attendance.ClockOut = &now
	attendance.UpdatedAt = now

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
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	_, _ = s.repo.CreateHistory(history)

	return updated, nil
}

func (s *attendanceService) GetAttendanceLogs(start_date string, end_date string, departementID string, employeeID string) ([]repository.AttendanceLogResponse, error) {
	return s.repo.GetAttendanceLogs(start_date, end_date, departementID, employeeID)
}
