package jokes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"workshop/internal/api"
)

const getJokePath = "/api?format=json"

type JokeClient struct {
	url string
}

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (j JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := j.url + getJokePath

	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request error: %s", resp.Status)
	}

	var data api.JokeResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
