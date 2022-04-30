package pathing

import "testing"

func TestBFSSearch_ValidGraph_ValidPath(t *testing.T) {
	var testGraph = map[string][]string{
		"A": {"B", "C"},
		"B": {"C", "A"},
		"C": {"D"},
		"D": {},
	}

	getLinksTest := func(article string) ([]string, error) {
		return testGraph[article], nil
	}

	path, err := bfs("A", "D", getLinksTest)
	if err != nil {
		t.Fatal(err)
	}

	if !equalsSlice(path, []string{"A", "B", "C", "D"}) && !equalsSlice(path, []string{"A", "C", "D"}) {
		t.Fatalf("Incorrect path: %v", path)
	}
}

func TestBFSSearch_NoPath_Error(t *testing.T) {
	var testGraph = map[string][]string{
		"A": {"B", "C"},
		"D": {},
	}

	getLinksTest := func(article string) ([]string, error) {
		return testGraph[article], nil
	}

	path, err := bfs("A", "D", getLinksTest)
	if err == nil {
		t.Fatal(err)
	}
	if len(path) != 0 {
		t.Fatalf("Path should be empty: %v", path)
	}
}

func equalsSlice(first []string, second []string) bool {
	if len(first) != len(second) {
		return false
	}

	for index, value := range first {
		if value != second[index] {
			return false
		}
	}
	return true
}

func TestGetPath_ValidMap_ValidPath(t *testing.T) {
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
