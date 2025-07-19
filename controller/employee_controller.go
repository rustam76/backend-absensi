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

type EmployeeController struct {
	service service.EmployeeService
}

// Gunakan pointer agar tidak copy struct tiap dipanggil
func NewEmployeeController(service service.EmployeeService) *EmployeeController {
	return &EmployeeController{service}
}

// GetAllEmployee godoc
// @Summary Get all employees
// @Description Get a list of all employees
// @Tags Employee
// @Accept json
// @Produce json
// @Success 200 {array} dto.EmployeeResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /employee [get]
func (c *EmployeeController) GetAllEmployee(ctx *fiber.Ctx) error {
	employees, err := c.service.GetAll()
	if err != nil {
		log.Println("Error getting all employees:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to retrieve employees",
		})
	}

	// Konversi ke DTO (jika belum dikonversi di service)
	var employeeDTOs []dto.EmployeeResponse
	for _, e := range employees {
		employeeDTOs = append(employeeDTOs, dto.EmployeeResponse{
			ID:              e.ID,
			EmployeeID:      e.EmployeeID,
			Name:            e.Name,
			Address:         e.Address,
			DepartementID:   e.DepartementID,
			DepartementName: e.DepartementName,
		})
	}

	return ctx.Status(http.StatusOK).JSON(employeeDTOs)
}

// GetEmployeeByID godoc
// @Summary Get an employee by ID
// @Description Get details of a specific employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} dto.EmployeeResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /employee/{id} [get]
func (c *EmployeeController) GetEmployeeByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	employee, err := c.service.GetByID(id)
	if err != nil {
		log.Println("Error getting employee by ID:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to retrieve employee",
		})
	}

	// Mapping model to DTO
	response := dto.EmployeeResponse{
		ID:              employee.ID,
		EmployeeID:      employee.EmployeeID,
		Name:            employee.Name,
		Address:         employee.Address,
		DepartementID:   employee.DepartementID,
		DepartementName: employee.Departement.DepartementName,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// CreateEmployee godoc
// @Summary Create new employee
// @Description Add a new employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param employee body dto.EmployeeRequest true "Employee details"
// @Success 201 {object} dto.EmployeeResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /employee [post]
func (c *EmployeeController) CreateEmployee(ctx *fiber.Ctx) error {
	var req dto.EmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	// Mapping DTO to model
	employee := model.Employee{
		EmployeeID:    req.EmployeeID,
		Name:          req.Name,
		Address:       req.Address,
		DepartementID: req.DepartementID,
	}

	createdEmployee, err := c.service.Create(employee)
	if err != nil {
		log.Println("Error creating employee:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to create employee",
		})
	}

	// Mapping model to DTO response
	response := dto.EmployeeResponse{
		ID:              createdEmployee.ID,
		EmployeeID:      createdEmployee.EmployeeID,
		Name:            createdEmployee.Name,
		Address:         createdEmployee.Address,
		DepartementID:   createdEmployee.DepartementID,
		DepartementName: createdEmployee.Departement.DepartementName,
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

// @Summary Update an employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Param employee body dto.EmployeeRequest true "Employee payload"
// @Success 200 {object} dto.EmployeeResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/employee/{id} [put]
func (c *EmployeeController) UpdateEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var req dto.EmployeeRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	// Mapping DTO ke Model
	employee := model.Employee{
		EmployeeID:    req.EmployeeID,
		Name:          req.Name,
		Address:       req.Address,
		DepartementID: req.DepartementID,
	}

	updatedEmployee, err := c.service.Update(id, employee)
	if err != nil {
		log.Println("Error updating employee:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to update employee",
		})
	}

	return ctx.Status(http.StatusOK).JSON(updatedEmployee)
}

// DeleteEmployee godoc
// @Summary Delete employee
// @Description Delete employee by ID
// @Tags Employee
// @Accept json
// @Produce json
// @Param id path string true "Employee ID"
// @Success 200 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /employee/{id} [delete]
func (c *EmployeeController) DeleteEmployee(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	err := c.service.Delete(id)
	if err != nil {
		log.Println("Error deleting employee:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to delete employee",
		})
	}

	return ctx.Status(http.StatusOK).JSON(utils.ErrorResponse{
		Message: "Employee deleted successfully",
	})
}
