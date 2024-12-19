package main

import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

type Account struct {
    ID      string  `json:"id"`
    Owner   string  `json:"owner"`
    Balance float64 `json:"balance"`
}

func main() {
    r := mux.NewRouter()
    
    // Register routes
    r.HandleFunc("/api/accounts", getAccounts).Methods("GET")
    r.HandleFunc("/api/accounts/{id}", getAccount).Methods("GET")
    
    log.Printf("Account Service starting on :8081")
    if err := http.ListenAndServe(":8081", r); err != nil {
        log.Fatal(err)
    }
}

func getAccounts(w http.ResponseWriter, r *http.Request) {
    accounts := []Account{
        {ID: "1", Owner: "John Doe", Balance: 1000.00},
    }
    json.NewEncoder(w).Encode(accounts)
}

func getAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    account := Account{ID: vars["id"], Owner: "John Doe", Balance: 1000.00}
    json.NewEncoder(w).Encode(account)
}
