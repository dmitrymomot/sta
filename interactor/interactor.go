package interactor

type (
	// HHandler interface
	hHandler interface {
		Exec(a, b, c bool) (string, error)
	}

	// KHandler interface
	kHandler interface {
		Calc(h string, d float64, e, f int) float64
	}

	// Interactor struct
	Interactor struct {
		hh hHandler
		kh kHandler
	}
)

// New is factory function,
// returns a new instance of the Interactor structure
func New(hh hHandler, kh kHandler) Interactor {
	return Interactor{hh: hh, kh: kh}
}

// Exec func handles options and return completed result
func (i Interactor) Exec(a, b, c bool, d float64, e, f int) (string, float64, error) {
	h, err := i.hh.Exec(a, b, c)
	if err != nil {
		return "", 0, err
	}

	return h, i.kh.Calc(h, d, e, f), nil
}
