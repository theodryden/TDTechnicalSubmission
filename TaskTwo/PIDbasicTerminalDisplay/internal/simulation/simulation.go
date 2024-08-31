package simulation

import (
	"math/rand"
)

// Simulate pressure changes based on aperture
func SimulatePressure(aperture float64) float64 {
	basePressure := 2.0
	// Calculate the maximum fluctuation based on the aperture size
	maxFluctuation := 1.0 - (aperture / 100.0)
	// Generate a random fluctuation in the range [-maxFluctuation, maxFluctuation]
	randFluctuation := (rand.Float64()*2.0 - 1.0) * maxFluctuation
	// Apply the fluctuation to the base pressure
	return basePressure + randFluctuation
}

// Ensures value keeps within the min max range
func Clamp(value, min, max float64) float64{
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}