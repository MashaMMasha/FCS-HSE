package animal

import (
	"fmt"
	"time"
)

type Type string

const (
	Predator  Type = "predator"
	Herbivore Type = "herbivore"
	Bird      Type = "bird"
	Aquarium  Type = "aquarium"
)

func (t Type) String() string {
	return string(t)
}

type Status string

const (
	Healthy Status = "healthy"
	Sick    Status = "sick"
)

type Animal struct {
	ID                int
	Name              string
	BirthDate         time.Time
	FavoriteFood      string
	FeedingScheduleID int
	Status            Status
	EnclosureID       int
	AnimalType        Type
}

func (a *Animal) Feed(food string) error {
	if food == a.FavoriteFood {
		return nil
	}
	return fmt.Errorf("incorrect food for %s", a.Name)
}
func (a *Animal) Heal() {
	a.Status = Healthy
}
func (a *Animal) ChangeEnclosure(enclosureID int) {
	a.EnclosureID = enclosureID
}
