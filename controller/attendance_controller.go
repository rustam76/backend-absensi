package controller

import (
	"backend-absensi/dto"
	"backend-absensi/model"
	"backend-absensi/service"
	"backend-absensi/utils"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AttendanceController struct {
	service service.AttendanceService
}

func NewAttendanceController(service service.AttendanceService) AttendanceController {
	return AttendanceController{service}
}

// ClockIn godoc
// @Summary Clock in attendance
// @Description Employee clock in
// @Tags Attendance
// @Accept json
// @Produce json
// @Param data body model.Attendance true "Attendance data"
// @Success 201 {object} model.Attendance
// @Failure 400 {object} utils.ErrorResponse
// @Router /api/attendance/clock-in [post]
func (c *AttendanceController) ClockIn(ctx *fiber.Ctx) error {
	var req model.Attendance
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{Message: "Invalid body"})
	}

	created, err := c.service.ClockIn(req)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{Message: "Failed to clock in"})
	}

	return ctx.Status(http.StatusCreated).JSON(created)
}

// ClockOut godoc
// @Summary Clock out attendance
// @Description Employee clock out
// @Tags Attendance
// @Accept json
// @Produce json
// @Param id path string true "Attendance ID"
// @Param data body map[string]string true "Employee ID"
// @Success 200 {object} model.Attendance
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/attendance/clock-out/{id} [put]
func (c *AttendanceController) ClockOut(ctx *fiber.Ctx) error {
	attendanceID := ctx.Params("id")
	var body struct {
		EmployeeID string `json:"employee_id"`
	}
	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{Message: "Invalid body"})
	}

	updated, err := c.service.ClockOut(attendanceID, body.EmployeeID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{Message: "Failed to clock out"})
	}

	return ctx.Status(http.StatusOK).JSON(updated)
}

// GetAttendanceLog godoc
// @Summary Get attendance logs
// @Description Get attendance logs with filter
// @Tags Attendance
// @Accept json
// @Produce json
// @Param start_date query string false "Start Date"
// @Param end_date query string false "End Date"
// @Param departement_id query string false "Departement ID"
// @Param employee_id query string false "Employee ID"
// @Success 200 {array} dto.AttendanceLogResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/attendance/logs [get]
func (c *AttendanceController) GetAttendanceLog(ctx *fiber.Ctx) error {
	start_date := ctx.Query("start_date")
	end_date := ctx.Query("end_date")
	departementID := ctx.Query("departement_id")
	employeeID := ctx.Query("employee_id")

	logs, err := c.service.GetAttendanceLogs(start_date, end_date, departementID, employeeID)
	if err != nil {
		log.Println("Error fetching attendance logs:", err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to fetch attendance logs",
		})
	}

	// Jika logs masih model, map ke DTO dulu
	var logsDTO []dto.AttendanceLogResponse
	for _, logEntry := range logs {
		logsDTO = append(logsDTO, dto.AttendanceLogResponse{
			EmployeeName:    logEntry.EmployeeName,
			DepartementName: logEntry.DepartementName,
			DateAttendance:  logEntry.DateAttendance,
			ClockIn:         logEntry.ClockIn,
			ClockOut:        logEntry.ClockOut,
			AttendanceID:    logEntry.AttendanceID,
			MaxClockInTime:  logEntry.MaxClockInTime,
			MaxClockOutTime: logEntry.MaxClockOutTime,
			StatusClockIn:   logEntry.StatusClockIn,
			StatusClockOut:  logEntry.StatusClockOut,
			IsLate:          logEntry.IsLate,
			IsLeaveEarly:    logEntry.IsLeaveEarly,
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(logsDTO)
}
