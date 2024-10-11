package lib

import "testing"

func TestInit(t *testing.T) {
	ry, err := NewRyzenAccess()
	if err != nil {
		t.Error(err)
		return
	}
	defer ry.Cleanup()
	err = ry.Refresh()
	if err != nil {
		t.Error(err)
	}
}
