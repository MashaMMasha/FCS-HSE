package ports

import "zoo/domain/animal"

type AnimalRepository interface {
	GetByID(id int) (*animal.Animal, error)
	Save(a *animal.Animal) error
	Delete(id int) error
	GetAll() ([]*animal.Animal, error)
}
