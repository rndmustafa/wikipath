package pathing

import "testing"

func TestFilterLinks_MixedInput_FilteredLinks(t *testing.T) {
	output := &ParseOutput{
		Parse: ParseObject{
			Links: []LinkObject{
				{ArticleName: "Flower"},
				{ArticleName: "Cannon"},
				{ArticleName: "File:Rose"},
				{ArticleName: "Rose"},
				{ArticleName: "Portal:Other"},
			},
		},
	}

	articles := filterLinks(output, "Rose")
	if !contains(articles, "Flower") {
		t.Fatalf("Flower not found in array %v", articles)
	}
	if !contains(articles, "Cannon") {
		t.Fatalf("Cannon not found in array %v", articles)
	}
	if contains(articles, "Rose") {
		t.Fatalf("Rose matches the article name and shouldn't be in the array %v", articles)
	}
	if contains(articles, "File:Rose") {
		t.Fatalf("File:Rose is a invalid article %v", articles)
	}
	if contains(articles, "Portal:Other") {
		t.Fatalf("Portal:Other is a invalid article %v", articles)
	}
}
