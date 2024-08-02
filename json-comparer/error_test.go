package jsoncomparer

import (
	"errors"
	"testing"
)

func TestCompareErr_Error(t *testing.T) {
	tests := []struct {
		name string
		e    *CompareErr
		want string
	}{
		{
			name: "print err for a map",
			e:    &CompareErr{keys: []string{"here", "is", "a", "map"}, err: errors.New("what ever")},
			want: string("failed at `here.is.a.map`: what ever"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Error(); got != tt.want {
				t.Errorf("CompareErr.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
