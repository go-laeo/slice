package slice

import "testing"

func TestContainString(t *testing.T) {
	type args struct {
		items []string
		one   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test_strings_should_succeed",
			args: args{
				items: []string{"go", "lang", "1.18"},
				one:   "go",
			},
			want: true,
		},
		{
			name: "test_strings_should_fails",
			args: args{
				items: []string{"go", "lang", "1.18"},
				one:   "hello",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contain(tt.args.items, tt.args.one); got != tt.want {
				t.Errorf("Contain() = %v, want %v", got, tt.want)
			}
		})
	}
}
