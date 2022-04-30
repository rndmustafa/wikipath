package pathing

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type ParseOutput struct {
	Parse ParseObject
}

type ParseObject struct {
	Title  string
	Pageid int
	Links  []LinkObject
}

type LinkObject struct {
	Ns          int
	Exists      string
	ArticleName string `json:"*"`
}

var client = http.Client{
	Timeout: time.Second * 10,
}

func getLinks(article string) ([]string, error) {
	parseOutput, err := wikipediaEndpointCall(article)
	if err != nil {
		return nil, err
	}

	links := filterLinks(parseOutput, article)

	return links, nil
}

func filterLinks(parseOutput *ParseOutput, parentArticle string) []string {
	filteredLinks := []string{}
	invalidPrefixes := []string{"File:", "Template:", "Category:", "Help:", "Special:", "Wikipedia:", "Portal:", "Template_talk:"}

	for _, linkObject := range parseOutput.Parse.Links {
		if validLink(linkObject.ArticleName, parentArticle, invalidPrefixes) {
			filteredLinks = append(filteredLinks, linkObject.ArticleName)
		}
	}

	return filteredLinks
}

func validLink(articleName string, parentArticle string, invalidPrefixes []string) bool {
	if articleName == parentArticle {
		return false
	}

	for _, prefix := range invalidPrefixes {
		if strings.HasPrefix(articleName, prefix) {
			return false
		}
	}

	return true
}

func wikipediaEndpointCall(article string) (*ParseOutput, error) {
	endpoint := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&format=json&page=%v&prop=links", url.QueryEscape(article))
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("Accept-Encoding", "gzip")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call %v: %w", endpoint, err)
	}

	var reader io.ReadCloser
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
	default:
		reader = resp.Body
	}

	respBody, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var parseOutput ParseOutput
	if err := json.Unmarshal(respBody, &parseOutput); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response body for article %v, body: %v error: %w", article, string(respBody), err)
	}

	return &parseOutput, nil
}
