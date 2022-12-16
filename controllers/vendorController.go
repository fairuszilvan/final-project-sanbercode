package controllers

import (
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/database"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/repository"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexVendor(c *gin.Context) {
	var (
		result gin.H
	)
	vendor, err := repository.IndexVendor(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": vendor,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertVendor(c *gin.Context) {
	var vendor structs.Vendor
	err := c.ShouldBindJSON(&vendor)
	if err != nil {
		panic(err)
	}
	err = repository.InsertVendor(database.DbConnection, vendor)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Insert Vendor",
	})
}

func UpdateVendor(c *gin.Context) {
	var vendor structs.Vendor
	id, _ := strconv.Atoi(c.Param("vendorID"))
	err := c.ShouldBindJSON(&vendor)
	if err != nil {
		panic(err)
	}
	vendor.ID = int(id)
	err = repository.UpdateVendor(database.DbConnection, vendor)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Update Vendor",
	})
}
func DeleteVendor(c *gin.Context) {
	var vendor structs.Vendor
	id, err := strconv.Atoi(c.Param("vendorID"))
	vendor.ID = int(id)
	err = repository.DeleteVendor(database.DbConnection, vendor)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Delete Vendor",
	})
}
