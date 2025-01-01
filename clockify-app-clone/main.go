package main

import (
	"github.com/Muhammad-Mahir157/clockify-app-clone/database"
	"github.com/Muhammad-Mahir157/clockify-app-clone/interface/api"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// err = postgresDb.MigrateTimelog(db)
	// if err != nil {
	// 	log.Fatal("Could not migrate the database ...")
	// }

	database.ConnectToDatabase()

	app := fiber.New()
	api.SetupRoutes(app)

	//r.SetupRoutes(app)
	app.Listen(":8080")
}
