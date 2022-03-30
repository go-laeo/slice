package slice

import (
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	type args struct {
		values []int
		fn     func(v int) bool
	}
	tests := []struct {
		name      string
		args      args
		wantV     int
		wantFound bool
	}{
		{
			name: "find_int_should_succeed",
			args: args{
				values: []int{1, 2, 3, 4, 5},
				fn:     func(v int) bool { return v == 5 },
			},
			wantV:     5,
			wantFound: true,
		},
		{
			name: "find_int_should_failed",
			args: args{
				values: []int{1, 2, 3, 4, 5},
				fn:     func(v int) bool { return v == 6 },
			},
			wantV:     0,
			wantFound: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotV, gotFound := Find(tt.args.values, tt.args.fn)
			if !reflect.DeepEqual(gotV, tt.wantV) {
				t.Errorf("Find() gotV = %v, want %v", gotV, tt.wantV)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Find() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
