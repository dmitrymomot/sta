package h

import "errors"

// Predefined result values
const (
	M string = "M"
	P string = "P"
	T string = "T"
)

type (
	// Handler interface
	Handler interface {
		Exec(a, b, c bool) (string, error)
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

// Exec function handles base rules
func (h Base) Exec(a, b, c bool) (string, error) {
	if a && b && !c {
		return M, nil
	} else if a && b && c {
		return P, nil
	} else if !a && b && c {
		return T, nil
	}

	if h.next != nil {
		return h.next.Exec(a, b, c)
	}

	return "", errors.New("could not find any match")
}
