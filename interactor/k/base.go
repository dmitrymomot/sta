package k

import "github.com/dmitrymomot/sta/interactor/h"

type (
	// Handler interface
	Handler interface {
		Calc(v string, d float64, e, f int) float64
	}

	// Base struct
	Base struct {
		next Handler
	}
)

// NewBase is factory function,
// returns a new instance of the Base structure
func NewBase(next Handler) Base {
	return Base{next: next}
}

// Calc function calculates and returns K value
func (i Base) Calc(v string, d float64, e, f int) float64 {
	switch v {
	case h.M:
		return d + (d * float64(e) / 10)
	case h.P:
		return d + (d * float64(e-f) / 25.5)
	case h.T:
		return d + (d * float64(f) / 30)
	}

	if i.next != nil {
		return i.next.Calc(v, d, e, f)
	}

	return 0
}
