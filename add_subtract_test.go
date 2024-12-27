// add_subtract_test.go
package calculator

import "testing"

func TestAdd(t *testing.T) {
    tests := []struct {
        a, b, want float64
    }{
        {2, 3, 5},
        {-1, 1, 0},
        {0, 0, 0},
        {3.14, 2.86, 6},
    }

    for _, tt := range tests {
        got := Add(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Add(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
        }
    }
}

func TestSubtract(t *testing.T) {
    tests := []struct {
        a, b, want float64
    }{
        {5, 3, 2},
        {-1, 1, -2},
        {0, 0, 0},
        {3.14, 2.14, 1},
    }

    for _, tt := range tests {
        got := Subtract(tt.a, tt.b)
        if got != tt.want {
            t.Errorf("Subtract(%v, %v) = %v; want %v", tt.a, tt.b, got, tt.want)
        }
    }
}
