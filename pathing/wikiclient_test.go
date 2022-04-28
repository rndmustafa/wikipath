package pathing

import "testing"

func TestGetLinks_ValidArticle_NoErrors(t *testing.T) {
	_, err := getLinks("bee")
	if err != nil {
		t.Fatal(err)
	}
}
