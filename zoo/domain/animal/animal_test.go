package animal_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zoo/domain/animal"
)

func TestAnimal_Feed(t *testing.T) {
	a := &animal.Animal{
		Name:         "Leo",
		FavoriteFood: "Meat",
	}

	t.Run("correct food", func(t *testing.T) {
		err := a.Feed("Meat")
		assert.NoError(t, err)
	})

	t.Run("incorrect food", func(t *testing.T) {
		err := a.Feed("Grass")
		assert.Error(t, err)
	})
}

func TestAnimal_Heal(t *testing.T) {
	a := &animal.Animal{
		Status: animal.Sick,
	}

	a.Heal()

	assert.Equal(t, animal.Healthy, a.Status)
}

func TestAnimal_ChangeEnclosure(t *testing.T) {
	a := &animal.Animal{
		EnclosureID: 1,
	}

	a.ChangeEnclosure(2)

	assert.Equal(t, 2, a.EnclosureID)
}
