package main

import (

	"./config"
	"./app/controllers"
	_ "./app/models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()
	PersonController := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/person/:id", PersonController.GetPerson)
	router.GET("/persons", PersonController.GetPersons)
	router.POST("/person", PersonController.CreatePerson)
	router.PUT("/person", PersonController.UpdatePerson)
	router.DELETE("/person/:id", PersonController.DeletePerson)
	router.Run(":3000")
}