package latlng

import "math"

const EarthRadius = 6371

type LatLng [2]float64

func (l LatLng) Lat() float64 {
	return l[0]
}

func (l LatLng) Lng() float64 {
	return l[1]
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
