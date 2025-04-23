package http

import (
	"encoding/json"
	"net/http"
	"zoo/application/services"
)

type ScheduleController struct {
	feedingService *services.FeedingService
}

func NewScheduleController(feedingService *services.FeedingService) *ScheduleController {
	return &ScheduleController{feedingService: feedingService}
}

type AddFeedingScheduleRequest struct {
	AnimalID        int    `json:"animal_id"`
	FoodType        string `json:"food_type"`
	FeedingInterval int    `json:"feeding_interval"`
}

type ChangeFeedingIntervalRequest struct {
	ID              int `json:"schedule_id"`
	FeedingInterval int `json:"feeding_interval"`
}

// AddFeedingSchedule godoc
// @Summary Создание нового расписания кормления
// @Description Добавляет новое расписание кормления для животного
// @Tags feeding
// @Accept json
// @Produce json
// @Param schedule body AddFeedingScheduleRequest true "Новое расписание кормления"
// @Success 201 {string} string "ID нового расписания"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /feeding/schedule [post]
func (c *ScheduleController) AddFeedingSchedule(w http.ResponseWriter, r *http.Request) {
	var req AddFeedingScheduleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := c.feedingService.AddFeedingSchedule(req.AnimalID, req.FoodType, req.FeedingInterval)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// ChangeFeedInterval godoc
// @Summary Изменение интервала кормления
// @Description Изменяет интервал кормления для расписания
// @Tags feeding
// @Accept json
// @Produce json
// @Param schedule body ChangeFeedingIntervalRequest true "Изменение интервала кормления"
// @Success 200 {string} string "Интервал кормления изменен"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /feeding/schedule/interval [put]
func (c *ScheduleController) ChangeFeedInterval(w http.ResponseWriter, r *http.Request) {
	var req ChangeFeedingIntervalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := c.feedingService.ChangeFeedInterval(req.ID, req.FeedingInterval)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type FeedAnimalRequest struct {
	AnimalID int    `json:"animal_id"`
	Food     string `json:"food"`
}

// FeedAnimal godoc
// @Summary Кормление животного
// @Description Кормит животное по расписанию
// @Tags feeding
// @Accept json
// @Produce json
// @Param animal body FeedAnimalRequest true "ID животного"
// @Success 200 {string} string "Животное покормлено"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Животное не найдено"
// @Failure 409 {string} string "Животное не голодно"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /feeding [post]
func (c *ScheduleController) FeedAnimal(w http.ResponseWriter, r *http.Request) {
	var req FeedAnimalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := c.feedingService.FeedAnimal(req.AnimalID, req.Food)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
