package controller

import (
	"backend-absensi/dto"
	"backend-absensi/model"
	"backend-absensi/service"
	"backend-absensi/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type DepartmentController struct {
	service service.DepartmentService
}

func NewDepartmentController(service service.DepartmentService) DepartmentController {
	return DepartmentController{service}
}

// GetAllDepartement godoc
// @Summary Get all departments
// @Description Get a list of all departments
// @Tags Department
// @Accept json
// @Produce json
// @Success 200 {array} dto.DepartmentResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/departements [get]
func (c *DepartmentController) GetAllDepartement(ctx *fiber.Ctx) error {
	departements, err := c.service.GetAll()
	if err != nil {
		log.Println("Error getting all departements:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to retrieve departements",
		})
	}

	// Mapping model ke DTO
	var response []dto.DepartmentResponse
	for _, d := range departements {
		response = append(response, dto.DepartmentResponse{
			ID:              d.ID,
			DepartementName: d.DepartementName,
			MaxClockInTime:  d.MaxClockInTime,
			MaxClockOutTime: d.MaxClockOutTime,
		})
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// GetDepartementByID godoc
// @Summary Get department by ID
// @Description Get a specific department by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} dto.DepartmentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/departement/{id} [get]
func (c *DepartmentController) GetDepartementByID(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("Error parsing ID:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid ID",
		})
	}
	id := uint(idInt)

	departement, err := c.service.GetByID(id)
	if err != nil {
		log.Println("Error getting departement by ID:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to retrieve departement",
		})
	}

	response := dto.DepartmentResponse{
		ID:              departement.ID,
		DepartementName: departement.DepartementName,
		MaxClockInTime:  departement.MaxClockInTime,
		MaxClockOutTime: departement.MaxClockOutTime,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// CreateDepartement godoc
// @Summary Create a new department
// @Description Create a new department based on request body
// @Tags Department
// @Accept json
// @Produce json
// @Param departement body dto.DepartmentRequest true "Department to create"
// @Success 201 {object} dto.DepartmentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/departements [post]
func (c *DepartmentController) CreateDepartement(ctx *fiber.Ctx) error {
	var req dto.DepartmentRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	departement := model.Departement{
		DepartementName: req.DepartementName,
		MaxClockInTime:  req.MaxClockInTime,
		MaxClockOutTime: req.MaxClockOutTime,
	}

	createdDepartement, err := c.service.Create(departement)
	if err != nil {
		log.Println("Error creating departement:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to create departement",
		})
	}

	response := dto.DepartmentResponse{
		ID:              createdDepartement.ID,
		DepartementName: createdDepartement.DepartementName,
		MaxClockInTime:  createdDepartement.MaxClockInTime,
		MaxClockOutTime: createdDepartement.MaxClockOutTime,
	}

	return ctx.Status(http.StatusCreated).JSON(response)
}

// UpdateDepartement godoc
// @Summary Update a department
// @Description Update an existing department
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Param body body dto.DepartmentRequest true "Updated department"
// @Success 200 {object} dto.DepartmentResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/departement/{id} [put]
func (c *DepartmentController) UpdateDepartement(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("Error parsing ID:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid ID",
		})
	}
	id := uint(idInt)

	var req dto.DepartmentRequest
	if err := ctx.BodyParser(&req); err != nil {
		log.Println("Error parsing request body:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid request body",
		})
	}

	departement := model.Departement{
		DepartementName: req.DepartementName,
		MaxClockInTime:  req.MaxClockInTime,
		MaxClockOutTime: req.MaxClockOutTime,
	}

	updatedDepartement, err := c.service.Update(id, departement)
	if err != nil {
		log.Println("Error updating departement:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to update departement",
		})
	}

	response := dto.DepartmentResponse{
		ID:              updatedDepartement.ID,
		DepartementName: updatedDepartement.DepartementName,
		MaxClockInTime:  updatedDepartement.MaxClockInTime,
		MaxClockOutTime: updatedDepartement.MaxClockOutTime,
	}

	return ctx.Status(http.StatusOK).JSON(response)
}

// DeleteDepartement godoc
// @Summary Delete a department
// @Description Delete a department by ID
// @Tags Department
// @Accept json
// @Produce json
// @Param id path int true "Department ID"
// @Success 200 {object} utils.ErrorResponse
// @Failure 400 {object} utils.ErrorResponse
// @Failure 500 {object} utils.ErrorResponse
// @Router /api/departement/{id} [delete]
func (c *DepartmentController) DeleteDepartement(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	idInt, err := strconv.Atoi(idParam)
	if err != nil {
		log.Println("Error parsing ID:", err)
		return ctx.Status(http.StatusBadRequest).JSON(utils.ErrorResponse{
			Message: "Invalid ID",
		})
	}
	id := uint(idInt)

	err = c.service.Delete(id)
	if err != nil {
		log.Println("Error deleting departement:", err)
		return ctx.Status(http.StatusInternalServerError).JSON(utils.ErrorResponse{
			Message: "Failed to delete departement",
		})
	}

	return ctx.Status(http.StatusOK).JSON(utils.ErrorResponse{
		Message: "Departement deleted successfully",
	})
}
