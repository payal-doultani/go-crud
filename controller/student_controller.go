package controller

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/payaldoultani/go-crud/managers"
	"github.com/payaldoultani/go-crud/request"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

// Custom email validation function
func validateEmail(fl validator.FieldLevel) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(fl.Field().String())
}

// InitDB initializes the database connection
func InitDB(database *sql.DB) {
	db = database
}

// CreateStudent handles the creation of a new student
func CreateStudent(c echo.Context) error {
	var req request.StudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.Name == "" || req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name and Email are required"})
	}

	// Initialize validator and register custom email validation
	validate := validator.New()
	validate.RegisterValidation("email_regex", validateEmail)

	// Validate the input
	if err := validate.Struct(req); err != nil {
		log.Println("Validation error:", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call manager to create the student
	createdStudent, err := managers.CreateStudent(db, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, createdStudent)
}

// GetAllStudents handles fetching all students
func GetAllStudents(c echo.Context) error {
	students, err := managers.GetAllStudents(db)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve students"})
	}

	return c.JSON(http.StatusOK, students)
}

// GetStudentById handles fetching a specific student by ID
func GetStudentById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	student, err := managers.GetStudentById(db, id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, student)
}

// UpdateStudent handles the update of a student by ID
func UpdateStudent(c echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	var req request.StudentRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	if req.Name == "" || req.Email == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Name and Email are required"})
	}

	// Validate the input
	validate := validator.New()
	validate.RegisterValidation("email_regex", validateEmail)
	if err := validate.Struct(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Call manager to update the student
	updatedStudent, err := managers.UpdateStudent(db, id, &req)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
	}

	return c.JSON(http.StatusOK, updatedStudent)
}

// DeleteStudent handles deleting a student by ID
func DeleteStudent(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid student ID"})
	}

	// Call the manager to delete the student
	err = managers.DeleteStudent(db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Student not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete student"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Student deleted successfully"})
}
