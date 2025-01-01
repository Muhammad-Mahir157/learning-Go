package api

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/database"
	"github.com/Muhammad-Mahir157/clockify-app-clone/infrastructure/postgresDb"
	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	timeLogRepo := postgresDb.NewTimeLogRepository(database.Db)

	timeLogService := usecase.NewTimeLogService(timeLogRepo)

	timeLogController := NewTimeLogController(timeLogService)

	routes := app.Group("/api/")
	routes.Post("/logTime", timeLogController.LogNewTime)
	routes.Put("/updateLoggedTime", timeLogController.UpdateLoggedTime)
	routes.Delete("/deleteLoggedTime/:Id", timeLogController.DeleteLoggedTime)
	routes.Get("/getLoggedTime", timeLogController.GetAllTimeLogs)
	routes.Get("/getLoggedTimeById/:Id", timeLogController.GetLoggedTimeById)
}
