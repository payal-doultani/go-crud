package managers

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/payaldoultani/go-crud/models"
	"github.com/payaldoultani/go-crud/request"
	"github.com/payaldoultani/go-crud/response"
	"github.com/payaldoultani/go-crud/service"
)

func CreateStudent(db *sql.DB, req request.StudentRequest) (*response.StudentResponse, error) {

	student := models.Student{
		Name:  req.Name,
		Email: req.Email,
	}

	log.Println("Manager: Creating student...")

	createdStudent, err := service.CreateStudent(db, &student)
	if err != nil {
		return nil, fmt.Errorf("failed to create student: %v", err)
	}

	return createdStudent, nil
}

func GetAllStudents(db *sql.DB) ([]response.StudentResponse, error) {
	log.Printf("Manager: Retrieving all students")

	// Call the service layer to get all students
	students, _, err := service.GetAllStudents(db)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve all students: %v", err)
	}

	return students, nil
}

func GetStudentById(db *sql.DB, id int) (response.StudentResponse, error) {
	log.Printf("Manager: Retrieving student with ID: %d", id)

	// Call the service layer to get the student by ID
	student, err := service.GetStudentById(db, id)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("failed to retrieve student with ID %d: %v", id, err)
	}

	return student, nil
}

func UpdateStudent(db *sql.DB, id int, req *request.StudentRequest) (response.StudentResponse, error) {
	student := &models.Student{
		Name:  req.Name,
		Email: req.Email,
	}

	updatedStudent, err := service.UpdateStudent(db, student, id)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("failed to update student: %v", err)
	}

	return updatedStudent, nil
}

func DeleteStudent(db *sql.DB, id int) error {

	err := service.DeleteStudent(db, id)
	if err != nil {
		return fmt.Errorf("failed to delete student: %v", err)
	}
	return nil
}
