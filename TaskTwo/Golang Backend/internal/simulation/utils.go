package simulation

import (
	"math/rand"
	"time"
)


// Random Delay Function
func RandomDelay() time.Duration {
	delayMiliseconds := rand.Intn(2001)
	return time.Duration(delayMiliseconds) * time.Millisecond
}

