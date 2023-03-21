package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error
type apiError struct {
	Error string
}

func makeHttpRequestHandler(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			//handle the error
			WriteJSON(w, http.StatusBadRequest, apiError{Error: err.Error()})
		}
	}
}

type ApiServer struct {
	listAddr string
}

func NewApiServer(listenAddress string) *ApiServer {
	return &ApiServer{
		listAddr: listenAddress,
	}
}

func (s *ApiServer) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHttpRequestHandler(s.handleAccount))

	http.ListenAndServe(s.listAddr, router)

}

func (s *ApiServer) handleAccount(writer http.ResponseWriter, request *http.Request) error {
	if request.Method == "GET" {
		return s.handleGetAccount(writer, request)
	}
	if request.Method == "DELETE" {
		return s.handleDeleteAccount(writer, request)
	}
	if request.Method == "POST" {
		return s.handleCreateAccount(writer, request)
	}
	return fmt.Errorf("method now allowed")
}

func (s *ApiServer) handleGetAccount(writer http.ResponseWriter, request *http.Request) error {
	account := NewAccount("Connor", "Nusser")
	return WriteJSON(writer, http.StatusOK, account)
}
func (s *ApiServer) handleDeleteAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
func (s *ApiServer) handleCreateAccount(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
func (s *ApiServer) handleTransfer(writer http.ResponseWriter, request *http.Request) error {
	return nil
}
