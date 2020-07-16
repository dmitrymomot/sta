package k

import "github.com/dmitrymomot/sta/interactor/h"

// Custom2 struct
type Custom2 struct {
	next Handler
}

// NewCustom2 is factory function,
// returns a new instance of the Custom2 structure
func NewCustom2(next Handler) Custom2 {
	return Custom2{next: next}
}

// Calc function calculates and returns K value
func (i Custom2) Calc(v string, d float64, e, f int) float64 {
	switch v {
	case h.M:
		return float64(f) + d + (d * float64(e) / 100)
	}

	if i.next != nil {
		return i.next.Calc(v, d, e, f)
	}

	return 0
}
