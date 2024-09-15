package controller

import (
 "github.com/gin-gonic/gin"
 "net/http"
 "gin-gorm-crud/data/request"
 "gin-gorm-crud/data/response"
 "gin-gorm-crud/service"
 "strconv"
)

// StudentController handles HTTP requests related to student operations.
type StudentController struct {
 StudentService service.StudentService
}

// NewStudentController creates and returns a new instance of StudentController.
// It initializes the controller with necessary dependencies.
func NewStudentController(service service.StudentService) *StudentController {
 return &StudentController{StudentService: service}
}

// FindAll retrieves all student records from the database.
func (controller *StudentController) FindAll(ctx *gin.Context) {
 data, err := controller.StudentService.FindAll()

 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   data,
 }
 
 ctx.JSON(http.StatusOK, res)
}


//FindByID finds Student by ID in the Database
func (controller *StudentController) FindByID(ctx *gin.Context) {
 studentID := ctx.Param("ID")
 ID, err := strconv.Atoi(studentID)

 data, err := controller.StudentService.FindByID(ID)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   data,
 }
 ctx.JSON(http.StatusOK, res)
}


// Create student Account from Scratch in the postgres DB
func (controller *StudentController) Create(ctx *gin.Context) {
 req := request.CreateStudentRequest{}
 ctx.ShouldBindJSON(&req)
 
 err := controller.StudentService.Create(req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}


// Update student records in the Postgres
func (controller *StudentController) Update(ctx *gin.Context) {
 req := request.UpdateStudentRequest{}
 err := ctx.ShouldBindJSON(&req)

 studentID := ctx.Param("ID")
 ID, err := strconv.Atoi(studentID)

 _, err = controller.StudentService.FindByID(ID)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }

 req.ID = ID

 err = controller.StudentService.Update(req)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}

// Delete the Student's record in the Database
func (controller *StudentController) Delete(ctx *gin.Context) {
 studentID := ctx.Param("ID")
 ID, err := strconv.Atoi(studentID)

 _, err = controller.StudentService.FindByID(ID)
 if err != nil {
  ctx.JSON(http.StatusNotFound, response.ErrorResponse{
   Code: 404,
   Message: err.Error(),
  })
  return
 }
 
 err = controller.StudentService.Delete(ID)
 if err != nil {
  ctx.JSON(http.StatusInternalServerError, response.ErrorResponse{
   Code: 500,
   Message: err.Error(),
  })
  return
 }

 res := response.Response{
  Code:   200,
  Status: "Ok",
  Data:   nil,
 }

 ctx.JSON(http.StatusOK, res)
}