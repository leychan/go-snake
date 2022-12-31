package main

import (
	"testing"
)

func TestDrawer_IsFood(t *testing.T) {
	l := Location{X: 1, Y: 1}
	d := &Drawer{Food: Location{X: 1, Y: 2}}
	got := d.IsFood(l)
	if false != got {
		t.Errorf("expected: %t, got: %t", false, got)
	}
}
