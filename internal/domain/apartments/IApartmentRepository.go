package apartments

type IApartmentRepository interface {
	GetById(id string) *Apartment
}
