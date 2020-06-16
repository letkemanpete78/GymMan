package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

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
	exerciseVersion := configuration.APIVersions.Exercise
	r.GET("/exercises/"+exerciseVersion, exerciseAPI.FindAll)
	r.GET("/exercises/"+exerciseVersion+"/:uuid", exerciseAPI.FindByUUID)
	r.POST("/exercises/"+exerciseVersion, exerciseAPI.Create)
	r.PUT("/exercises/"+exerciseVersion+"/:uuid", exerciseAPI.Update)
	r.DELETE("/exercises/"+exerciseVersion+"/:uuid", exerciseAPI.Delete)

	var port string = ":" + strconv.FormatUint(configuration.Server.Port, 10)
	err := http.ListenAndServe(port, r)
	if err != nil {
		panic(err)
	}
}
