// Package gotator is a REST API client for https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/PubTator/
// see also: https://www.ncbi.nlm.nih.gov/research/bionlp/APIs/usage/
package gotator

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Client is needed to access API functions
type Client struct {
	BaseURL    *url.URL
	httpClient *http.Client
}

// WordSpan contains begin and end position of strike's text
type WordSpan struct {
	Begin int `json:"begin"`
	End   int `json:"end"`
}

// Annotation is one 'strike', i.e. found gene or other concept
type Annotation struct {
	// Type, i.e. Gene, Disease, Chemical, Species, Mutation.
	Obj  string   `json:"obj"`
	Span WordSpan `json:"span"`
}

// Article wraps whole response body
// TODO: PubTator API seems to have problems with multiple ids, so we always request one
type Article struct {
	Db    string       `json:"sourcedb"`
	ID    string       `json:"sourceid"`
	Text  string       `json:"text"`
	Annos []Annotation `json:"denotations"`
}

// NewClient creates new API client to be used for all functions
func (c *Client) NewClient(baseURL string) (*Client, error) {
	base, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, err
	}
	httpClient := http.DefaultClient
	client := Client{base, httpClient}

	return &client, nil
}

// GetAllAnnotations gets list of annotations for one article by its pmid
// E.g.: https://www.ncbi.nlm.nih.gov/CBBresearch/Lu/Demo/RESTful/tmTool.cgi/BioConcept/28785587/JSON/
func (c *Client) GetAllAnnotations(id string) (*Article, error) {
	u, err := url.Parse(fmt.Sprintf("%v/BioConcept/%v/JSON/", c.BaseURL, id))
	if err != nil {
		return nil, err
	}
	//log.Printf("Gotator: http get: %v", u.String())
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var article Article
	err = json.NewDecoder(resp.Body).Decode(&article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}
