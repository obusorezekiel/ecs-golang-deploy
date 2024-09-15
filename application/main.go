package main

import (
 "github.com/go-playground/validator/v10"
 "net/http"
 "gin-gorm-crud/config"
 "gin-gorm-crud/controller"
 "gin-gorm-crud/utils"
 "gin-gorm-crud/repository"
 "gin-gorm-crud/router"
 "gin-gorm-crud/service"
 "time"
 "log"
)

func main() {
 //Database
 db:= config.DatabaseConnection()
 validate := validator.New()

 //Init Repository
 StudentRepository := repository.NewStudentRepositoryImpl(db)

 //Init Service
 StudentService, err := service.NewStudentServiceImpl(StudentRepository, validate)
 if err != nil {
  // Handle error appropriately, such as logging and exiting
  log.Fatalf("Error initializing student service: %v", err)
 }

 //Init controller
 StudentController := controller.NewStudentController(StudentService)

 //Router
 routes := router.StudentRouter(StudentController)

 server := &http.Server{
  Addr:           ":8888",
  Handler:        routes,
  ReadTimeout:    10 * time.Second,
  WriteTimeout:   10 * time.Second,
  MaxHeaderBytes: 1 << 20,
 }

 err = server.ListenAndServe()
 utils.ErrorPanic(err)

}