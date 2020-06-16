package exercise

import (
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Repository gorm db object
type Repository struct {
	DB *gorm.DB
}

// ProvideExerciseRepostiory is the provider repostiory object
func ProvideExerciseRepostiory(DB *gorm.DB) Repository {
	return Repository{DB: DB}
}

// FindAll is the repository method to returns all exercise objects from the database
func (p *Repository) FindAll() []Exercise {
	var exercises []Exercise
	p.DB.Find(&exercises)

	return exercises
}

// FindByUUID is the repository method to return exercise  from the database given the ID value
func (p *Repository) FindByUUID(uuid string) Exercise {
	var exercise Exercise
	p.DB.First(&exercise, uuid)

	return exercise
}

// Save is the repository method to save the exercise object into the database
func (p *Repository) Save(exercise Exercise) Exercise {
	if exercise.UUID == "" {
		exercise.UUID = uuid.New().String()
	}
	p.DB.Save(&exercise)

	return exercise
}

// Delete is the repository method to remove the exercise object from the database
func (p *Repository) Delete(exercise Exercise) {
	p.DB.Delete(&exercise)
}
