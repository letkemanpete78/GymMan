package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/letkemanpete78/gymman/exercise"
)

func initDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", os.Getenv("exercise.db"))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&exercise.Exercise{})

	return db
}

func main() {
	db := initDB()
	defer db.Close()

	exerciseAPI := InitAPI(db)

	r := gin.Default()

	r.GET("/exercises", exerciseAPI.FindAll)
	r.GET("/exercises/:id", exerciseAPI.FindByID)
	r.POST("/exercises", exerciseAPI.Create)
	r.PUT("/exercises/:id", exerciseAPI.Update)
	r.DELETE("/exercises/:id", exerciseAPI.Delete)

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
