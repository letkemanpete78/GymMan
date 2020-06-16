package exercise

// DTO used for json data transfer objects
type DTO struct {
	UUID        string `json:"uuid,strings"`
	Name        string `json:"name,string"`
	Description string `json:"description,string"`
}
