package check

import "testing"

func TestAdd(t *testing.T) {
	c := New()
	c.Add(Check{})
	if len(c.Checks) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	c := New()
	c.Add(Check{})
	results := c.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
