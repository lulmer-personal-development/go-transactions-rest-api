package main

import (
    "fmt"
    "log"
    "os"
    "flag"
    "net/http"

    "github.com/gorilla/mux"

    "gochallenge/transactionservice/service"
    )

func main() {
  var transactionFile string

  flag.StringVar(&transactionFile, "transactions", "transactions.json", "Specify file of transactions.")
  flag.Parse()

  if len(transactionFile) == 0 {
    fmt.Println("Usage: main.go --transactions <transaction/file>")
    os.Exit(1)
  }

  ts := transactionservice.NewTransactionInterface(transactionFile)

  router := mux.NewRouter().StrictSlash(true)
  router.HandleFunc("/transactions", ts.GetTransactions).Methods("GET")
  router.HandleFunc("/transactions/sorted", ts.GetTransactionsNewestToOldest).Methods("GET")
  fmt.Println("Serving transactions on port 8000")
  log.Fatal(http.ListenAndServe(":8000", router))
}
