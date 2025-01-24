package main

import (
	"log"

	"github.com/payaldoultani/go-crud/config"
	"github.com/payaldoultani/go-crud/controller"
	"github.com/payaldoultani/go-crud/db"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	mysqlCfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load MySQL configuration: %v", err)
	}

	err = db.InitDB(mysqlCfg)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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
