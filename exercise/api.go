package exercise

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// API defines the model to the exercise service
type API struct {
	Service Service
}

// ProvideAPI defines the model for the exercise provider
func ProvideAPI(p Service) API {
	return API{Service: p}
}

// FindAll returns all exercise records
func (p *API) FindAll(c *gin.Context) {
	exercises := p.Service.FindAll()

	c.JSON(http.StatusOK, gin.H{"exercises": ToDTOs(exercises)})
}

// FindByID finds the exercise record by primary key/id
func (p *API) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.Service.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"exercise": ToDTO(exercise)})
}

// Create inserts a record into database
func (p *API) Create(c *gin.Context) {
	var exerciseDTO DTO
	err := c.BindJSON(&exerciseDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	createdExercise := p.Service.Save(ToExercise(exerciseDTO))

	c.JSON(http.StatusOK, gin.H{"exercise": ToDTO(createdExercise)})
}

// Update saves the updated record to the database
func (p *API) Update(c *gin.Context) {
	var exerciseDTO DTO
	err := c.BindJSON(&exerciseDTO)
	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.Service.FindByID(uint(id))
	if exercise == (Exercise{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	exercise.Name = exerciseDTO.Name
	exercise.Description = exerciseDTO.Description
	p.Service.Save(exercise)

	c.Status(http.StatusOK)
}

// Delete removes an exercise record from database
func (p *API) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	exercise := p.Service.FindByID(uint(id))
	if exercise == (Exercise{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Service.Delete(exercise)

	c.Status(http.StatusOK)
}
