package h

import (
	"testing"
)

type testCustom2Next struct {
	next Handler
}

// Exec function handles custom 1 rules
func (h testCustom2Next) Exec(a, b, c bool) (string, error) {
	return "P", nil
}

func TestCustom2_Exec(t *testing.T) {
	type fields struct {
		next Handler
	}
	type args struct {
		a bool
		b bool
		c bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{"case M", fields{}, args{a: true, b: false, c: true}, M, false},
		{"case T", fields{}, args{a: true, b: true, c: false}, T, false},
		{"case with next", fields{testCustom2Next{}}, args{a: false, b: false, c: false}, "P", false},
		{"error", fields{}, args{a: false, b: false, c: false}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Custom2{
				next: tt.fields.next,
			}
			got, err := h.Exec(tt.args.a, tt.args.b, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Custom2.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Custom2.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
