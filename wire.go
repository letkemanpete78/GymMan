package main

import (
	"exercise"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initAPI(db *gorm.DB) exercise.API {
	wire.Build(exercise.ProvideExerciseRepostiory, exercise.ProvideService, exercise.ProvideAPI)

	return exercise.API{}
}
