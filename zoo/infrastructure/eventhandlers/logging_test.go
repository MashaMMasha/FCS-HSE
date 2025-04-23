package eventhandlers_test

import (
	"bytes"
	"log"
	"testing"
	"zoo/domain/events"
	"zoo/infrastructure/eventhandlers"

	"github.com/stretchr/testify/assert"
)

func TestLoggingEventHandler_HandleAnimalMoved(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)
	defer log.SetOutput(nil)

	handler := eventhandlers.NewLoggingEventHandler()

	event := events.AnimalMovedEvent{
		AnimalID:       1,
		OldEnclosureID: 101,
		NewEnclosureID: 102,
	}

	handler.HandleAnimalMoved(event)

	expectedLog := "Животное с ID=1 перемещено из вольера 101 в 102"
	assert.Contains(t, logBuffer.String(), expectedLog)
}

func TestLoggingEventHandler_HandleFeedingTime(t *testing.T) {
	var logBuffer bytes.Buffer
	log.SetOutput(&logBuffer)
	defer log.SetOutput(nil)

	handler := eventhandlers.NewLoggingEventHandler()

	event := events.FeedingTimeEvent{
		AnimalID: 1,
		FoodType: "Meat",
	}

	handler.HandleFeedingTime(event)

	expectedLog := "Животное покормлено в 1 (съел Meat)"
	assert.Contains(t, logBuffer.String(), expectedLog)
}
