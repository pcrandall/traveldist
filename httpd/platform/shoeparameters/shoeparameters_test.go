package shoeparameters

import "testing"

func TestGetAll(t *testing.T) {
	param := GetShoeParameters()
	if param.Check < 1 {
		t.Errorf("no params")
	}
}
