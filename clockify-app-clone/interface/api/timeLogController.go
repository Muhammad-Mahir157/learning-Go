package api

import (
	"net/http"

	"github.com/Muhammad-Mahir157/clockify-app-clone/application/interfaces"

	responseMapper "github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/mapper"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/request"
	"github.com/gofiber/fiber/v2"
)

type TimeLogController struct {
	service interfaces.TimeLogService
}

func NewTimeLogController(app *fiber.App, timeLogService interfaces.TimeLogService) *TimeLogController {
	c := &TimeLogController{
		service: timeLogService,
	}

	routes := app.Group("/api/")
	routes.Post("/logTime", c.LogNewTime)
	//api.Put("/updateLoggedTime", r.UpdateLoggedTime)
	//api.Delete("/deleteLoggedTime", r.DeleteLoggedTime)
	routes.Get("/getLoggedTime", c.GetAllTimeLogs)

	return c
}

func (c *TimeLogController) LogNewTime(dbContext *fiber.Ctx) error {
	controllerRequest := request.TimeLogRequest{}

	err := dbContext.BodyParser(&controllerRequest)
	if err != nil {
		dbContext.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	serviceRequestModel := controllerRequest.ToServiceLogTimeRequest()
	_, err = c.service.AddTimeLog(serviceRequestModel)
	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could'nt log time"})
		return err
	}

	dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{"message": "Time logged successfully!"})

	return nil
}

func (c *TimeLogController) GetAllTimeLogs(dbContext *fiber.Ctx) error {
	existingTimeLogs, err := c.service.GetAllTimeLogs()

	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could'nt fetch logged time"})
		return err
	}

	response := responseMapper.ToTimeLogListResponse(existingTimeLogs.List)

	dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Logged time fetched successfully",
			"data":    response,
		})

	return nil
}
