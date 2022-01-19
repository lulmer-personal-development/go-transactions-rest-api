package transactionservice

import "net/http"

type TransactionInterface interface {
  GetTransactions(w http.ResponseWriter, r *http.Request)
  GetTransactionsNewestToOldest(w http.ResponseWriter, r *http.Request)
}