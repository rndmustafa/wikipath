package pathing

import "testing"

func TestGetPath(t *testing.T) {
	cameFrom := map[string]string{
		"D": "C",
		"C": "B",
		"B": "A",
	}
	path := getPath("D", cameFrom)

	valueAtIndexEquals(path, 0, "A", t)
	valueAtIndexEquals(path, 1, "B", t)
	valueAtIndexEquals(path, 2, "C", t)
	valueAtIndexEquals(path, 3, "D", t)
}

func valueAtIndexEquals(path []string, index int, value string, t *testing.T) {
	if path[index] != value {
		t.Fatalf("value at index %v was not %v: %v", index, value, path)
	}
}
