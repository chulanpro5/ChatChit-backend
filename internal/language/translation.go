package language

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TranslationRequest struct {
	Text           string `json:"text"`
	TargetLanguage string `json:"target_language"`
}

type TranslationResponse struct {
	Translation string `json:"translation"`
}

func TranslateText(text, targetLanguage string, apiUrl string) (string, error) {
	url := apiUrl + "/translate"

	// Create a JSON request body
	requestBody, err := json.Marshal(TranslationRequest{
		Text:           text,
		TargetLanguage: targetLanguage,
	})

	if err != nil {
		return "", err
	}

	// Make a POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Parse response body
	var response TranslationResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	return response.Translation, nil
}
