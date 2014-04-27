package game_engine

import (
    "testing"
)

func TestNations(t *testing.T) {
    if int(red) != 0 {
        t.Fatalf("Red should be first 0")
    }
    if int(blue) != 1 {
        t.Fatalf("Blue should be second 1")
    }
    if int(green) != 2 {
        t.Fatalf("Red should be third 2")
    }
    if int(yellow) != 3 {
        t.Fatalf("Red should be fourth 3")
    }
}
