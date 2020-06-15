package main

import (
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
	"github.com/letkemanpete78/gymman/exercise"
)

func InitAPI(db *gorm.DB) exercise.API {
	wire.Build(exercise.ProvideExerciseRepostiory, exercise.ProvideService, exercise.ProvideAPI)

	return exercise.API{}
}
