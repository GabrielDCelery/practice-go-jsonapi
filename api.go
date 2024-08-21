package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

func sendJSON(response http.ResponseWriter, status int, data any) error {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(status)
	return json.NewEncoder(response).Encode(data)
}

type HttpHandlerFunction func(response http.ResponseWriter, request *http.Request) error

type HttpError struct {
	Error string `json:"error"`
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
	router.HandleFunc("/account", wrapHttpHandlerFunc(server.handleGetAccounts)).Methods("GET")
	router.HandleFunc("/account", wrapHttpHandlerFunc(server.handleCreateAccount)).Methods("POST")
	router.HandleFunc("/account/{id}", wrapHttpHandlerFunc(server.handleGetAccount)).Methods("GET")
	router.HandleFunc("/account/{id}", wrapHttpHandlerFunc(server.handleDeleteAccount)).Methods("DELETE")
	log.Printf(fmt.Sprintf("Server started listening at address: %s", server.listenAddress))
	http.ListenAndServe(server.listenAddress, router)
}

func (server Server) handleGetAccount(response http.ResponseWriter, request *http.Request) error {
	vars := mux.Vars(request)
	accountID := vars["id"]
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(accountID) {
		return fmt.Errorf("Invalid account ID: %s", accountID)
	}
	err, account := server.store.GetAccountByID(accountID)
	if err != nil {
		return err
	}
	return sendJSON(response, http.StatusOK, account)
}

func (server Server) handleGetAccounts(response http.ResponseWriter, _ *http.Request) error {
	err, accounts := server.store.GetAccounts()
	if err != nil {
		return err
	}
	return sendJSON(response, http.StatusOK, accounts)
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
	vars := mux.Vars(request)
	accountID := vars["id"]
	uuidRegex := regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[1-5][0-9a-fA-F]{3}-[89abAB][0-9a-fA-F]{3}-[0-9a-fA-F]{12}$`)
	if !uuidRegex.MatchString(accountID) {
		return fmt.Errorf("Invalid account ID: %s", accountID)
	}
	return server.store.DeleteAccountByID(accountID)
}

func (server Server) handleTransfer(response http.ResponseWriter, request *http.Request) error {
	return nil
}
