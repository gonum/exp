package ivp

import "errors"

// Configuration represents a change in default parameters in the solving
// of IVPs. Parameters are often needed to enable step size adaptivity and to
// overwrite solver default values.
type Configuration func(*parameters) error

type parameters struct {
	// step length
	minStep, maxStep float64
	// Relative and absolute error tolerances
	rtol, atol float64
	// parameters for step size selection.
	// The new step size is subject to restriction  minStepRatio <= hnew/hold <= maxStepRatio
	minStepRatio, maxStepRatio float64
}

// ConfigStepLimits sets minimum step length and max step length
// an integrator can take when using an adaptive method.
func ConfigStepLimits(minStep, maxStep float64) Configuration {
	return func(p *parameters) error {
		if minStep <= 0 || minStep > maxStep {
			return errors.New("minimum step length too small or greater than max step length")
		}
		p.maxStep, p.minStep = maxStep, minStep
		return nil
	}
}

// ConfigScalarTolerance Sets relative and absolute tolerances for controlling step error.
// If value passed is zero the corresponding tolerance is not modified.
func ConfigScalarTolerance(rtol, atol float64) Configuration {
	return func(p *parameters) error {
		if rtol < 0 || atol < 0 {
			return errors.New("negative error tolerance")
		}
		if rtol > 0 {
			p.rtol = rtol
		}
		if atol > 0 {
			p.atol = atol
		}
		return nil
	}
}

// ConfigAdaptiveRatioLimits sets parameters for step size selection.
// The new step size is subject to restriction  min <= hnew/hold <= max.
// If min or max are zero then value is not changed
func ConfigAdaptiveRatioLimits(min, max float64) Configuration {
	return func(p *parameters) error {
		if min < 0 || max < min {
			return errors.New("bad step size selection ratio")
		}
		if min > 0 {
			p.minStepRatio = min
		}
		if max > 0 {
			p.maxStepRatio = max
		}
		return nil
	}
}
