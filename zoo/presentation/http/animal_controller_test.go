package http_test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"zoo/application/services"
	"zoo/domain/enclosure"
	"zoo/infrastructure/repositories"
	controller "zoo/presentation/http"
)

func TestAnimalController2(t *testing.T) {
	animalRepo := repositories.NewInMemoryAnimalRepository()
	enclosureRepo := repositories.NewInMemoryEnclosureRepository()
	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)
	controller := controller.NewAnimalController(service)

	enclosure1 := &enclosure.Enclosure{ID: 1, MaxCapacity: 2, AnimalType: "predator", AnimalIDs: make(map[int]struct{})}
	enclosure2 := &enclosure.Enclosure{ID: 2, MaxCapacity: 2, AnimalType: "predator", AnimalIDs: make(map[int]struct{})}
	enclosureRepo.Save(enclosure1)
	enclosureRepo.Save(enclosure2)

	t.Run("AddAnimal success", func(t *testing.T) {
		reqBody := `{"name":"Lion","favorite_food":"Meat","enclosure_id":1,"animal_type":"predator"}`
		req := httptest.NewRequest(http.MethodPost, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.AddAnimal(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "1", rec.Body.String())
	})

	t.Run("Add second animal", func(t *testing.T) {
		reqBody := `{"name":"Tiger","favorite_food":"Meat","enclosure_id":1,"animal_type":"predator"}`
		req := httptest.NewRequest(http.MethodPost, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.AddAnimal(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, "2", rec.Body.String())
	})

	t.Run("GetAllAnimals", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/animals", nil)
		rec := httptest.NewRecorder()

		controller.GetAllAnimals(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.NewDecoder(rec.Body).Decode(&response)
		assert.NoError(t, err)

		animals := response["animals"].([]interface{})
		assert.Equal(t, 2, len(animals))

		firstAnimal := animals[0].(map[string]interface{})
		assert.Equal(t, float64(1), firstAnimal["id"])
		assert.Equal(t, "Lion", firstAnimal["name"])

		secondAnimal := animals[1].(map[string]interface{})
		assert.Equal(t, float64(2), secondAnimal["id"])
		assert.Equal(t, "Tiger", secondAnimal["name"])
	})

	t.Run("ChangeEnclosure invalid enclosure", func(t *testing.T) {
		reqBody := `{"animal_id":1,"enclosure_id":999}`
		req := httptest.NewRequest(http.MethodPost, "/animal/move", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.ChangeEnclosure(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("ChangeEnclosure invalid animal", func(t *testing.T) {
		reqBody := `{"animal_id":999,"enclosure_id":1}`
		req := httptest.NewRequest(http.MethodPost, "/animal/move", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.ChangeEnclosure(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("GetAnimal success", func(t *testing.T) {
		reqBody := `{"id":1}`
		req := httptest.NewRequest(http.MethodGet, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.GetAnimal(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var response map[string]interface{}
		err := json.NewDecoder(rec.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, float64(1), response["id"])
		assert.Equal(t, "Lion", response["name"])
		assert.Equal(t, "Meat", response["favorite_food"])
		assert.Equal(t, float64(1), response["enclosure_id"])
		assert.Equal(t, "predator", response["animal_type"])
	})

	t.Run("GetAnimal not found", func(t *testing.T) {
		reqBody := `{"id":999}`
		req := httptest.NewRequest(http.MethodGet, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.GetAnimal(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("DeleteAnimal success", func(t *testing.T) {
		reqBodyAdd := `{"name":"Tiger","favorite_food":"Meat","enclosure_id":1,"animal_type":"predator"}`
		reqAdd := httptest.NewRequest(http.MethodPost, "/animal", bytes.NewBufferString(reqBodyAdd))
		reqAdd.Header.Set("Content-Type", "application/json")
		recAdd := httptest.NewRecorder()
		controller.AddAnimal(recAdd, reqAdd)

		reqBody := `{"id":2}`
		req := httptest.NewRequest(http.MethodDelete, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.DeleteAnimal(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		_, err := animalRepo.GetByID(2)
		assert.Error(t, err)
	})

	t.Run("DeleteAnimal not found", func(t *testing.T) {
		reqBody := `{"id":999}`
		req := httptest.NewRequest(http.MethodDelete, "/animal", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		controller.DeleteAnimal(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

}
