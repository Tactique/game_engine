package game

import (
	"testing"
)

func TestTerrainConsts(t *testing.T) {
	if int(plains) != 0 {
		t.Fatalf("Plains should be 0")
	}
}
