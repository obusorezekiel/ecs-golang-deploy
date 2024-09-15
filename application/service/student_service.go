package service

import (
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
)

// StudentService defines the contract for student-related business logic.
type StudentService interface {
 Create(student request.CreateStudentRequest) error
 Update(student request.UpdateStudentRequest) error
 Delete(studentID int) error
 FindByID(studentID int) (student response.StudentResponse, err error)
 FindAll() (students []response.StudentResponse, err error)
}