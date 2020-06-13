package exercise

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// ExerciseRepository gorm db object
type ExerciseRepository struct {
	DB *gorm.DB
}

// ProvideExerciseRepostiory is the provider repostiory object
func ProvideExerciseRepostiory(DB *gorm.DB) ExerciseRepository {
	return ExerciseRepository{DB: DB}
}

// FindAll is the repository method to returns all exercise objects from the database
func (p *ExerciseRepository) FindAll() []Exercise {
	var exercises []Exercise
	p.DB.Find(&exercises)

	return exercises
}

// FindByID is the repository method to return exercise  from the database given the ID value
func (p *ExerciseRepository) FindByID(id uint) Exercise {
	var exercise Exercise
	p.DB.First(&exercise, id)

	return exercise
}

// Save is the repository method to save the exercise object into the database
func (p *ExerciseRepository) Save(exercise Exercise) Exercise {
	p.DB.Save(&exercise)

	return exercise
}

// Delete is the repository method to remove the exercise object from the database
func (p *ExerciseRepository) Delete(exercise Exercise) {
	p.DB.Delete(&exercise)
}
