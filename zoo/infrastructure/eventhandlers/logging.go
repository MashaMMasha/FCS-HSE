package eventhandlers

import (
	"log"
	"zoo/application/ports"
	"zoo/domain/events"
)

type LoggingEventHandler struct{}

func NewLoggingEventHandler() ports.EventHandler {
	return &LoggingEventHandler{}
}

func (h *LoggingEventHandler) HandleAnimalMoved(event events.AnimalMovedEvent) {
	log.Printf("Животное с ID=%d перемещено из вольера %d в %d", event.AnimalID, event.OldEnclosureID, event.NewEnclosureID)
}

func (h *LoggingEventHandler) HandleFeedingTime(event events.FeedingTimeEvent) {
	log.Printf("Животное покормлено в %d (съел %s)", event.AnimalID, event.FoodType)
}
