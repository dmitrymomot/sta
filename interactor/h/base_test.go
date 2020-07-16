package h

import (
	"testing"
)

type testBaseNext struct {
	next Handler
}

// Exec function handles custom 1 rules
func (h testBaseNext) Exec(a, b, c bool) (string, error) {
	return "Z", nil
}

func TestBase_Exec(t *testing.T) {
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
		{"case M", fields{}, args{a: true, b: true, c: true}, P, false},
		{"case C", fields{}, args{a: true, b: true, c: false}, M, false},
		{"case T", fields{}, args{a: false, b: true, c: true}, T, false},
		{"case with next", fields{testBaseNext{}}, args{a: false, b: false, c: false}, "Z", false},
		{"error", fields{}, args{a: false, b: false, c: false}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Base{
				next: tt.fields.next,
			}
			got, err := h.Exec(tt.args.a, tt.args.b, tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("Base.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Base.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
