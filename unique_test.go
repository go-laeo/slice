package slice

import (
	"reflect"
	"testing"
)

func TestUniqueInts(t *testing.T) {
	type args struct {
		src []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "unique_ints_should_succeed",
			args: args{
				src: []int{0, 0, 1, 0, 2, 3, 4, 3, 1, 0},
			},
			want: []int{0, 1, 2, 3, 4},
		},
		{
			name: "unique_ints_with_one_element_should_succeed",
			args: args{
				src: []int{0},
			},
			want: []int{0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniqueStrings(t *testing.T) {
	type args struct {
		src []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "unique_strings_should_succeed",
			args: args{src: []string{"go", "go", "lang", " ", "1.18"}},
			want: []string{"go", "lang", " ", "1.18"},
		},
		{
			name: "unique_strings_with_one_element_should_succeed",
			args: args{src: []string{"go1.18"}},
			want: []string{"go1.18"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.src); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}
