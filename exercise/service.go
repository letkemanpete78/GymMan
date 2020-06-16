package exercise

// Service is the object used for the repository
type Service struct {
	Repository Repository
}

// ProvideService is the provider used for the repoitory service
func ProvideService(p Repository) Service {
	return Service{Repository: p}
}

// FindAll is the service method to returns all exercise objects from the database
func (p *Service) FindAll() []Exercise {
	return p.Repository.FindAll()
}

// FindByUUID is the service method to return exercise  from the database given the ID value
func (p *Service) FindByUUID(uuid string) Exercise {
	return p.Repository.FindByUUID(uuid)
}

// Save is the service method to save the exercise object into the database
func (p *Service) Save(exercise Exercise) Exercise {
	p.Repository.Save(exercise)

	return exercise
}

// Delete is the service method to remove the exercise object from the database
func (p *Service) Delete(exercise Exercise) {
	p.Repository.Delete(exercise)
}
