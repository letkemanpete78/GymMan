package exercise

// ToExercise converts an ExerciseDTO object to Exercise object
func ToExercise(exerciseDTO ExerciseDTO) Exercise {
	return Exercise{Name: exerciseDTO.Name, Description: exerciseDTO.Description, Image: exerciseDTO.Image}
}

// ToExerciseDTO converts an Exercise object to ExerciseDTO object
func ToExerciseDTO(exercise Exercise) ExerciseDTO {
	return ExerciseDTO{ID: exercise.ID, Name: exercise.Name, Description: exercise.Description, Image: exercise.Image}
}

// ToExerciseDTOs converts an array of ExerciseDTO objects to an array of Exercise objects
func ToExerciseDTOs(exercises []Exercise) []ExerciseDTO {
	exercisedtos := make([]ExerciseDTO, len(exercises))

	for i, itm := range exercises {
		exercisedtos[i] = ToExerciseDTO(itm)
	}

	return exercisedtos
}
