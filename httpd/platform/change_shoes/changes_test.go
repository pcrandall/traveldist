package change

import "testing"

func TestAdd(t *testing.T) {
	c := New()
	c.Add(Change{})
	if len(c.Changes) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	c := New()
	c.Add(Change{})
	results := c.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}
