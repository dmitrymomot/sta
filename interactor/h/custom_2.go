package h

import "errors"

// Custom2 struct
type Custom2 struct {
	next Handler
}

// NewCustom2 is factory function,
// returns a new instance of the Custom2 structure
func NewCustom2(next Handler) Custom2 {
	return Custom2{next: next}
}

// Exec function handles custom 1 rules
func (h Custom2) Exec(a, b, c bool) (string, error) {
	if a && b && !c {
		return T, nil
	} else if a && !b && c {
		return M, nil
	}

	if h.next != nil {
		return h.next.Exec(a, b, c)
	}

	return "", errors.New("could not find any match")
}
