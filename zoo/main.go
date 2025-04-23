package main

import (
	"net/http"
	"zoo/infrastructure/eventhandlers"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "zoo/docs"

	"zoo/application/services"
	"zoo/infrastructure/repositories"
	controllers "zoo/presentation/http"
)

func main() {
	animalRepo := repositories.NewInMemoryAnimalRepository()
	enclosureRepo := repositories.NewInMemoryEnclosureRepository()
	feedingRepo := repositories.NewInMemoryScheduleRepository()
	eventHandler := eventhandlers.NewLoggingEventHandler()
	animalService := services.NewAnimalTransferService(animalRepo, enclosureRepo, eventHandler)
	feedingService := services.NewFeedingService(animalRepo, feedingRepo, eventHandler)
	animalController := controllers.NewAnimalController(animalService)
	enclosureController := controllers.NewEnclosureController(animalService)
	feedingController := controllers.NewScheduleController(feedingService)

	mux := http.NewServeMux()
	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	mux.HandleFunc("/animal", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			animalController.AddAnimal(w, r)
		case http.MethodDelete:
			animalController.DeleteAnimal(w, r)
		case http.MethodGet:
			animalController.GetAnimal(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/animal/move", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			animalController.ChangeEnclosure(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	mux.HandleFunc("/animals", animalController.GetAllAnimals)
	mux.HandleFunc("/enclosures", enclosureController.GetAllEnclosures)
	mux.HandleFunc("/enclosure", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			enclosureController.AddEnclosure(w, r)
		case http.MethodDelete:
			enclosureController.DeleteEnclosure(w, r)
		case http.MethodGet:
			enclosureController.GetEnclosure(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/feeding/schedule", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			feedingController.AddFeedingSchedule(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/feeding", feedingController.FeedAnimal)
	mux.HandleFunc("feeding/schedule/interval", feedingController.ChangeFeedInterval)

	http.ListenAndServe(":8080", mux)
}
