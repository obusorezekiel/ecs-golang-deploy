package repository

import (
 "errors"
 "gin-gorm-crud/model"
 "gin-gorm-crud/data/request"
 "gorm.io/gorm"
)

// StudentRepositoryImpl defines the methods for interacting with student data in the repository.
type StudentRepositoryImpl struct {
 Db *gorm.DB
}

// NewStudentRepositoryImpl creates a new instance of StudentRepositoryImpl.
// It returns a pointer to the newly created StudentRepository
func NewStudentRepositoryImpl(Db *gorm.DB) (StudentRepository) {
 return &StudentRepositoryImpl{Db: Db}
}

// FindAll is an implementation of the StudentRepository interface.
func (t StudentRepositoryImpl) FindAll() (students []model.Student, err error) {
 results := t.Db.Find(&students)
 if results.Error != nil {
  return nil, results.Error
 }
 
 return students, nil
}

// FindByID retrieves a student by their ID from the repository.
func (t StudentRepositoryImpl) FindByID(studentID int) (student model.Student, err error) {
 result := t.Db.Find(&student, studentID)
 
 if result.Error != nil {
  return model.Student{}, result.Error
 }

 if result.RowsAffected == 0 {
  return model.Student{}, errors.New("Student is not found")
 }

 return student, nil
}

// Save stores a new student or updates an existing student's details in the repository
func (t *StudentRepositoryImpl) Save(student model.Student) error {
 result := t.Db.Create(&student)
 if result.Error != nil {
  return result.Error
 }

 return nil
}

// Update modifies an existing student's details in the repository.
func (t *StudentRepositoryImpl) Update(student model.Student) error {
 var data = request.UpdateStudentRequest{
  ID:   student.ID,
  Name: student.Name,
  Description: student.Description,
 }

 result := t.Db.Model(&student).Updates(data)
 if result.Error != nil {
  return result.Error
 }

 return nil
}

// Delete removes a student from the repository by their ID.
func (t *StudentRepositoryImpl) Delete(studentID int) error {
 var student model.Student

 result := t.Db.Where("id = ?", studentID).Delete(&student)
 if result.Error != nil {
  return result.Error
 }

 return nil
}