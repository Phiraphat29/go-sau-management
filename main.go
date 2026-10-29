package main

import (
	"fmt"

	"github.com/gofiber/fiber/v3"
)

func main() {
	db, err := ConnectDatabase()

	if err != nil {
		fmt.Println(err, db)
	}

	app := fiber.New()

	studentRoute(app, db)

	if err := app.Listen(":1112"); err != nil {
		fmt.Println(err)
	}
}
