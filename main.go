package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/letkemanpete78/gymman/config"
	"github.com/letkemanpete78/gymman/exercise"
	"github.com/spf13/viper"
)

func initDB() *gorm.DB {
	// db, err := gorm.Open("sqlite3", os.Getenv("exercise.db"))
	db, err := gorm.Open("sqlite3", "/home/pete/Desktop/gymman/exercise.db")
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

	/* */

	db := initDB()
	defer db.Close()

	exerciseAPI := InitAPI(db)

	r := gin.Default()

	exerciseVersion := configuration.APIVersions.Exercise
	rgAPI := r.Group("/exercises/" + exerciseVersion)
	rgAPI.GET("/", exerciseAPI.FindAll)
	rgAPI.GET("/:uuid", exerciseAPI.FindByUUID)
	rgAPI.POST("/", exerciseAPI.Create)
	rgAPI.PUT("/:uuid", exerciseAPI.Update)
	rgAPI.DELETE("/:uuid", exerciseAPI.Delete)

	r.LoadHTMLGlob("templates/exercise/*")
	rg := r.Group("/exercises/")
	rg.GET("/create", renderCreateExercise)
	rg.GET("/list", renderListExcercises)

	var port = getPort(&configuration)
	r.Run(port)
}

func renderCreateExercise(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", gin.H{"Exercises": nil})
}

func renderListExcercises(c *gin.Context) {
	var values []exercise.Exercise
	db := initDB()
	defer db.Close()
	db.Find(&exercise.Exercise{}).Scan(&values)
	c.HTML(http.StatusOK, "list.html", gin.H{"Exercises": values})
}

func getPort(config *config.Configuration) string {
	log.Printf("port for this application is %d", config.Server.Port)
	return ":" + strconv.FormatUint(config.Server.Port, 10)
}
