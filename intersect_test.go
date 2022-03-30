package slice

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "intersect_ints_should_succeed",
			args: args{
				a: []int{1, 1, 2, 3, 2, 3, 4},
				b: []int{1, 2, 3},
			},
			want: []int{1, 1, 2, 3, 2, 3},
		},
		{
			name: "intersect_empty_ints_a_should_succeed",
			args: args{
				a: []int{},
				b: []int{1, 2, 3},
			},
			want: []int{},
		},
		{
			name: "intersect_empty_ints_b_should_succeed",
			args: args{
				a: []int{1, 2, 2, 12, 3},
				b: []int{},
			},
			want: []int{},
		},
		{
			name: "intersect_ints_with_even_elements_should_succeed",
			args: args{
				a: []int{1, 1, 2, 3, 2, 3},
				b: []int{1, 2},
			},
			want: []int{1, 1, 2, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}
