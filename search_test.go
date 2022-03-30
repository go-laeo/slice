package slice

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		values []int
		fn     func(v int) bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "search_int_should_succeed",
			args: args{
				values: []int{1, 2, 3, 4, 5},
				fn:     func(v int) bool { return v == 3 },
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Search(tt.args.values, tt.args.fn); got != tt.want {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
