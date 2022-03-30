package slice

import (
	"reflect"
	"strconv"
	"testing"
)

func TestReduce(t *testing.T) {
	type args struct {
		values  []int
		fn      func(carry string, item int) string
		initial string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "reduce_ints_to_string_should_succeed",
			args: args{
				values:  []int{1, 2, 3},
				fn:      func(carry string, item int) string { return carry + strconv.Itoa(item) },
				initial: "",
			},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Reduce(tt.args.values, tt.args.fn, tt.args.initial); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reduce() = %v, want %v", got, tt.want)
			}
		})
	}
}
