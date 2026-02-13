package main
import (
    "testing"
)

func TestExample(t *testing.T) {
    in := 5
    out := 5
    result := Example(in)
    if result != out {
        t.Errorf(`Hello(%d) gave %d, was expecting %d`, in, result, out)
    }
}
