package ports

import (
	"zoo/domain/schedule"
)

type ScheduleRepository interface {
	GetByID(id int) (*schedule.FeedingSchedule, error)
	GetByAnimalID(animalID int) (*schedule.FeedingSchedule, error)
	Save(schedule *schedule.FeedingSchedule) error
	Delete(id int) error
}
