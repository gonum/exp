package ivp

/* idea drafting for possible future shape of IVPs with an input vector.

// U is a vector which is a function of the current
// state. Put simply, the next input is a function of all current state variables
// and, possibly, current input as well.
//  U_next = F_u(t, X, U)
type NAIVP interface {
	// The input functions (ufunc) are not differential equations but rather
	// calculated directly from a given x vector and current input vector.
	IVP
	// Where F_u is ufunc as returned by Equations()
	// An initial value problem is characterized by boundary conditions imposed
	// on the state vector X at the beginning of the integration domain. These
	// boundary conditions are returned by the IV() method for the state vector
	// as x0 and for the input vector as u0.
	Inputs() (ufunc func(y []float64, t float64, x []float64))
	// Dimensions of x state vector and u inputs.	// If problem has no input functions then u supplied and ufunc returned
	// may be nil. x equations my not be nil.
	Dims() (nx, nu int)
}

*/
