package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sendJSON(response http.ResponseWriter, status int, data any) error {
	response.WriteHeader(status)
	response.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(response).Encode(data)
}

type HttpHandlerFunction func(response http.ResponseWriter, request *http.Request) error

type HttpError struct {
	Error string
}

func wrapHttpHandlerFunc(handler HttpHandlerFunction) http.HandlerFunc {
	return func(response http.ResponseWriter, request *http.Request) {
		if err := handler(response, request); err != nil {
			sendJSON(response, http.StatusBadRequest, HttpError{Error: err.Error()})
		}
	}
}

type Server struct {
	listenAddress string
}

func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
	}
}

func (server *Server) Run() {
	router := mux.NewRouter()
	router.HandleFunc("/account", wrapHttpHandlerFunc(server.handleAccount))
	log.Printf(fmt.Sprintf("Server started listening at address: %s", server.listenAddress))
	http.ListenAndServe(server.listenAddress, router)
}

func (server Server) handleAccount(response http.ResponseWriter, request *http.Request) error {
	switch request.Method {
	case "GET":
		return server.handleGetAccount(response, request)
	case "POST":
		return server.handleCreateAccount(response, request)
	case "DELETE":
		return server.handleDeleteAccount(response, request)
	}
	return fmt.Errorf("Method not allowed %s", request.Method)
}

func (server Server) handleGetAccount(response http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server Server) handleCreateAccount(response http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server Server) handleDeleteAccount(response http.ResponseWriter, request *http.Request) error {
	return nil
}

func (server Server) handleTransfer(response http.ResponseWriter, request *http.Request) error {
	return nil
}
