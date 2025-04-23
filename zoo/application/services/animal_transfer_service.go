package services

import (
	"fmt"
	"time"
	"zoo/application/ports"
	"zoo/domain/animal"
	"zoo/domain/enclosure"
	"zoo/domain/events"
)

type AnimalTransferService struct {
	animalRepo    ports.AnimalRepository
	enclosureRepo ports.EnclosureRepository
	eventHandler  ports.EventHandler
}

func NewAnimalTransferService(
	animalRepo ports.AnimalRepository,
	enclosureRepo ports.EnclosureRepository,
	eventHandler ports.EventHandler,
) *AnimalTransferService {
	return &AnimalTransferService{
		animalRepo:    animalRepo,
		enclosureRepo: enclosureRepo,
		eventHandler:  eventHandler,
	}
}

func (s *AnimalTransferService) TransferAnimal(animalID int, toEnclosureID int) (*events.AnimalMovedEvent, error) {
	animal, err := s.animalRepo.GetByID(animalID)
	if err != nil {
		return nil, fmt.Errorf("animal not found: %w", err)
	}

	fromEnclosure, err := s.enclosureRepo.GetByID(animal.EnclosureID)
	if err != nil {
		return nil, fmt.Errorf("first enclosure not found: %w", err)
	}

	toEnclosure, err := s.enclosureRepo.GetByID(toEnclosureID)
	if err != nil {
		return nil, fmt.Errorf("second enclosure not found: %w", err)
	}

	if err = fromEnclosure.RemoveAnimal(animal.ID); err != nil {
		return nil, fmt.Errorf("failed to remove from old enclosure: %w", err)
	}
	if err = toEnclosure.AddAnimal(animal); err != nil {
		_ = fromEnclosure.AddAnimal(animal)
		return nil, fmt.Errorf("failed to add to new enclosure: %w", err)
	}

	animal.ChangeEnclosure(toEnclosureID)

	if err = s.animalRepo.Save(animal); err != nil {
		return nil, fmt.Errorf("failed to update animal: %w", err)
	}
	if err = s.enclosureRepo.Save(fromEnclosure); err != nil {
		return nil, err
	}
	if err = s.enclosureRepo.Save(toEnclosure); err != nil {
		return nil, err
	}
	event := events.AnimalMovedEvent{
		AnimalID:       animal.ID,
		OldEnclosureID: fromEnclosure.ID,
		NewEnclosureID: toEnclosure.ID,
		Timestamp:      time.Now(),
	}
	s.eventHandler.HandleAnimalMoved(event)

	return &event, nil
}

func (s *AnimalTransferService) AddAnimal(enclosureID int, food string, name string, animalType animal.Type) (int, error) {
	enclosure, err := s.enclosureRepo.GetByID(enclosureID)
	if err != nil {
		return 0, fmt.Errorf("enclosure not found: %w", err)
	}

	newAnimal := &animal.Animal{
		Name:         name,
		BirthDate:    time.Now(),
		FavoriteFood: food,
		Status:       animal.Healthy,
		EnclosureID:  enclosureID,
		AnimalType:   animalType,
	}
	
	if err = s.animalRepo.Save(newAnimal); err != nil {
		return 0, fmt.Errorf("failed to save new animal: %w", err)
	}

	if err = enclosure.AddAnimal(newAnimal); err != nil {
		return 0, fmt.Errorf("cannot add animal to enclosure: %w", err)
	}

	if err = s.animalRepo.Save(newAnimal); err != nil {
		return 0, fmt.Errorf("failed to save new animal: %w", err)
	}

	if err = s.enclosureRepo.Save(enclosure); err != nil {
		return 0, fmt.Errorf("failed to update enclosure: %w", err)
	}

	return newAnimal.ID, nil
}

func (s *AnimalTransferService) DeleteAnimal(animalID int) error {
	animal, err := s.animalRepo.GetByID(animalID)
	if err != nil {
		return fmt.Errorf("animal not found: %w", err)
	}

	enclosure, err := s.enclosureRepo.GetByID(animal.EnclosureID)
	if err != nil {
		return fmt.Errorf("enclosure not found for animal %d: %w", animal.ID, err)
	}

	if err = enclosure.RemoveAnimal(animal.ID); err != nil {
		return fmt.Errorf("failed to remove animal from enclosure: %w", err)
	}

	if err = s.enclosureRepo.Save(enclosure); err != nil {
		return fmt.Errorf("failed to update enclosure: %w", err)
	}

	if err = s.animalRepo.Delete(animalID); err != nil {
		return fmt.Errorf("failed to delete animal: %w", err)
	}

	return nil
}

func (s *AnimalTransferService) AddEnclosure(animalType animal.Type, maxCap int) (int, error) {
	enclosure := &enclosure.Enclosure{
		ID:          0,
		AnimalCount: 0,
		AnimalType:  animalType,
		MaxCapacity: maxCap,
		AnimalIDs:   make(map[int]struct{}),
	}

	if err := s.enclosureRepo.Save(enclosure); err != nil {
		return 0, fmt.Errorf("failed to save enclosure: %w", err)
	}

	return enclosure.ID, nil
}

func (s *AnimalTransferService) DeleteEnclosure(enclosureID int) error {
	enclosure, err := s.enclosureRepo.GetByID(enclosureID)
	if err != nil {
		return fmt.Errorf("enclosure not found: %w", err)
	}

	if enclosure.AnimalCount > 0 {
		return fmt.Errorf("enclosure is not empty, transfer animals first")
	}

	if err = s.enclosureRepo.Delete(enclosureID); err != nil {
		return fmt.Errorf("failed to delete enclosure: %w", err)
	}

	return nil
}

func (s *AnimalTransferService) GetAnimalByID(animalID int) (*animal.Animal, error) {
	animal, err := s.animalRepo.GetByID(animalID)
	if err != nil {
		return nil, fmt.Errorf("animal not found: %w", err)
	}
	return animal, nil
}

func (s *AnimalTransferService) GetAllAnimals() ([]*animal.Animal, error) {
	animals, err := s.animalRepo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get animals: %w", err)
	}
	return animals, nil
}

func (s *AnimalTransferService) GetAllEnclosures() ([]*enclosure.Enclosure, error) {
	enclosures := s.enclosureRepo.GetAll()
	return enclosures, nil
}

func (s *AnimalTransferService) GetEnclosureByID(enclosureID int) (*enclosure.Enclosure, error) {
	enclosure, err := s.enclosureRepo.GetByID(enclosureID)
	if err != nil {
		return nil, fmt.Errorf("enclosure not found: %w", err)
	}
	return enclosure, nil
}
