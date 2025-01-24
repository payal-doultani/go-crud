package controller

import (
	"log"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/payaldoultani/go-crud/managers"
	"github.com/payaldoultani/go-crud/request"
	"github.com/payaldoultani/go-crud/response"

	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/labstack/echo/v4"
)
 
var validate = validator.New()
 
func CreateStudent(c echo.Context) error {
    var req request.StudentRequest
 
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request"})
    }
    log.Println("req----->")
 
    if err := validate.Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Validation failed"})
    }
    log.Println("req2----->")
 
    createdStudent, err := managers.CreateStudent(req)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, response.ErrorResponse{Message: "Error creating student"})
    }
 
    return c.JSON(http.StatusCreated, createdStudent)
}

func GetAllStudents(c echo.Context) error {
    var req request.Req
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request parameters"})
    }
 
    if req.PageNo <= 0 {
        req.PageNo = 1
    }
    if req.PageSize <= 0 {
        req.PageSize = 2
    }
 
    if c.QueryParam("per_page") != "" {
        pageSize, err := strconv.Atoi(c.QueryParam("per_page"))
        if err == nil && pageSize > 0 {
            req.PageSize = pageSize
        }
    }
 
    o := orm.NewOrm()
 
    studentListResponse, err := managers.GetAllStudents(o, req)
    if err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Failed to fetch students"})
    }
 
    return c.JSON(http.StatusOK, studentListResponse)
}

func GetStudentById(c echo.Context) error {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid student ID"})
    }
 
    student, err := managers.GetStudentById(id)
    if err != nil {
 
        return c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Student not found"})
    }
 
    return c.JSON(http.StatusOK, student)
}
 
func UpdateStudent(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid ID"})
    }
 
    var req request.StudentRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid request"})
    }
    if err := validate.Struct(req); err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Validation failed"})
    }
 
    updatedStudent, err := managers.UpdateStudent(id, req)
    if err != nil {
       
            return c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Student not found"})
    }
 
    return c.JSON(http.StatusOK, updatedStudent)
}
func DeleteStudent(c echo.Context) error {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        return c.JSON(http.StatusBadRequest, response.ErrorResponse{Message: "Invalid ID"})
    }
    if err := managers.DeleteStudent(id); err != nil {
       
            return c.JSON(http.StatusNotFound, response.ErrorResponse{Message: "Student not found"})
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "Student successfully deleted"})
}
