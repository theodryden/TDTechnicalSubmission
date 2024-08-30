package main

import "main.go/internal/server"

func main() {
    server.RunSimulation()
}


/*
package main

import (
	"fmt"
	"time"
	"github.com/felixge/pidctrl"
	"main.go/internal/simulation"

)


func main() {

	// Initial configuration
	aperture := 50.0
	targetPressure := 2.0
	currentPressure := 2.0

    // Initialize the PID controller with tuning and setValue to 2
    pid := pidctrl.NewPIDController(1.0, 0.1, 0.05)
	pid.SetOutputLimits(0,100)
    pid.Set(targetPressure)


	// Create a simulation loop
	for {
		// Simulate pressure fluctuations
		currentPressure = simulation.SimulatePressure(aperture)

		// Compute the PID based on the current pressure (starts at 2.0)
		adjustment := pid.Update(currentPressure)

		// Adjust the aperture based on PID output + Clamp keeps it within 0-100 range
		aperture = simulation.Clamp(aperture+adjustment, 0,100)
		fmt.Printf("Aperture: %.2f%% \nPressure Sensor: %.2f bars\n", aperture, currentPressure,)

		// Random delay to simulate aperture open/close time
		time.Sleep(simulation.RandomDelay())
	}



}
*/


