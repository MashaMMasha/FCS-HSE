package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"zoo/application/services"
	"zoo/domain/animal"
)

type AnimalController struct {
	animalService *services.AnimalTransferService
}

func NewAnimalController(animalService *services.AnimalTransferService) *AnimalController {
	return &AnimalController{animalService: animalService}
}

type AddAnimalRequest struct {
	Name         string      `json:"name"`
	FavoriteFood string      `json:"favorite_food"`
	EnclosureID  int         `json:"enclosure_id"`
	AnimalType   animal.Type `json:"animal_type"`
}

// AddAnimal godoc
// @Summary Создание нового животного
// @Description Добавляет новое животное в вольер
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body AddAnimalRequest true "Новое животное"
// @Success 201 {string} string "ID нового животного"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /animal [post]
func (c *AnimalController) AddAnimal(w http.ResponseWriter, r *http.Request) {
	var req AddAnimalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	id, err := c.animalService.AddAnimal(req.EnclosureID, req.FavoriteFood, req.Name, req.AnimalType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))
}

type DeleteAnimalRequest struct {
	ID int `json:"id"`
}

// DeleteAnimal godoc
// @Summary Удаление животного
// @Description Удаляет животное из вольера
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body DeleteAnimalRequest true "ID животного"
// @Success 204 {string} string "Успешно удалено"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Животное не найдено"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /animal [delete]
func (c *AnimalController) DeleteAnimal(w http.ResponseWriter, r *http.Request) {
	var req DeleteAnimalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	err := c.animalService.DeleteAnimal(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

type GetAnimalRequest struct {
	ID int `json:"id"`
}

type GetAnimalResponse struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	FavoriteFood string      `json:"favorite_food"`
	EnclosureID  int         `json:"enclosure_id"`
	AnimalType   animal.Type `json:"animal_type"`
}

// GetAnimal godoc
// @Summary Получение информации о животном
// @Description Возвращает информацию о животном по ID
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body GetAnimalRequest true "ID животного"
// @Success 200 {object} GetAnimalResponse
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Животное не найдено"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /animal [get]
func (c *AnimalController) GetAnimal(w http.ResponseWriter, r *http.Request) {
	var req GetAnimalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	animal, err := c.animalService.GetAnimalByID(req.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetAnimalResponse{
		ID:           animal.ID,
		Name:         animal.Name,
		FavoriteFood: animal.FavoriteFood,
		EnclosureID:  animal.EnclosureID,
		AnimalType:   animal.AnimalType,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type GetAllAnimalsResponse struct {
	Animals []GetAnimalResponse `json:"animals"`
}

// GetAllAnimals godoc
// @Summary Получение всех животных
// @Description Возвращает список всех животных
// @Tags animals
// @Accept json
// @Produce json
// @Success 200 {object} GetAllAnimalsResponse
// @Failure 500 {string} string "Ошибка сервера"
// @Router /animals [get]
func (c *AnimalController) GetAllAnimals(w http.ResponseWriter, r *http.Request) {
	animals, err := c.animalService.GetAllAnimals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := GetAllAnimalsResponse{
		Animals: make([]GetAnimalResponse, len(animals)),
	}

	for i, animal := range animals {
		response.Animals[i] = GetAnimalResponse{
			ID:           animal.ID,
			Name:         animal.Name,
			FavoriteFood: animal.FavoriteFood,
			EnclosureID:  animal.EnclosureID,
			AnimalType:   animal.AnimalType,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

type ChangeEnclosureRequest struct {
	AnimalID    int `json:"animal_id"`
	EnclosureID int `json:"enclosure_id"`
}

// ChangeEnclosure godoc
// @Summary Перемещение животного в другой вольер
// @Description Перемещает животное в другой вольер
// @Tags animals
// @Accept json
// @Produce json
// @Param animal body ChangeEnclosureRequest true "ID животного и ID нового вольера"
// @Success 200 {string} string "Успешно перемещено"
// @Failure 400 {string} string "Неверный формат запроса"
// @Failure 404 {string} string "Животное или вольер не найдены"
// @Failure 500 {string} string "Ошибка сервера"
// @Router /animal/move [post]
func (c *AnimalController) ChangeEnclosure(w http.ResponseWriter, r *http.Request) {
	var req ChangeEnclosureRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	_, err := c.animalService.TransferAnimal(req.AnimalID, req.EnclosureID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Animal moved successfully"))
}
