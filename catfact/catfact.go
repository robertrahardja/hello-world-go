package catfact

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// CatFact represents a cat fact from the API
type CatFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

// FetchCatFact fetches a random cat fact from the API
func FetchCatFact() (CatFact, error) {
	url := "https://catfact.ninja/fact"
	resp, err := http.Get(url)
	if err != nil {
		return CatFact{}, fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var catFact CatFact
	if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
		return CatFact{}, fmt.Errorf("error decoding JSON: %w", err)
	}

	return catFact, nil
}
