package app

import (
	"fmt"
	"net/http"
)

type HTTPHandler struct {
	userUseCase *UserUseCase
}

func NewHTTPHandler(userUseCase *UserUseCase) *HTTPHandler {
	return &HTTPHandler{
		userUseCase: userUseCase,
	}
}

func (h *HTTPHandler) StartServer() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {})
	http.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {})

	fmt.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
