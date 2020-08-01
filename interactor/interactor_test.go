package interactor

import (
	"testing"

	"github.com/dmitrymomot/sta/interactor/h"
	"github.com/dmitrymomot/sta/interactor/k"
)

func TestInteractor_Exec(t *testing.T) {
	type fields struct {
		hh hHandler
		kh kHandler
	}
	type args struct {
		a bool
		b bool
		c bool
		d float64
		e int
		f int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   float64
		wantErr bool
	}{
		{"only base", fields{h.NewBase(nil), k.NewBase(nil)}, args{a: true, b: true, c: true, d: 25.5, e: 3, f: 2}, h.P, 26.5, false},
		{
			"full",
			fields{
				h.NewCustom2(h.NewBase(nil)),
				k.NewCustom2(k.NewCustom1(k.NewBase(nil))),
			},
			args{a: true, b: true, c: true, d: 3.3, e: 10, f: 10},
			h.P,
			6.93,
			false,
		},
		{"error", fields{h.NewBase(nil), k.NewBase(nil)}, args{a: false, b: false, c: false, d: 1, e: 1, f: 1}, "", 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := Interactor{
				hh: tt.fields.hh,
				kh: tt.fields.kh,
			}
			got, got1, err := i.Exec(tt.args.a, tt.args.b, tt.args.c, tt.args.d, tt.args.e, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Interactor.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Interactor.Exec() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Interactor.Exec() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
