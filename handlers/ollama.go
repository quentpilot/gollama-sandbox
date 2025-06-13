package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func ollamaGenerateHandler(req OllamaRequest, res any) *HttpError {
	jsonData, _ := json.Marshal(req)
	resp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if err := json.Unmarshal(body, &res); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}
