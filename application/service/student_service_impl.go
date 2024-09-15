package service

import (
 "github.com/go-playground/validator/v10"
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
 "gin-gorm-crud/model"
 "gin-gorm-crud/repository"
 "errors"
)

// StudentServiceImpl is the concrete implementation of the StudentService interface,
// providing the business logic for handling student-related operations.
type StudentServiceImpl struct {
 StudentRepository repository.StudentRepository
 Validate      *validator.Validate
}

// NewStudentServiceImpl creates a new instance of StudentServiceImpl with the provided repository.
// It sets up the service layer for handling student-related operations.
func NewStudentServiceImpl(StudentRepository repository.StudentRepository, validate *validator.Validate) (service StudentService, err error) {
 if validate == nil {
  return nil, errors.New("validator instance cannot be nil")
 }
 return &StudentServiceImpl{
  StudentRepository: StudentRepository,
  Validate:       validate,
 }, err
}

// FindAll retrieves all students from the repository.
func (t StudentServiceImpl) FindAll() (students []response.StudentResponse, err error) {
 result, err := t.StudentRepository.FindAll()
 if err != nil {
  return nil, err
 }

 
 for _, value := range result {
	student := response.StudentResponse{
	 ID:   value.ID,
	 Name: value.Name,
	 Description: value.Description,
	}
	students = append(students, student)
   }
   return students, nil
  }

// FindByID retrieves all students information by ID from the repository.
func (t StudentServiceImpl) FindByID(studentID int) (response.StudentResponse, error) {
 data, err := t.StudentRepository.FindByID(studentID)
 if err != nil {
  return response.StudentResponse{}, err
 }

 res := response.StudentResponse{
  ID:   data.ID,
  Name: data.Name,
  Description: data.Description,
 }
 return res, nil
}

// Create adds a new student to the repository.
func (t *StudentServiceImpl) Create(student request.CreateStudentRequest) (err error) {
 err = t.Validate.Struct(student)

 if err != nil {
  return err
 }

 m := model.Student{
  Name: student.Name,
  Description: student.Description,
 }

 t.StudentRepository.Save(m)

 return nil
}

// Update existing information to the repository.
func (t *StudentServiceImpl) Update(student request.UpdateStudentRequest) (err error) {
 data, err := t.StudentRepository.FindByID(student.ID)
 
 if err != nil {
  return err
 }

 data.Name = student.Name
 data.Description = student.Description
 t.StudentRepository.Update(data)
 return nil
}

// Delete existing information to the repository.
func (t *StudentServiceImpl) Delete(studentID int) (err error) {
 err = t.StudentRepository.Delete(studentID)

 if err != nil {
  return err
 }

 return nil
}