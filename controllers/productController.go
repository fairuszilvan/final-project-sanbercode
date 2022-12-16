package controllers

import (
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/database"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/repository"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexProduct(c *gin.Context) {
	var (
		result gin.H
	)
	product, err := repository.IndexProduct(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": product,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertProduct(c *gin.Context) {
	var product structs.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}
	err = repository.InsertProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Insert Produk",
	})
}

func UpdateProduct(c *gin.Context) {
	var product structs.Product
	id, _ := strconv.Atoi(c.Param("proID"))
	err := c.ShouldBindJSON(&product)
	if err != nil {
		panic(err)
	}
	product.ID = int(id)
	err = repository.UpdateProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Update Produk",
	})
}
func DeleteProduct(c *gin.Context) {
	var product structs.Product
	id, err := strconv.Atoi(c.Param("proID"))
	product.ID = int(id)
	err = repository.DeleteProduct(database.DbConnection, product)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Delete Produk",
	})
}
