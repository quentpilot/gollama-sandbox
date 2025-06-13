package handlers

import (
	"fmt"
	"github/quentpilot/gollama-sandbox/parser"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AskHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req AskRequest
		if err := c.BindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
			return
		}

		ollamaReq := OllamaRequest{
			Model:  "llama3.2",
			Prompt: req.Prompt,
			Stream: false,
		}

		var res *AskResponse
		err := ollamaGenerateHandler(ollamaReq, &res)
		if err != nil {
			c.JSON(err.Status(), gin.H{"message": err.Error()})
			return
		}

		c.JSON(http.StatusOK, res)
	}
}

func SummarizeHtmlHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Call endpoint /summarize")
		url := c.Query("url")
		if url == "" {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Missing url value"})
			return
		}

		content, err := parser.ParseHtmlContent(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Cannot parse url content: " + err.Error()})
			return
		}

		// Todo: API optimized to summarize wikipedia pages by parsing sections and RAG them with vectors
		prompt := fmt.Sprintf("You are a journalist. Summarize the following content in french, with maximum of three sentences: %s", content)

		ollamaReq := OllamaRequest{
			Model:  "llama3.2",
			Prompt: prompt,
			Stream: false,
		}
		var resume AskResponse
		ollamaGenerateHandler(ollamaReq, &resume)

		c.JSON(http.StatusOK, resume)
	}
}
