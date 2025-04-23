package http

import (
	"encoding/json"
	"net/http"
	"zoo/application/services"
	"zoo/domain/animal"
)

type EnclosureController struct {
	enclosureService *services.AnimalTransferService
}

func NewEnclosureController(enclosureService *services.AnimalTransferService) *EnclosureController {
	return &EnclosureController{enclosureService: enclosureService}
}

type AddEnclosureRequest struct {
	// Тип животного
	// required: true
	// enum: predator,herbivore,bird,aquarium
	// example: aquarium
	AnimalType animal.Type `json:"animal_type"`

	// Вместимость вольера
	// required: true
	// minimum: 1
	// example: 10
	Capacity int `json:"capacity"`
}

type DeleteEnclosureRequest struct {
	// ID вольера
	// required: true
	// example: 1
	ID int `json:"id"`
}

// AddEnclosure godoc
// @Summary Создание нового вольера
// @Description Добавляет новый вольер
// @Tags enclosures
// @Accept json
// @Produce json
// @Param enclosure body AddEnclosureRequest true "Данные нового вольера"
// @Success 201 {string} string "ID нового вольера"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /enclosure [post]
func (c *EnclosureController) AddEnclosure(w http.ResponseWriter, r *http.Request) {
	var req AddEnclosureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Валидация
	if req.Capacity <= 0 {
		http.Error(w, "capacity must be positive", http.StatusBadRequest)
		return
	}

	id, err := c.enclosureService.AddEnclosure(req.AnimalType, req.Capacity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

// DeleteEnclosure godoc
// @Summary Удаление вольера
// @Description Удаляет вольер по ID
// @Tags enclosures
// @Accept json
// @Produce json
// @Param id body DeleteEnclosureRequest true "ID вольера"
// @Success 204 "Вольер удален"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /enclosure [delete]
func (c *EnclosureController) DeleteEnclosure(w http.ResponseWriter, r *http.Request) {
	var req DeleteEnclosureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := c.enclosureService.DeleteEnclosure(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

type GetAllEnclosuresResponse struct {
	ID          int         `json:"id"`
	AnimalCount int         `json:"animal_count"`
	AnimalType  animal.Type `json:"animal_type"`
	MaxCapacity int         `json:"max_capacity"`
	AnimalIDs   []int       `json:"animal_ids"`
}

// GetAllEnclosures godoc
// @Summary Получение всех вольеров
// @Description Возвращает список всех вольеров
// @Tags enclosures
// @Accept json
// @Produce json
// @Success 200 {array} enclosure.Enclosure "Список вольеров"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /enclosures [get]
func (c *EnclosureController) GetAllEnclosures(w http.ResponseWriter, r *http.Request) {
	enclosures, err := c.enclosureService.GetAllEnclosures()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := make([]GetAllEnclosuresResponse, len(enclosures))
	for i, enclosure := range enclosures {
		ids := make([]int, 0, len(enclosure.AnimalIDs))
		for id := range enclosure.AnimalIDs {
			ids = append(ids, id)
		}
		response[i] = GetAllEnclosuresResponse{
			ID:          enclosure.ID,
			AnimalCount: enclosure.AnimalCount,
			AnimalType:  enclosure.AnimalType,
			MaxCapacity: enclosure.MaxCapacity,
			AnimalIDs:   ids,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type GetEnclosureRequest struct {
	ID int `json:"id"`
}

// GetEnclosure godoc
// @Summary Получение вольера по ID
// @Description Возвращает вольер по ID
// @Tags enclosures
// @Accept json
// @Produce json
// @Param id path int true "ID вольера"
// @Success 200 {object} GetAllEnclosuresResponse
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Вольер не найден"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /enclosure [get]
func (c *EnclosureController) GetEnclosure(w http.ResponseWriter, r *http.Request) {
	var req GetEnclosureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	enclosure, err := c.enclosureService.GetEnclosureByID(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetAllEnclosuresResponse{
		ID:          enclosure.ID,
		AnimalCount: enclosure.AnimalCount,
		AnimalType:  enclosure.AnimalType,
		MaxCapacity: enclosure.MaxCapacity,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
