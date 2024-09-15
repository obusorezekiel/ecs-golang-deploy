package repository

import (
 "gin-gorm-crud/model"
)

// StudentRepository defines the methods for interacting with student data in the repository.
type StudentRepository interface {
 FindAll() ([]model.Student, error)
 FindByID(studentID int) (student model.Student, err error)
 Save(student model.Student) error
 Update(student model.Student) error
 Delete(studentID int) error
}