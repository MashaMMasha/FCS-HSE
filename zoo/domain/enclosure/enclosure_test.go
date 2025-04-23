package enclosure_test

import (
	"testing"
	"zoo/domain/animal"
	"zoo/domain/enclosure"
)

func TestEnclosureCreation(t *testing.T) {
	e := enclosure.Enclosure{ID: 1, AnimalCount: 2, AnimalType: "predator", MaxCapacity: 3, AnimalIDs: make(map[int]struct{})}
	if e.ID != 1 {
		t.Errorf("Expected ID 1, got %d", e.ID)
	}
	if e.AnimalCount != 2 {
		t.Errorf("Expected AnimalCount 2, got %d", e.AnimalCount)
	}
	if e.AnimalType != "predator" {
		t.Errorf("Expected AnimalType 'predator', got %s", e.AnimalType)
	}
	if e.MaxCapacity != 3 {
		t.Errorf("Expected MaxCapacity 3, got %d", e.MaxCapacity)

	}
	if len(e.AnimalIDs) != 0 {
		t.Errorf("Expected AnimalIDs to be empty, got %v", e.AnimalIDs)
	}
}

func TestEnclosure_AddAnimal(t *testing.T) {
	e := enclosure.Enclosure{ID: 1, AnimalCount: 0, AnimalType: "predator", MaxCapacity: 3, AnimalIDs: make(map[int]struct{})}
	animalID := 1

	err := e.AddAnimal(&animal.Animal{ID: animalID, AnimalType: "predator"})
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if e.AnimalCount != 1 {
		t.Errorf("Expected AnimalCount 1, got %d", e.AnimalCount)
	}
	if _, ok := e.AnimalIDs[animalID]; !ok {
		t.Errorf("Expected AnimalIDs to contain %d", animalID)
	}
}

func TestEnclosure_RemoveAnimal(t *testing.T) {
	e := enclosure.Enclosure{ID: 1, AnimalCount: 1, AnimalType: "predator", MaxCapacity: 3, AnimalIDs: map[int]struct{}{1: {}}}
	animalID := 1

	err := e.RemoveAnimal(animalID)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if e.AnimalCount != 0 {
		t.Errorf("Expected AnimalCount 0, got %d", e.AnimalCount)
	}
	if _, ok := e.AnimalIDs[animalID]; ok {
		t.Errorf("Expected AnimalIDs to not contain %d", animalID)
	}
}
