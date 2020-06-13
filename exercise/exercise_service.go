package exercise

// ExerciseService is the object used for the repository
type ExerciseService struct {
	ExerciseRepository ExerciseRepository
}

// ProvideExerciseService is the provider used for the repoitory service
func ProvideExerciseService(p ExerciseRepository) ExerciseService {
	return ExerciseService{ExerciseRepository: p}
}

// FindAll is the service method to returns all exercise objects from the database
func (p *ExerciseService) FindAll() []Exercise {
	return p.ExerciseRepository.FindAll()
}

// FindByID is the service method to return exercise  from the database given the ID value
func (p *ExerciseService) FindByID(id uint) Exercise {
	return p.ExerciseRepository.FindByID(id)
}

// Save is the service method to save the exercise object into the database
func (p *ExerciseService) Save(exercise Exercise) Exercise {
	p.ExerciseRepository.Save(exercise)

	return exercise
}

// Delete is the service method to remove the exercise object from the database
func (p *ExerciseService) Delete(exercise Exercise) {
	p.ExerciseRepository.Delete(exercise)
}
