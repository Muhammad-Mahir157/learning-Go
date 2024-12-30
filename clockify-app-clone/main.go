package main

import (
	"log"
	"os"

	"github.com/Muhammad-Mahir157/clockify-app-clone/application/services"
	"github.com/Muhammad-Mahir157/clockify-app-clone/infrastructure/postgresDb"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// type Repository struct {
// 	DB *gorm.DB
// }

// type TimeLog struct {
// 	StartTime   string `json:"startTime"`
// 	EndTime     string `json:"endTime"`
// 	Date        string `json:"date"`
// 	LoggedHours uint   `json:"loggedHours"`
// 	Description string `json:"description"`
// }

// func (r *Repository) LogNewTime(dbContext *fiber.Ctx) error {
// 	logTime := TimeLog{}

// 	err := dbContext.BodyParser(&logTime)
// 	if err != nil {
// 		dbContext.Status(http.StatusUnprocessableEntity).JSON(
// 			&fiber.Map{"message": "request failed"})
// 		return err
// 	}

// 	err = r.DB.Create(&logTime).Error
// 	if err != nil {
// 		dbContext.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "Could'nt log time"})
// 		return err
// 	}

// 	dbContext.Status(http.StatusOK).JSON(
// 		&fiber.Map{"message": "Time logged successfully!"})

// 	return nil
// }

// func (r *Repository) GetLoggedTime(dbContext *fiber.Ctx) error {
// 	loggedTimeList := &[]entities.TimeLog{}

// 	err := r.DB.Find(loggedTimeList).Error
// 	if err != nil {
// 		dbContext.Status(http.StatusBadRequest).JSON(
// 			&fiber.Map{"message": "Could'nt fetch logged time"})
// 		return err
// 	}

// 	dbContext.Status(http.StatusOK).JSON(
// 		&fiber.Map{
// 			"message": "Logged time fetched successfully",
// 			"data":    loggedTimeList,
// 		})

// 	return nil
// }

// func (r *Repository) SetupRoutes(app *fiber.App) {
// 	api := app.Group("/api/")
// 	api.Post("/logTime", r.LogNewTime)
// 	//api.Put("/updateLoggedTime", r.UpdateLoggedTime)
// 	//api.Delete("/deleteLoggedTime", r.DeleteLoggedTime)
// 	api.Get("/getLoggedTime", r.GetLoggedTime)
// }

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	config := &postgresDb.Config{
		Host:     os.Getenv("DB_Host"),
		Port:     os.Getenv("DB_Port"),
		User:     os.Getenv("DB_User"),
		Password: os.Getenv("DB_Password"),
		DBName:   os.Getenv("DB_Name"),
		SSLMode:  os.Getenv("DB_Mode"),
	}

	db, err := postgresDb.NewConnection(config)
	if err != nil {
		log.Fatal("Could not load the database ...")
	}

	err = postgresDb.MigrateTimelog(db)
	if err != nil {
		log.Fatal("Could not migrate the database ...")
	}

	timeLogRepo := postgresDb.NewTimeLogRepository(db)

	timeLogService := services.NewTimeLogService(timeLogRepo)

	app := fiber.New()
	_ = api.NewTimeLogController(app, timeLogService)

	// r := Repository{
	// 	DB: db,
	// }

	//r.SetupRoutes(app)
	app.Listen(":8080")
}
