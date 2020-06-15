package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/letkemanpete78/gymman/config"
	"github.com/letkemanpete78/gymman/exercise"
	"github.com/spf13/viper"
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

	/*
		https://github.com/devilsray/golang-viper-config-example
	*/
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	var configuration config.Configuration

	if configErr := viper.ReadInConfig(); configErr != nil {
		log.Fatalf("Error reading config file, %s", configErr)
	}
	configErr := viper.Unmarshal(&configuration)
	if configErr != nil {
		log.Fatalf("unable to decode into struct, %v", configErr)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)
	/* */

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
