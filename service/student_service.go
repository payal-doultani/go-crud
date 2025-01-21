package service

import (
	"database/sql"
	"fmt"

	"github.com/payaldoultani/go-crud/models"
	"github.com/payaldoultani/go-crud/response"
)

func CreateStudent(db *sql.DB, student *models.Student) (*response.StudentResponse, error) {
	query := "INSERT INTO student (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, student.Name, student.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to create student: %v", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get the last inserted student ID: %v", err)
	}

	createdStudent := &response.StudentResponse{
		ID:    int(lastInsertID),
		Name:  student.Name,
		Email: student.Email,
	}

	return createdStudent, nil
}

func GetAllStudents(db *sql.DB) ([]response.StudentResponse, int, error) {
	query := "SELECT * FROM student"
	rows, err := db.Query(query)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get the student list: %v", err)
	}

	var students []response.StudentResponse
	for rows.Next() {
		var student response.StudentResponse
		if err := rows.Scan(&student.ID, &student.Name, &student.Email); err != nil {
			return nil, 0, fmt.Errorf("scan error: %v", err)
		}

		students = append(students, student)
	}

	var total int
	totalQuery := "SELECT FOUND_ROWS()"
	if err := db.QueryRow(totalQuery).Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("total count error: %v", err)
	}

	return students, total, nil
}

func GetStudentById(db *sql.DB, id int) (response.StudentResponse, error) {
	selectStatement := `
        SELECT * 
        FROM student 
        WHERE id = ?`

	var student response.StudentResponse

	err := db.QueryRow(selectStatement, id).Scan(
		&student.ID,
		&student.Name,
		&student.Email,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return response.StudentResponse{}, fmt.Errorf("student with id %d not found", id)
		}
		return response.StudentResponse{}, fmt.Errorf("failed to retrieve student: %v", err)
	}

	return student, nil
}

func UpdateStudent(db *sql.DB, student *models.Student, id int) (response.StudentResponse, error) {
	sqlStatement := `
        UPDATE student
        SET name = ?, email = ?
        WHERE id = ?`

	_, err := db.Exec(sqlStatement, student.Name, student.Email, id)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("failed to update student: %v", err)
	}

	selectStatement := `
        SELECT *
        FROM student
        WHERE id = ?`

	var updatedStudent response.StudentResponse
	err = db.QueryRow(selectStatement, id).Scan(
		&updatedStudent.ID,
		&updatedStudent.Name,
		&updatedStudent.Email,
	)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("failed to retrieve updated student: %v", err)
	}

	return updatedStudent, nil
}

func DeleteStudent(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM student WHERE id = ?", id)
	if err != nil {
		return fmt.Errorf("failed to delete student: %v", err)
	}
	
	return nil
}
