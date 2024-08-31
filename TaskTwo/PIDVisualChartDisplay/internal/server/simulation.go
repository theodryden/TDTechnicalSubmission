package server

import (
    "log"
    "encoding/json"
    "net/http"
    "time"
    "github.com/felixge/pidctrl"
    "main.go/internal/simulation"
)

// Global variables for simulation
var (
    aperture       = 50.0
    targetPressure = 2.0
    currentPressure = 2.0
    pid             = pidctrl.NewPIDController(1.0, 0.1, 0.05)
)

func init() {
    pid.SetOutputLimits(0, 100)
    pid.Set(targetPressure)
}

func RunSimulation() {
    http.HandleFunc("/api/pressure", corsMiddleware(pressureHandler))
    http.HandleFunc("/api/aperture", corsMiddleware(apertureHandler))
    go startSimulation() // Start simulation in a separate goroutine
    http.ListenAndServe(":8080", nil)

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

func startSimulation() {
    for {
        // Simulate pressure fluctuations
        currentPressure = simulation.SimulatePressure(aperture)

        // Compute the PID adjustment
        adjustment := pid.Update(currentPressure)

        // Adjust the aperture based on PID output + Clamp keeps it within 0-100 range
        aperture = simulation.Clamp(aperture+adjustment, 0, 100)

        // Random delay to simulate aperture open/close time
        time.Sleep(simulation.RandomDelay())
    }
}

func pressureHandler(w http.ResponseWriter, r *http.Request) {
    // Return the current pressure as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]float64{"pressure": currentPressure})
}

func apertureHandler(w http.ResponseWriter, r *http.Request) {
    // Return the current aperture as JSON
    w.Header().Set("Content-Type","application/json")
    json.NewEncoder(w).Encode(map[string]float64{"aperture": aperture})
}


func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    }
}
