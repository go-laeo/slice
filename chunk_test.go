package slice

import (
	"reflect"
	"testing"
)

func TestChunkInts(t *testing.T) {
	type args struct {
		values []int
		length int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "chunk_ints_should_succeed",
			args: args{values: []int{0, 1, 2, 3, 1, 2, 3, 2, 3}, length: 3},
			want: [][]int{{0, 1, 2}, {3, 1, 2}, {3, 2, 3}},
		},
		{
			name: "chunk_ints_with_single_element_shoud_succeed",
			args: args{values: []int{0}, length: 3},
			want: [][]int{{0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.values, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunkStrings(t *testing.T) {
	type args struct {
		values []string
		length int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "chunk_strings_should_succeed",
			args: args{values: []string{"hello", "g", "o", "1.18", "generics"}, length: 2},
			want: [][]string{{"hello", "g"}, {"o", "1.18"}, {"generics"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.values, tt.args.length); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}
