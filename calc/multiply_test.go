// multiply_test.go
package calculator

import "testing"

func TestMultiply(t *testing.T) {
    tests := []struct {
        a, b, want float64
    }{
        {2, 3, 6},
        {-1, 1, -1},
        {0, 5, 0},
        {3.5, 2, 7},
    }

    for _, tt := range tests {
        got := Multiply(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Multiply(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
        }
    }
}
