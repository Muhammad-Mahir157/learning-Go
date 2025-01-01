package main

import (
	"log"
	"os"

	"github.com/Muhammad-Mahir157/clockify-app-clone/infrastructure/postgresDb"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api"
	"github.com/Muhammad-Mahir157/clockify-app-clone/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

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

	timeLogService := usecase.NewTimeLogService(timeLogRepo)

	app := fiber.New()
	_ = api.NewTimeLogController(app, timeLogService)

	// r := Repository{
	// 	DB: db,
	// }

	//r.SetupRoutes(app)
	app.Listen(":8080")
}
