package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"zoo/application/services"
	enclosure "zoo/domain/enclosure"
	"zoo/infrastructure/repositories"
	controller "zoo/presentation/http"

	"github.com/stretchr/testify/assert"
)

func TestEnclosureController(t *testing.T) {
	animalRepo := repositories.NewInMemoryAnimalRepository()
	enclosureRepo := repositories.NewInMemoryEnclosureRepository()
	service := services.NewAnimalTransferService(animalRepo, enclosureRepo, nil)
	enclosureController := controller.NewEnclosureController(service)

	t.Run("AddEnclosure success", func(t *testing.T) {
		reqBody := `{"animal_type":"predator","capacity":5}`
		req := httptest.NewRequest(http.MethodPost, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.AddEnclosure(rec, req)

		assert.Equal(t, http.StatusCreated, rec.Code)

		var response map[string]int
		err := json.NewDecoder(rec.Body).Decode(&response)
		assert.NoError(t, err)
		assert.Equal(t, 1, response["id"])
	})

	t.Run("AddEnclosure invalid capacity", func(t *testing.T) {
		reqBody := `{"animal_type":"predator","capacity":0}`
		req := httptest.NewRequest(http.MethodPost, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.AddEnclosure(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), "capacity must be positive")
	})

	t.Run("GetEnclosure success", func(t *testing.T) {
		enc := &enclosure.Enclosure{ID: 2, AnimalType: "herbivore", MaxCapacity: 3}
		enclosureRepo.Save(enc)

		reqBody := `{"id":2}`
		req := httptest.NewRequest(http.MethodGet, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.GetEnclosure(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var response controller.GetAllEnclosuresResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		assert.NoError(t, err)

		assert.Equal(t, 2, response.ID)
		assert.Equal(t, "herbivore", string(response.AnimalType))
		assert.Equal(t, 3, response.MaxCapacity)
	})

	t.Run("GetEnclosure not found", func(t *testing.T) {
		reqBody := `{"id":999}`
		req := httptest.NewRequest(http.MethodGet, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.GetEnclosure(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})

	t.Run("GetAllEnclosures", func(t *testing.T) {
		enc := &enclosure.Enclosure{ID: 3, AnimalType: "bird", MaxCapacity: 10}
		enclosureRepo.Save(enc)

		req := httptest.NewRequest(http.MethodGet, "/enclosures", nil)
		rec := httptest.NewRecorder()

		enclosureController.GetAllEnclosures(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var response []controller.GetAllEnclosuresResponse
		err := json.NewDecoder(rec.Body).Decode(&response)
		assert.NoError(t, err)

		assert.GreaterOrEqual(t, len(response), 2)

		found := false
		for _, enc := range response {
			if enc.ID == 3 {
				found = true
				assert.Equal(t, "bird", string(enc.AnimalType))
				assert.Equal(t, 10, enc.MaxCapacity)
				break
			}
		}
		assert.True(t, found)
	})

	t.Run("DeleteEnclosure success", func(t *testing.T) {
		enc := &enclosure.Enclosure{ID: 4, AnimalType: "aquarium", MaxCapacity: 20}
		enclosureRepo.Save(enc)

		reqBody := `{"id":4}`
		req := httptest.NewRequest(http.MethodDelete, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.DeleteEnclosure(rec, req)

		assert.Equal(t, http.StatusNoContent, rec.Code)

		_, err := enclosureRepo.GetByID(4)
		assert.Error(t, err)
	})

	t.Run("DeleteEnclosure not found", func(t *testing.T) {
		reqBody := `{"id":999}`
		req := httptest.NewRequest(http.MethodDelete, "/enclosure", bytes.NewBufferString(reqBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()

		enclosureController.DeleteEnclosure(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)
	})
}
