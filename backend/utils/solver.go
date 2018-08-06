package utils

import (
	"encoding/json"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/shirakiya/sharqen/backend/conf"
)

type (
	CustomQueryResponse struct {
		CustomQuery string `json:"custom_query"`
	}

	Answer struct {
		Choice string `json:"choice"`
		Index  int    `json:"index"`
	}

	MatchCounts struct {
		Basic           []int `json:"basic"`
		Biased          []int `json:"biased"`
		Bigram          []int `json:"Bigram"`
		Transliteration []int `json:"transliteration"`
	}

	Choice struct {
		Candidates   []string `json:"candidates"`
		Normalized   string   `json:"normalized"`
		Surface      []string `json:"surface"`
		Unnormalized string   `json:"unnormalized"`
	}

	SolveResultResponse struct {
		Answer        Answer      `json:"answer"`
		Choices       []Choice    `json:"choices"`
		MatchCounts   MatchCounts `json:"match_counts"`
		Question      string      `json:"question"`
		SearchResults []string    `json:"search_results"`
	}
)

var client = &http.Client{Timeout: time.Duration(3) * time.Second}

func buildUrl(endpoint string) string {
	parsedUrl, _ := url.Parse(conf.GetConfig().SolverService.Url)
	parsedUrl.Path = path.Join(parsedUrl.Path, "custom")

	return parsedUrl.String()
}

func get(url string, values url.Values) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.URL.RawQuery = values.Encode()
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}

func GetCustomQuery(question string, choices []string) CustomQueryResponse {
	values := url.Values{}
	values.Add("question", question)
	for _, choice := range choices {
		values.Add("choices[]", choice)
	}

	resp := get(buildUrl("custom"), values)
	defer resp.Body.Close()

	customQueryResponse := CustomQueryResponse{}
	json.NewDecoder(resp.Body).Decode(&customQueryResponse)

	return customQueryResponse
}

func GetSolveResult(question string, choices []string) SolveResultResponse {
	values := url.Values{}
	values.Add("question", question)
	for _, choice := range choices {
		values.Add("choices[]", choice)
	}

	resp := get(buildUrl("solve"), values)
	defer resp.Body.Close()

	solveResultResponse := SolveResultResponse{}
	json.NewDecoder(resp.Body).Decode(&solveResultResponse)

	return solveResultResponse
}
