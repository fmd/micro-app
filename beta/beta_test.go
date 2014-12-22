package beta

import "testing"

var tests = []uint64{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

func TestFib(t *testing.T) {
    fn := Fib()
    for i, v := range tests {
        if val := fn(); val != v {
            t.Fatalf("at index %d, expected %d, got %d.", i, v, val)
        }
    }
}