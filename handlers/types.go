package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	return &Server{
		Engine: gin.Default(),
	}
}

type AskRequest struct {
	Prompt string `json:"prompt"`
}

type AskResponse struct {
	Response string `json:"response"`
}

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type HttpError struct {
	Err  error
	Code int
}

func NewHttpError(err error, status int) *HttpError {
	return &HttpError{
		Err:  err,
		Code: status,
	}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HttpError: %s", e.Err.Error())
}

func (e *HttpError) Status() int {
	return e.Code
}

func NewOllamaRequest(prompt string, model string) *OllamaRequest {
	return &OllamaRequest{
		Model:  model,
		Prompt: prompt,
		Stream: false,
	}
}
