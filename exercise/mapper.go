package exercise

// ToExercise converts an DTO object to Exercise object
func ToExercise(exerciseDTO DTO) Exercise {
	return Exercise{Name: exerciseDTO.Name, Description: exerciseDTO.Description, UUID: exerciseDTO.UUID}
}

// ToDTO converts an Exercise object to DTO object
func ToDTO(exercise Exercise) DTO {
	return DTO{Name: exercise.Name, Description: exercise.Description, UUID: exercise.UUID}
}

// ToDTOs converts an array of DTO objects to an array of Exercise objects
func ToDTOs(exercises []Exercise) []DTO {
	exercisedtos := make([]DTO, len(exercises))

	for i, itm := range exercises {
		exercisedtos[i] = ToDTO(itm)
	}

	return exercisedtos
}
