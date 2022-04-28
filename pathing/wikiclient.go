package pathing

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type LinksEndpointOutput struct {
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

func getLinks(article string) ([]string, error) {
	linksEndpointOutput, err := wikipediaEndpointCall(article)
	if err != nil {
		return nil, err
	}

	links := filterLinks(linksEndpointOutput, article)

	return links, nil
}

func filterLinks(linksEndpointOutput *LinksEndpointOutput, parentArticle string) []string {
	filteredLinks := []string{}
	invalidPrefixes := []string{"File:", "Template:", "Category:", "Help:", "Special:", "Wikipedia:", "Portal:", "Template_talk:"}

	for _, linkObject := range linksEndpointOutput.Parse.Links {
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

func wikipediaEndpointCall(article string) (*LinksEndpointOutput, error) {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	endpoint := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=parse&format=json&page=%v&prop=links", article)
	resp, err := client.Get(endpoint)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var linkEndpointOutput LinksEndpointOutput
	if err := json.Unmarshal(respBody, &linkEndpointOutput); err != nil {
		return nil, err
	}

	return &linkEndpointOutput, nil
}
