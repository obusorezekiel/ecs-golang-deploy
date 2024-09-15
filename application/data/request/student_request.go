package request

// CreateStudentRequest represents the data required to create a new student record.
// It contains fields for the student's name, and description.
type CreateStudentRequest struct {
 Name string `validate:"required,min=1,max=200" json:"name"`
 Description string `validate:"required,min=1,max=200" json:"description"`
}

// UpdateStudentRequest represents the data required to update student record.
// It contains fields for the student's name, description, and other necessary details.
type UpdateStudentRequest struct {
 ID   int    `validate:"required"`
 Name string `validate:"required,max=200,min=1" json:"name"`
 Description string `validate:"required,max=200,min=1" json:"description"`
}