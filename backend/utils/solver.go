package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/shirakiya/sharqen/backend/conf"
)

type (
	CustomQueryRequest struct {
		Question string   `json:"question"`
		Choices  []string `json:"choices"`
		Solver   string   `json:"solver"`
	}

	CustomQueryResponse struct {
		CustomQuery string `json:"custom_query"`
	}
)

type (
	SolveRequest struct {
		Question string   `json:"question"`
		Choices  []string `json:"choices"`
		Solver   string   `json:"solver"`
	}

	SolveResponse struct {
		Answer  string `json:"answer"`
		Choices []struct {
			Candidates   []string `json:"candidates"`
			Normalized   string   `json:"normalized"`
			Surface      []string `json:"surface"`
			Unnormalized string   `json:"unnormalized"`
		} `json:"choices"`
		MatchCounts struct {
			Basic           []int `json:"basic"`
			Biased          []int `json:"biased"`
			Bigram          []int `json:"Bigram"`
			Transliteration []int `json:"transliteration"`
		} `json:"match_counts"`
		SearchResults []string `json:"search_results"`
	}
)

var client = &http.Client{Timeout: time.Duration(3) * time.Second}

func buildURL(endpoint string) string {
	parsedURL, _ := url.Parse(conf.GetConfig().SolverService.Url)
	parsedURL.Path = path.Join(parsedURL.Path, endpoint)

	return parsedURL.String()
}

// get is not used anymore..
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

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("Response Status Code: %d", resp.StatusCode)
		panic("Response had something error.")
	}

	return resp
}

func post(url string, paramsBytes []byte) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(paramsBytes))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		log.Printf("Response Status Code: %d", resp.StatusCode)
		panic("Response had something error.")
	}

	return resp
}

// GetCustomQuery returns question processed by Solver.
func GetCustomQuery(question string, choices []string) CustomQueryResponse {
	url := buildURL("custom")

	requestParams := CustomQueryRequest{
		Question: question,
		Choices:  choices,
		Solver:   "shiraki",
	}
	requestParamsBytes, _ := json.Marshal(requestParams)

	resp := post(url, requestParamsBytes)
	defer resp.Body.Close()

	customQueryResponse := CustomQueryResponse{}
	json.NewDecoder(resp.Body).Decode(&customQueryResponse)

	return customQueryResponse
}

// GetSolveResult returns answer from Solver.
func GetSolveResult(question string, choices []string) SolveResponse {
	url := buildURL("solve")

	requestParams := SolveRequest{
		Question: question,
		Choices:  choices,
		Solver:   "shiraki",
	}
	requestParamsBytes, _ := json.Marshal(requestParams)

	resp := post(url, requestParamsBytes)
	defer resp.Body.Close()

	solveResponse := SolveResponse{}
	json.NewDecoder(resp.Body).Decode(&solveResponse)

	return solveResponse
}
