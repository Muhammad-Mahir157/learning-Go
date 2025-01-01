package api

import (
	"net/http"

	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/adaptor"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api/dto/models"
	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type TimeLogController struct {
	service usecase.TimeLogUsecase
}

func NewTimeLogController(app *fiber.App, timeLogService usecase.TimeLogUsecase) *TimeLogController {
	c := &TimeLogController{
		service: timeLogService,
	}

	routes := app.Group("/api/")
	routes.Post("/logTime", c.LogNewTime)
	routes.Put("/updateLoggedTime", c.UpdateLoggedTime)
	routes.Delete("/deleteLoggedTime/:Id", c.DeleteLoggedTime)
	routes.Get("/getLoggedTime", c.GetAllTimeLogs)
	routes.Get("/getLoggedTimeById/:Id", c.GetLoggedTimeById)

	return c
}

func (c *TimeLogController) LogNewTime(dbContext *fiber.Ctx) error {
	controllerRequest := models.AddTimeLogRequest{}

	err := dbContext.BodyParser(&controllerRequest)
	if err != nil {
		dbContext.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed"})
		return err
	}

	addRequestModel := controllerRequest.ToAddLogTimeRequest()
	newlyAddedTimeLog, err := c.service.AddTimeLog(addRequestModel)
	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could'nt add time log"})
		return err
	}

	response := adaptor.ToTimeLogResponse(newlyAddedTimeLog)
	return dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Time log added successfully!",
			"data":    response,
		})

}

func (c *TimeLogController) GetAllTimeLogs(dbContext *fiber.Ctx) error {
	existingTimeLogs, err := c.service.GetAllTimeLogs()

	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could'nt fetch logged time"})
		return err
	}

	response := adaptor.ToTimeLogListResponse(existingTimeLogs.List)

	return dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Logged time fetched successfully",
			"data":    response,
		})
}

func (c *TimeLogController) UpdateLoggedTime(dbContext *fiber.Ctx) error {
	controllerRequest := models.UpdateTimeLogRequest{}

	err := dbContext.BodyParser(&controllerRequest)
	if err != nil {
		dbContext.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Couldn't process request"})
		return err
	}

	updateRequestModel := controllerRequest.ToUpdateLogTimeRequest()

	timeLogUpdated, err := c.service.UpdateTimeLog(updateRequestModel)
	if err != nil {
		dbContext.Status(http.StatusNotFound).JSON(
			&fiber.Map{
				"message": "Could'nt update logged time against provided Id",
			})
		return err
	}

	//mapping the timeLogFound to response ...
	response := adaptor.ToTimeLogResponse(timeLogUpdated)

	return dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Time Log successfully updated",
			"data":    response,
		})
}

func (c *TimeLogController) DeleteLoggedTime(dbContext *fiber.Ctx) error {
	idRaw := dbContext.Params("Id")

	if idRaw == "" {
		dbContext.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Id cannot be empty",
			})
		return nil
	}

	id, err := uuid.Parse(idRaw)
	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Invalid timeLog Id format",
			})
	}

	timeLogDeleted, err := c.service.DeleteTimeLog(id)
	if err != nil {
		dbContext.Status(http.StatusNotFound).JSON(
			&fiber.Map{
				"message": "Could'nt delete logged time against provided Id",
			})
		return err
	}

	return dbContext.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "Time Log successfully deleted",
			"data":    timeLogDeleted,
		})
}

func (c *TimeLogController) GetLoggedTimeById(dbContext *fiber.Ctx) error {
	idRaw := dbContext.Params("Id")

	if idRaw == "" {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Id cannot be empty",
			})
		return nil
	}

	id, err := uuid.Parse(idRaw)
	if err != nil {
		dbContext.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Invalid timeLog Id format",
			})
	}

	timeLogFound, err := c.service.GetTimeLogById(id)
	if err != nil {
		return dbContext.Status(http.StatusNotFound).JSON(
			&fiber.Map{
				"message": "Time Log not found against provided Id",
			})
	}

	//mapping the timeLogFound to response ...
	response := adaptor.ToTimeLogResponse(timeLogFound)
	return dbContext.Status(http.StatusFound).JSON(
		&fiber.Map{
			"message": "Time Log found successfully",
			"data":    response,
		})
}
