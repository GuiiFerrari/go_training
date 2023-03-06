package algelin

import (
	"testing"
)

func Test_inner_product_2d(t *testing.T) {
	type args struct {
		point_a *Point2D
		point_b *Point2D
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "Test 1",
			args: args{
				point_a: &Point2D{X: 1.0, Y: 0.},
				point_b: &Point2D{X: 0., Y: 1.},
			},
			want: 0.,
		},
		{
			name: "Test 2",
			args: args{
				point_a: &Point2D{X: 0., Y: 0.},
				point_b: &Point2D{X: 0., Y: 1.},
			},
			want: 0.,
		},
		{
			name: "Test 3",
			args: args{
				point_a: &Point2D{X: 0., Y: 1.},
				point_b: &Point2D{X: 0., Y: 1.},
			},
			want: 1.,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InnerProduct2d(tt.args.point_a, tt.args.point_b); got != tt.want {
				t.Errorf("inner_product_2d() = %v, want %v", got, tt.want)
			}
		})
	}
}
