package ports

import "zoo/domain/events"

type EventHandler interface {
	HandleAnimalMoved(event events.AnimalMovedEvent)
	HandleFeedingTime(event events.FeedingTimeEvent)
}
