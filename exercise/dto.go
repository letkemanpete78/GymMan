package exercise

// DTO used for json data transfer objects
type DTO struct {
	ID          uint   `json:"id,string,omitempty"`
	Name        string `json:"name,string"`
	Description string `json:"description,string"`
}
