package main

import (
	"rest-gin-gorm/exercise"

	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func initExerciseAPI(db *gorm.DB) exercise.ExerciseAPI {
	wire.Build(exercise.ProvideExerciseRepostiory, exercise.ProvideExerciseService, exercise.ProvideExerciseAPI)

	return exercise.ExerciseAPI{}
}
