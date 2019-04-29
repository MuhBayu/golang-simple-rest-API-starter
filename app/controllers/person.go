package controllers

import (
	"net/http"
	
	"../models"
	"github.com/gin-gonic/gin"
)

func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	var id string = c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"data": err.Error(),
			"success": false,
		}
	} else {
		result = gin.H{
			"data": person,
			"success": true,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []models.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"data": nil,
			"success": true,
		}
	} else {
		result = gin.H{
			"data": persons,
			"count":  len(persons),
			"success": false,
		}
	}

	c.JSON(http.StatusOK, result)
}

func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	first_name, last_name := c.PostForm("first_name"), c.PostForm("last_name")
	person.First_Name = first_name
	person.Last_Name = last_name
	idb.DB.Create(&person)
	result = gin.H{
		"data": person,
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")
	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	var (
		person    models.Person
		newPerson models.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "data not found",
		}
	}
	newPerson.First_Name = first_name
	newPerson.Last_Name = last_name
	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "update failed",
		}
	} else {
		result = gin.H{
			"success": true,
			"message": "successfully updated data",
		}
	}
	c.JSON(http.StatusOK, result)
}

func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person models.Person
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "data not found",
		}
	}
	err = idb.DB.Delete(&person).Error
	if err != nil {
		result = gin.H{
			"success": false,
			"message": "delete failed",
		}
	} else {
		result = gin.H{
			"success": true,
			"message": "Data deleted successfully",
		}
	}

	c.JSON(http.StatusOK, result)
}
