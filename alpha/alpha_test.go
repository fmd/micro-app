package alpha

import (
    "testing"
    "fmt"
)

func TestFizz(t *testing.T) {
    f := Fizz()
    if len(f) != 100 {
        t.Fatalf("There aren't 100 in the fizzbuzz!")
    }
    if f[25] != "Buzz" {
        fmt.Println(f[25])
        t.Fatalf("Error with modulo 5.")
    }
    if f[27] != "Fizz" {
        t.Fatalf("Error with modulo 3.")
    }
    if f[15] != "FizzBuzz" {
        t.Fatalf("Error with modulo combination.")
    }
    if f[7] != "7" {
        t.Fatalf("Error leaving numbers alone.")
    }
}