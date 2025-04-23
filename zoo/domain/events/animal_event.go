package events

import "time"

type AnimalMovedEvent struct {
	AnimalID       int
	OldEnclosureID int
	NewEnclosureID int
	Timestamp      time.Time
}

func NewAnimalMovedEvent(animalID, oldEnclosureID, newEnclosureID int) AnimalMovedEvent {
	return AnimalMovedEvent{
		AnimalID:       animalID,
		OldEnclosureID: oldEnclosureID,
		NewEnclosureID: newEnclosureID,
		Timestamp:      time.Now(),
	}
}
