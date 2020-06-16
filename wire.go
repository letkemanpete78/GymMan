package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/letkemanpete78/gymman/exercise"
)

// initAPI is the basic db inject method
func initAPI(db *gorm.DB) exercise.API {
	wire.Build(exercise.ProvideExerciseRepostiory, exercise.ProvideService, exercise.ProvideAPI)

	return exercise.API{}
}
