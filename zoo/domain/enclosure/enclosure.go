package enclosure

import (
	"fmt"
	"zoo/domain/animal"
)

type Enclosure struct {
	ID          int
	AnimalCount int
	AnimalType  animal.Type
	MaxCapacity int
	AnimalIDs   map[int]struct{}
}

func (e *Enclosure) AddAnimal(a *animal.Animal) error {
	if e.AnimalType != a.AnimalType {
		return fmt.Errorf("this enclosure is only for %s", e.AnimalType)
	}
	if e.MaxCapacity <= e.AnimalCount {
		return fmt.Errorf("no space left in enclosure %d", e.ID)
	}
	e.AnimalCount++
	e.AnimalIDs[a.ID] = struct{}{}
	return nil
}

func (e *Enclosure) RemoveAnimal(animalID int) error {
	if _, ok := e.AnimalIDs[animalID]; !ok {
		return fmt.Errorf("animal %d not found in enclosure %d", animalID, e.ID)
	}
	delete(e.AnimalIDs, animalID)
	e.AnimalCount--
	return nil
}
