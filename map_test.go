package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	type args struct {
		values []int
		fn     func(v int) string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "map_ints_to_strings_should_succees",
			args: args{
				values: []int{1, 2, 3},
				fn:     func(v int) string { return strconv.Itoa(v) },
			},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.values, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}
