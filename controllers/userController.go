package controllers

import (
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/database"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/repository"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexUser(c *gin.Context) {
	var (
		result gin.H
	)
	user, err := repository.IndexUser(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": user,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertUser(c *gin.Context) {
	var user structs.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	reggexType := `(.jpg|.png|.gif)`
	matched, err := regexp.MatchString(reggexType, user.Profile_url)
	if !matched {
		panic(err)
	}
	err = repository.InsertUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Insert User",
	})
}

func UpdateUser(c *gin.Context) {
	var user structs.User
	id, _ := strconv.Atoi(c.Param("userID"))
	err := c.ShouldBindJSON(&user)
	if err != nil {
		panic(err)
	}
	user.ID = int(id)
	err = repository.UpdateUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Update User",
	})
}
func DeleteUser(c *gin.Context) {
	var user structs.User
	id, err := strconv.Atoi(c.Param("userID"))
	user.ID = int(id)
	err = repository.DeleteUser(database.DbConnection, user)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Delete User",
	})
}
