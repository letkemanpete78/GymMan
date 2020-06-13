package exercise

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ExerciseAPI defines the model to the exercise service
type ExerciseAPI struct {
	ExerciseService ExerciseService
}

// ProvideExerciseAPI defines the model for the exercise provider
func ProvideExerciseAPI(p ExerciseService) ExerciseAPI {
	return ExerciseAPI{ExerciseService: p}
}

// FindAll returns all exercise records
func (p *ExerciseAPI) FindAll(c *gin.Context) {
	exercises := p.ExerciseService.FindAll()

	c.JSON(http.StatusOK, gin.H{"exercises": ToExerciseDTOs(exercises)})
}

// FindByID finds the exercise record by primary key/id
func (p *ExerciseAPI) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.ExerciseService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"exercise": ToExerciseDTO(exercise)})
}

// Create inserts a record into database
func (p *ExerciseAPI) Create(c *gin.Context) {
	var exerciseDTO ExerciseDTO
	err := c.BindJSON(&exerciseDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdExercise := p.ExerciseService.Save(ToExercise(exerciseDTO))

	c.JSON(http.StatusOK, gin.H{"exercise": ToExerciseDTO(createdExercise)})
}

// Update saves the updated record to the database
func (p *ExerciseAPI) Update(c *gin.Context) {
	var exerciseDTO ExerciseDTO
	err := c.BindJSON(&exerciseDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.ExerciseService.FindByID(uint(id))
	if exercise == (Exercise{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	exercise.Name = exerciseDTO.Name
	exercise.Description = exerciseDTO.Description
	exercise.Image = exerciseDTO.Image
	p.ExerciseService.Save(exercise)

	c.Status(http.StatusOK)
}

// Delete removes an exercise record from database
func (p *ExerciseAPI) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.ExerciseService.FindByID(uint(id))
	if exercise == (Exercise{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.ExerciseService.Delete(exercise)

	c.Status(http.StatusOK)
}
