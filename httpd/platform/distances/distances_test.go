package distances

import "testing"

func TestAdd(t *testing.T) {
	d := New()
	d.Add(Distance{})
	if len(d.Distances) != 1 {
		t.Errorf("Item was not added")
	}
}

func TestGetAll(t *testing.T) {
	d := New()
	d.Add(Distance{})
	results := d.GetAll()
	if len(results) != 1 {
		t.Errorf("Item was not added")
	}
}