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

func IndexLaporan(c *gin.Context) {
	var (
		result gin.H
	)
	laporan, err := repository.IndexLaporan(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": laporan,
		}
	}
	c.JSON(http.StatusOK, result)
}

func InsertLaporan(c *gin.Context) {
	var laporan structs.Laporan
	err := c.ShouldBindJSON(&laporan)
	if err != nil {
		panic(err)
	}

	reggexType := `(.png)`
	matched, err := regexp.MatchString(reggexType, laporan.Bukti_pengesahan)
	if !matched {
		panic(err)
	}
	err = repository.InsertLaporan(database.DbConnection, laporan)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Insert Laporan",
	})
}

func UpdateLaporan(c *gin.Context) {
	var laporan structs.Laporan
	id, _ := strconv.Atoi(c.Param("laporanID"))
	err := c.ShouldBindJSON(&laporan)
	if err != nil {
		panic(err)
	}
	laporan.ID = int(id)
	err = repository.UpdateLaporan(database.DbConnection, laporan)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Update Laporan",
	})
}
func DeleteLaporan(c *gin.Context) {
	var laporan structs.Laporan
	id, err := strconv.Atoi(c.Param("laporanID"))
	laporan.ID = int(id)
	err = repository.DeleteLaporan(database.DbConnection, laporan)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"result": "Sucess Delete Laporan",
	})
}
func CreateView(c *gin.Context) {
	var (
		result gin.H
	)
	err := repository.CreateView(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	}
	c.JSON(http.StatusOK, result)
}
func ViewLaporan(c *gin.Context) {
	var (
		result gin.H
	)
	view, err := repository.ViewLaporan(database.DbConnection)
	if err != nil {
		result = gin.H{
			"result": err,
		}
	} else {
		result = gin.H{
			"result": view,
		}
	}
	c.JSON(http.StatusOK, result)
}
