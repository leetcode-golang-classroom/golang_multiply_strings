package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	num1 := "123"
	num2 := "456"
	for idx := 0; idx < b.N; idx++ {
		multiply(num1, num2)
	}
}
func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "num1 = \"2\", num2 = \"3\"",
			args: args{num1: "2", num2: "3"},
			want: "6",
		},
		{
			name: "num1 = \"123\", num2 = \"456\"",
			args: args{num1: "123", num2: "456"},
			want: "56088",
		},
		{
			name: "num1 = \"408\", num2 = \"5\"",
			args: args{num1: "408", num2: "5"},
			want: "2040",
		},
		{
			name: "num1 = \"408\", num2 = \"0\"",
			args: args{num1: "408", num2: "0"},
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.args.num1, tt.args.num2); got != tt.want {
				t.Errorf("multiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
