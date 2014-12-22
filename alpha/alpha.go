package alpha

import (
    "fmt"
)

func Fizz() []string {
    r := []string{}
    for i := 0; i < 100; i++ {
        s := ""
        if i % 3 == 0 {
            s = "Fizz"
        }
        if i % 5 == 0 {
            s = fmt.Sprintf("%sBuzz", s)
        }
        if len(s) == 0 {
            s = fmt.Sprintf("%d", i)
        }
        r = append(r, s)
    }
    return r
}