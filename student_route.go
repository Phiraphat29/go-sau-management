// Create API Endpoint
package main

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func studentRoute(app *fiber.App, db *gorm.DB) {
	// * C (POST /api/students)
	app.Post("/api/students", func(c fiber.Ctx) error {
		// Stored value from client
		student := new(Student)

		// bind data
		if err := c.Bind().Body(student); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		// validate
		if err := createStudent(db, student); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		// if data successfully insert to database
		return c.Status(201).JSON(fiber.Map{
			"message": "Created",
			"data":    student,
		})
	})

	// * R (GET /api/students)
	app.Get("/api/students", func(c fiber.Ctx) error {
		// stored data for showing in client
		var students []Student

		// * -------------------------------------------
		// stored data for query string condition
		StuName := c.Query("StuName")
		StuLevel := c.Query("StuLevel")

		var err error

		if StuName != "" {
			// fetching StuName condition
			students, err = getStudentByName(db, StuName)
		} else if StuLevel != "" {
			// fetching StuLevel condition
			StuLevelInt, err := strconv.Atoi(StuLevel)
			if err != nil {
				return c.Status(400).JSON(fiber.Map{
					"message": "Invalid request parameter",
				})
			}
			students, err = getStudentByLevel(db, StuLevelInt)
		} else {
			// stored data
			students, err = getAllStudent(db)
		}
		// * -------------------------------------------

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		// validate
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    students,
		})
	})

	// * R (GET /api/students/:id)
	app.Get("/api/students/:id", func(c fiber.Ctx) error {
		// stored PK
		id := c.Params("id")

		// convert ID to number
		idInt, err := strconv.Atoi(id)

		// fetching data and stored in student
		student, err := getStudentByID(db, idInt)

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).JSON(fiber.Map{
				"message": "Student not found",
			})
		}

		// validate
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Success",
			"data":    student,
		})
	})

	// * U (PUT /api/students/:id)
	app.Put("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		student := new(Student)

		if err := c.Bind().Body(student); err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request body",
			})
		}

		if err := updateStudent(db, idInt, student); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Student Update Successfully",
			"data":    student,
		})
	})

	// * D (DELETE /api/students/:id)
	app.Delete("/api/students/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		idInt, err := strconv.Atoi(id)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request parameter",
			})
		}

		if err := deleteStudent(db, idInt); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}

		return c.Status(200).JSON(fiber.Map{
			"message": "Student delete successfully",
		})
	})
}
