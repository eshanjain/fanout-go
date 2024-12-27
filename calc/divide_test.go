// divide_test.go
package calculator

import "testing"

func TestDivide(t *testing.T) {
    tests := []struct {
        a, b, want float64
    }{
        {6, 3, 2},
        {-1, 1, -1},
        {0, 5, 0},
        {10, 2, 5},
        {7, 0, 0}, // Test division by zero
    }

    for _, tt := range tests {
        got := Divide(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Divide(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
        }
    }
}
