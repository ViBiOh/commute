package coordinates

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParisKrakow(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		expected := 1275.569971915251

		got := LatLng{48.8566, 2.3522}.
			DistanceInKilometer(
				LatLng{50.0647, 19.9450})

		assert.Equal(t, expected, got)
	})

	t.Run("example", func(t *testing.T) {
		t.Parallel()

		expected := 5.926734214010531

		got := LatLng{2.990353, 101.533913}.
			DistanceInKilometer(
				LatLng{2.960148, 101.577888})

		assert.Equal(t, expected, got)
	})
}
