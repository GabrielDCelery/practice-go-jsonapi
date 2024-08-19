package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func sendJSON(response http.ResponseWriter, status int, data any) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
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
	store         Store
}

func NewServer(listenAddress string, store Store) *Server {
	return &Server{
		listenAddress: listenAddress,
		store:         store,
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
	mockAccount := NewAccount("Gabor", "Zeller")
	return sendJSON(response, http.StatusOK, mockAccount)
}

func (server Server) handleCreateAccount(response http.ResponseWriter, request *http.Request) error {
	createAccountRequest := &CreateAccountRequest{}
	if err := json.NewDecoder(request.Body).Decode(createAccountRequest); err != nil {
		return err
	}
	account := NewAccount(createAccountRequest.FirstName, createAccountRequest.LastName)
	if err := server.store.CreateAccount(account); err != nil {
		return err
	}
	return sendJSON(response, http.StatusOK, account)
}

func (server Server) handleDeleteAccount(response http.ResponseWriter, request *http.Request) error {
	deleteAccountRequest := &DeleteAccountRequest{}
	if err := json.NewDecoder(request.Body).Decode(deleteAccountRequest); err != nil {
		return err
	}
	if err := server.store.DeleteAccountByID(deleteAccountRequest.ID); err != nil {
		return err
	}
	return sendJSON(response, http.StatusOK, deleteAccountRequest)
}

func (server Server) handleTransfer(response http.ResponseWriter, request *http.Request) error {
	return nil
}
