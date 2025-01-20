package main

import (
	"log"

	"github.com/payaldoultani/go-crud/controller"
	"github.com/payaldoultani/go-crud/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}()

	database, err := db.GetDB()
	if err != nil {
		log.Fatalf("Error getting database instance: %v", err)
	}
	controller.InitDB(database)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", database)
			return next(c)
		}
	})
	e.POST("/student", controller.CreateStudent)
	e.GET("/students", controller.GetAllStudents)
	e.GET("/student/:id", controller.GetStudentById)
	e.PUT("/student/:id", controller.UpdateStudent)
	e.DELETE("/student/:id", controller.DeleteStudent)

	e.Logger.Fatal(e.Start(":8090"))
}
