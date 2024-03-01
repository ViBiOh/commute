package coordinates

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

const EarthRadius = 6371

var ErrLatLngNeedsExactlyTwoValues = errors.New("LatLng needs exactly two values")

type LatLng [2]float64

func NewLatLng(input []float64) (LatLng, error) {
	if len(input) != 2 {
		return LatLng{}, ErrLatLngNeedsExactlyTwoValues
	}

	return LatLng{input[0], input[1]}, nil
}

func ParseLatLng(raw string) (LatLng, error) {
	parts := strings.Split(raw, ",")
	if len(parts) != 2 {
		return LatLng{}, fmt.Errorf("`%s`: %w", raw, ErrLatLngNeedsExactlyTwoValues)
	}

	lat, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return LatLng{}, fmt.Errorf("latitude is not a float `%s`", raw)
	}

	lng, err := strconv.ParseFloat(parts[1], 64)
	if err != nil {
		return LatLng{}, fmt.Errorf("longitude is not a float `%s`", raw)
	}

	return LatLng{lat, lng}, nil
}

func (ll LatLng) String() string {
	return fmt.Sprintf("%f,%f", ll[0], ll[1])
}

func (ll LatLng) LngLat() string {
	return fmt.Sprintf("%f,%f", ll[1], ll[0])
}

func (ll LatLng) Lat() float64 {
	return ll[0]
}

func (ll LatLng) Lng() float64 {
	return ll[1]
}

func (ll LatLng) IsWithin(b LatLng, distance float64) bool {
	return ll.DistanceInKilometer(b) <= distance
}

func (l LatLng) DistanceInKilometer(b LatLng) float64 {
	// 2R × sin⁻¹(√[sin²((θ₂ - θ₁)/2) + cosθ₁ × cosθ₂ × sin²((φ₂ - φ₁)/2)])

	sinLat := math.Sin(degreesToRadians(b.Lat()-l.Lat()) / 2)
	sinLng := math.Sin(degreesToRadians(b.Lng()-l.Lng()) / 2)

	a := sinLat*sinLat + sinLng*sinLng*math.Cos(degreesToRadians(l.Lat()))*math.Cos(degreesToRadians(b.Lat()))

	return 2 * EarthRadius * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
}

func degreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

type List []LatLng

func (l List) Center() LatLng {
	if len(l) == 0 {
		return LatLng{}
	}

	var output LatLng

	for _, coord := range l {
		output[0] += coord[0]
		output[1] += coord[1]
	}

	output[0] = output[0] / float64(len(l))
	output[1] = output[1] / float64(len(l))

	return output
}

func (l List) Clusters(distance float64) List {
	var groups List

	var coords List = l

	for len(coords) > 0 {
		var next List

		current := coords[0]
		group := List{current}

		for i := 1; i < len(coords); i++ {
			following := coords[i]

			if current.IsWithin(following, distance) {
				group = append(group, following)
			} else {
				next = append(next, following)
			}
		}

		groups = append(groups, group.Center())
		coords = next
	}

	return groups
}
