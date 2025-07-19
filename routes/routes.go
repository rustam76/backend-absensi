package routes

import (
	"backend-absensi/config"
	"backend-absensi/controller"
	"backend-absensi/middleware"
	"backend-absensi/repository"
	"backend-absensi/service"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	config.ConnectDB()

	employeeController := controller.NewEmployeeController(service.NewEmployeeService(repository.NewEmployeeRepository(config.DB)))

	api := app.Group("/api")
	// API routes employee (protected)
	employee := api.Group("/employee", middleware.JWTMiddleware)
	employee.Get("/", employeeController.GetAllEmployee)
	employee.Get("/:id", employeeController.GetEmployeeByID)
	employee.Post("/", employeeController.CreateEmployee)
	employee.Put("/:id", employeeController.UpdateEmployee)
	employee.Delete("/:id", employeeController.DeleteEmployee)

	// API routes departement
	departementController := controller.NewDepartmentController(service.NewDepartmentService(repository.NewDepartementRepository(config.DB)))

	departement := api.Group("/departement", middleware.JWTMiddleware)
	departement.Get("/", departementController.GetAllDepartement)
	departement.Get("/:id", departementController.GetDepartementByID)
	departement.Post("/", departementController.CreateDepartement)
	departement.Put("/:id", departementController.UpdateDepartement)
	departement.Delete("/:id", departementController.DeleteDepartement)

	// Api Attendance
	attendanceController := controller.NewAttendanceController(service.NewAttendanceService(repository.NewAttendanceRepository(config.DB)))

	attendance := api.Group("/attendance", middleware.JWTMiddleware)
	attendance.Post("/clock-in", attendanceController.ClockIn)
	attendance.Put("/clock-out/:id", attendanceController.ClockOut)
	// List Log Absensi Karyawan
	attendance.Get("/logs", attendanceController.GetAttendanceLog)

	// Api Login
	authController := controller.NewAuthController(service.NewEmployeeService(repository.NewEmployeeRepository(config.DB)))
	api.Post("/login/:employeeId", authController.Login)

}
