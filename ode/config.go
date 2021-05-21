package ode

// Parameters represents the most common configuration parameters
type Parameters struct {
	// Permissible tolerance given an adaptive method
	AbsTolerance float64
	// Minimum/Maximum step allowed for a single iteration
	MinStep, MaxStep float64
}

var DefaultParam = Parameters{}
