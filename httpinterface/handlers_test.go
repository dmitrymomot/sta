package httpinterface

import (
	"net/http"
	"reflect"
	"testing"
)

func TestInteractor(t *testing.T) {
	type args struct {
		i interactor
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Interactor(tt.args.i); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Interactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
