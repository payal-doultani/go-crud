package managers

import (
	"errors"
	"fmt"
	"log"

	"github.com/beego/beego/v2/client/orm"
	"github.com/payaldoultani/go-crud/models"
	"github.com/payaldoultani/go-crud/request"
	"github.com/payaldoultani/go-crud/response"
)

func CreateStudent(req request.StudentRequest) (response.StudentResponse, error) {
	o := orm.NewOrm()
	log.Println("reqmanagers----->")

	var student = models.Student{
		Name:  req.Name,
		Email: req.Email,
	}
	log.Println("reqmanagers 2----->")
	_, err := o.Insert(&student)
	if err != nil {
		return response.StudentResponse{}, err
	}
	responseStudent := &response.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
	}

	return *responseStudent, nil
}

func GetAllStudents(o orm.Ormer, req request.Req) (response.GetAllStudentsResponse, error) {
	var students []response.StudentResponse

	query := o.QueryTable(&models.Student{})

	if req.Filter != "" {
		query = query.Filter("name__icontains", req.Filter) 
	}

	fmt.Printf("PageNo: %d, PageSize: %d\n", req.PageNo, req.PageSize)

	offset := (req.PageNo - 1) * req.PageSize
	query = query.Limit(req.PageSize).Offset(offset)

	fmt.Printf("Limit: %d, Offset: %d\n", req.PageSize, offset)

	if _, err := query.All(&students); err != nil {
		return response.GetAllStudentsResponse{}, fmt.Errorf("failed to fetch students: %v", err)
	}

	countQuery := o.QueryTable(&models.Student{})
	if req.Filter != "" {
		countQuery = countQuery.Filter("name__icontains", req.Filter)
	}

	total, err := countQuery.Count()
	if err != nil {
		return response.GetAllStudentsResponse{}, fmt.Errorf("failed to count students: %v", err)
	}

	lastPage := (total + int64(req.PageSize) - 1) / int64(req.PageSize)
	if lastPage == 0 {
		lastPage = 1
	}

	return response.GetAllStudentsResponse{
		Students:    students,
		PageNo:      req.PageNo,
		PageSize:    req.PageSize,
		TotalCount:  int(total),
		LastPage:    int(lastPage),
		CurrentPage: req.PageNo,
	}, nil
}

func GetStudentById(id int) (response.StudentResponse, error) {
	o := orm.NewOrm()

	var student models.Student

	if err := o.QueryTable(new(models.Student)).Filter("ID", id).One(&student); err != nil {
		if err == orm.ErrNoRows {
			return response.StudentResponse{}, fmt.Errorf("student not found")
		}
		return response.StudentResponse{}, fmt.Errorf("failed to fetch student: %v", err)
	}

	studentResponse := response.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
	}

	return studentResponse, nil
}

func UpdateStudent(id int, req request.StudentRequest) (response.StudentResponse, error) {
	o := orm.NewOrm()

	student := models.Student{ID: id}
	err := o.Read(&student)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("student not found")
	}

	student.Name = req.Name
	student.Email = req.Email

	_, err = o.Update(&student)
	if err != nil {
		return response.StudentResponse{}, fmt.Errorf("failed to update student: %v", err)
	}
	responseStudent := &response.StudentResponse{
		ID:    student.ID,
		Name:  student.Name,
		Email: student.Email,
	}

	return *responseStudent, nil
}

func DeleteStudent(id int) error {
	o := orm.NewOrm()
	existingStudent := models.Student{ID: id}
	if err := o.Read(&existingStudent); err != nil {
		if err == orm.ErrNoRows {
			return errors.New("student not found")
		}
		return err
	}
	if _, err := o.Delete(&existingStudent); err != nil {
		return err
	}
	return nil
}
