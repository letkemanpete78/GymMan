package exercise

// ExerciseDTO used for json data transfer objects
type ExerciseDTO struct {
	ID          uint   `json:"id,string,omitempty"`
	Name        string `json:"string"`
	Description string `json:"description,string"`
	Image       string `json:"string"`
}
