package k

import (
	"testing"

	"github.com/dmitrymomot/sta/interactor/h"
)

type testCustom2Next struct {
	next Handler
}

// Calc function calculates and returns K value
func (i testCustom2Next) Calc(v string, d float64, e, f int) float64 {
	switch v {
	case "Z":
		return 12345
	}
	return 0
}

func TestCustom2_Calc(t *testing.T) {
	type fields struct {
		next Handler
	}
	type args struct {
		v string
		d float64
		e int
		f int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
	}{
		{"case M", fields{}, args{h.M, 3.3, 10, 10}, 13.63},
		{"case with next", fields{testBaseNext{}}, args{"Z", 3.3, 0, 10}, 12345},
		{"case undefined", fields{}, args{"", 3.3, 10, 10}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Custom2{
				next: tt.fields.next,
			}
			if got := i.Calc(tt.args.v, tt.args.d, tt.args.e, tt.args.f); got != tt.want {
				t.Errorf("Custom2.Calc() = %v, want %v", got, tt.want)
			}
		})
	}
}
