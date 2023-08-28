package ts_duration

import (
	"fmt"
	"testing"
	"time"
)

func TestDurationParseFromString(t *testing.T) {
	// Using ParseDuration() function
	hr, _ := time.ParseDuration("3h")
	comp, _ := time.ParseDuration("5h30m40s")
	seconds, _ := time.ParseDuration("3600s")

	fmt.Println("Time Duration 1: ", hr)
	fmt.Println("Time Duration 2: ", comp)
	fmt.Println("Time Duration 3: ", seconds)
	fmt.Println("Seconds: ", hr.Seconds())
}
