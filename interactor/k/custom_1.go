package k

import (
	"github.com/dmitrymomot/sta/interactor/h"
)

// Custom1 struct
type Custom1 struct {
	next Handler
}

// NewCustom1 is factory function,
// returns a new instance of the Custom1 structure
func NewCustom1(next Handler) Custom1 {
	return Custom1{next: next}
}

// Calc function calculates and returns K value
func (i Custom1) Calc(v string, d float64, e, f int) float64 {
	switch v {
	case h.P:
		return 2*d + (d * float64(e) / 100)
	}

	if i.next != nil {
		return i.next.Calc(v, d, e, f)
	}

	return 0
}
