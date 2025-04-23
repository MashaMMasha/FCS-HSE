package services

import "zoo/application/ports"

type ZooStatisticsService struct {
	animalRepo    ports.AnimalRepository
	enclosureRepo ports.EnclosureRepository
}

func NewZooStatisticsService(
	animalRepo ports.AnimalRepository,
	enclosureRepo ports.EnclosureRepository,
) *ZooStatisticsService {
	return &ZooStatisticsService{
		animalRepo:    animalRepo,
		enclosureRepo: enclosureRepo,
	}
}
func (s *ZooStatisticsService) GetAnimalCount() (int, error) {
	animals, err := s.animalRepo.GetAll()
	if err != nil {
		return 0, err
	}
	return len(animals), nil
}
func (s *ZooStatisticsService) GetEnclosureCount() (int, error) {
	enclosures := s.enclosureRepo.GetAll()
	return len(enclosures), nil
}
