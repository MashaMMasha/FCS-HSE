package ports

import (
	"zoo/domain/enclosure"
)

type EnclosureRepository interface {
	GetByID(id int) (*enclosure.Enclosure, error)
	Save(a *enclosure.Enclosure) error
	Delete(id int) error
	GetAll() []*enclosure.Enclosure
}
