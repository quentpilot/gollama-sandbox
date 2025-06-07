package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PromptRequest struct {
	Prompt string `json:"prompt"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func main() {
	r := gin.Default()

	r.POST("/ask", func(c *gin.Context) {
		var req PromptRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		ollamaReq := OllamaRequest{
			Model:  "llama3",
			Prompt: req.Prompt,
			Stream: false,
		}

		jsonData, _ := json.Marshal(ollamaReq)
		resp, err := http.Post("http://ollama:11434/api/generate", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ollama unreachable", "details": err.Error()})
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)

		var ollamaResp OllamaResponse
		if err := json.Unmarshal(body, &ollamaResp); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid Ollama response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": ollamaResp.Response})
	})

	r.Run(":8080")
}
