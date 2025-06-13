package main

import (
	"fmt"
	"github/quentpilot/gollama-sandbox/handlers"
)

func main() {
	s := handlers.NewServer()

	s.Engine.POST("/ask", handlers.AskHandler())
	s.Engine.GET("/summarize", handlers.SummarizeHtmlHandler())

	fmt.Println("API Ollama Sandbox is running on localhost:8080")

	s.Engine.Run(":8080")
}
