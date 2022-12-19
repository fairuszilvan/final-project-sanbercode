package main

import (
	"database/sql"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/controllers"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/database"
	"fairuszilvan-final-project-bds-sanbercode-golang-batch-40/middleware"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	db  *sql.DB
	err error
)

func main() {
	err = godotenv.Load("config/.env")
	if err != nil {
		fmt.Println("failed load file environment")
	} else {
		fmt.Println("success load file environment")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err = sql.Open("postgres", psqlInfo)
	err = db.Ping()
	if err != nil {
		fmt.Println("DB Connection Failed")
		panic(err)
	} else {
		fmt.Println("DB Connection Success")
	}

	database.DbMigrate(db)

	defer db.Close()

	router := gin.Default()
	router.GET("/products", middleware.Login, controllers.IndexProduct)
	router.POST("/products", middleware.Login, controllers.InsertProduct)
	router.PUT("/products/products-update/:proID", middleware.Login, controllers.UpdateProduct)
	router.DELETE("/products/products-delete/:proID", middleware.Login, controllers.DeleteProduct)

	router.GET("/vendor", middleware.Login, controllers.IndexVendor)
	router.POST("/vendor", middleware.Login, controllers.InsertVendor)
	router.PUT("/vendor/vendor-update/:vendorID", middleware.Login, controllers.UpdateVendor)
	router.DELETE("/vendor/vendor-delete/:vendorID", middleware.Login, controllers.DeleteVendor)

	router.GET("/users", middleware.Login, controllers.IndexUser)
	router.POST("/users", controllers.InsertUser)
	router.PUT("/users/users-update/:userID", middleware.Login, controllers.UpdateUser)
	router.DELETE("/users/users-delete/:userID", middleware.Login, controllers.DeleteUser)

	router.GET("/laporan", middleware.Login, controllers.IndexLaporan)
	router.POST("/laporan", middleware.Login, controllers.InsertLaporan)
	router.PUT("/laporan/laporan-update/:laporanID", middleware.Login, controllers.UpdateLaporan)
	router.DELETE("/laporan/laporan-delete/:laporanID", middleware.Login, controllers.DeleteLaporan)

	router.GET("/view", middleware.Login, controllers.CreateView)
	router.GET("/view/show", middleware.Login, controllers.ViewLaporan)

	router.Run(":" + os.Getenv("PORT"))
}
