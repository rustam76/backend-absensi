package controller

import (
	"backend-absensi/dto"
	"backend-absensi/service"
	"backend-absensi/utils"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthController struct {
	service service.EmployeeService
}

func NewAuthController(service service.EmployeeService) AuthController {
	return AuthController{service}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param employeeId path string true "Employee ID"
// @Success 200 {object} dto.LoginResponseDTO
// @Failure 400 {object} utils.ErrorResponse
// @Failure 401 {object} utils.ErrorResponse
// @Router /login [post]
func (c *AuthController) Login(ctx *fiber.Ctx) error {
	employeeId := ctx.Params("employeeId")
	employee, err := c.service.GetEmployeeByID(employeeId)
	if err != nil {
		log.Println("Error getting employee by ID:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to retrieve employee",
		})
	}

	token, err := utils.GenerateJWT(uint(employee.ID), employee.Departement.DepartementName)
	if err != nil {
		log.Println("Error generating token:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to generate token",
		})
	}

	response := dto.LoginResponseDTO{
		Token: token,
		User: dto.LoginUserInfoResponseDTO{
			EmployeeID:      employee.EmployeeID,
			Name:            employee.Name,
			Address:         employee.Address,
			Departement:     employee.Departement.DepartementName,
			MaxClockInTime:  employee.Departement.MaxClockInTime,
			MaxClockOutTime: employee.Departement.MaxClockOutTime,
		},
	}

	return ctx.Status(http.StatusOK).JSON(response)

}
