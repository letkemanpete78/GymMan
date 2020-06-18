package exercise

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	//c.JSON(http.StatusOK, gin.H{"exercises": ToDTOs(exercises)})
	c.JSON(http.StatusOK, gin.H{"exercises": exercises})
}

// FindByUUID finds the exercise record by primary key/id
func (p *API) FindByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	exercise := p.Service.FindByUUID(uuid)

	// c.JSON(http.StatusOK, gin.H{"exercise": ToDTO(exercise)})
	c.JSON(http.StatusOK, gin.H{"exercise": exercise})

}

// Create inserts a record into database
func (p *API) Create(c *gin.Context) {
	var exerciseDTO DTO
	exerciseDTO.Description = c.PostForm("description")
	exerciseDTO.Name = c.PostForm("name")
	exerciseDTO.UUID = uuid.New().String()

	// err := c.BindJSON(&exerciseDTO)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	c.Status(http.StatusBadRequest)
	// 	return
	// }

	createdExercise := p.Service.Save(ToExercise(exerciseDTO))
	// c.JSON(http.StatusOK, gin.H{"exercise": ToDTO(createdExercise)})
	c.JSON(http.StatusOK, gin.H{"exercise": createdExercise})
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

	uuid := c.Param("uuid")
	exercise := p.Service.FindByUUID(uuid)
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
	uuid := c.Param("uuid")
	exercise := p.Service.FindByUUID(uuid)
	if exercise == (Exercise{}) {
		c.Status(http.StatusBadRequest)
		return
	}

	p.Service.Delete(exercise)

	c.Status(http.StatusOK)
}
