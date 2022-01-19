package transactionservice

import (
  "fmt"
  "log"
  "io/ioutil"
  "sort"
  "strings"
  "net/http"
  "encoding/json"
  
  "gochallenge/transactionservice/model"
)

type TransactionService struct {
  TransactionFile string
  Transactions []model.Transaction
}

func NewTransactionInterface(transactionFile string) *TransactionService {
  return &TransactionService{
    TransactionFile: transactionFile,
  }
}

func (s *TransactionService) GetTransactions(w http.ResponseWriter, r *http.Request) {
  err := s.GetTransactionsFromFile()
  if err != nil {
    w.Write([]byte(err.Error()))
    log.Fatal("Error getting transactions from json file: ", err)
  }

  fmt.Println(s.Transactions)
  transactionDisplays := s.createTransactionDisplays()
  fmt.Println("Current Transactions:")
  s.printTransactionDisplays(transactionDisplays)
  json.NewEncoder(w).Encode(transactionDisplays)
}

func (s *TransactionService)GetTransactionsNewestToOldest(w http.ResponseWriter, r *http.Request) {
  err :=  s.GetTransactionsFromFile()
  if err != nil {
    w.Write([]byte(err.Error()))
    log.Fatal("Error getting transactions from json file: ", err)
  }

  sort.Slice(s.Transactions, func(i, j int) bool {
    return s.Transactions[i].PostedTimeStamp > s.Transactions[j].PostedTimeStamp
  })

  transactionDisplays := s.createTransactionDisplays()
  fmt.Println("Current Transactions:")
  s.printTransactionDisplays(transactionDisplays)
  json.NewEncoder(w).Encode(transactionDisplays)
}

func (s *TransactionService)GetTransactionsFromFile() error {
  transactionFileContent, err := ioutil.ReadFile(s.TransactionFile)
  if err != nil {
    return err
  }

  err = json.Unmarshal(transactionFileContent, &s.Transactions)
  if err != nil {
    return err
  }

  return nil
}

func (s *TransactionService)createTransactionDisplays() []model.TransactionDisplay {
  var transactionDisplays []model.TransactionDisplay

  for _, transaction := range s.Transactions {
    transactionDisplays = append(transactionDisplays, model.TransactionDisplay{
      ID:                   transaction.ID,
      Amount:               transaction.Amount,
	    MessageType:          transaction.MessageType,
	    CreatedAt:            transaction.CreatedAt,
      TransactionID:        transaction.TransactionID,
      PAN:                  s.HidePan(transaction.PAN),
      TransactionCategory:  transaction.TransactionCategory,
      PostedTimeStamp:      transaction.PostedTimeStamp,
      TransactionType:      transaction.TransactionType,
      SendingAccount:       transaction.SendingAccount,
      ReceivingAccount:     transaction.ReceivingAccount,
      TransactionNote:      transaction.TransactionNote,
    })
  }

  return transactionDisplays
}

func (s *TransactionService)HidePan(panInt int) string {
  panStr := fmt.Sprintf("%d", panInt)
  hiddenDigits := strings.Repeat("*", len(panStr)-4)
  lastFourDigits := panStr[len(panStr)-4:]

  return hiddenDigits + lastFourDigits
}

func (s *TransactionService)printTransactionDisplays(transactionDisplays []model.TransactionDisplay) {
  for _, transaction := range transactionDisplays {
    fmt.Printf("Transaction #%d:\n", transaction.ID)
    fmt.Printf("\tAmount: %.2f\n", float32(transaction.Amount)/100.0)
    fmt.Printf("\tMessage Type: %s\n", transaction.MessageType)
    fmt.Printf("\tCreated At: %s\n", transaction.CreatedAt)
    fmt.Printf("\tTransaction ID: %d\n", transaction.TransactionID)
    fmt.Printf("\tPAN: %s\n", transaction.PAN)
    fmt.Printf("\tTransaction Category: %s\n", transaction.TransactionCategory)
    fmt.Printf("\tPosted TimeStamp: %s\n", transaction.PostedTimeStamp)
    fmt.Printf("\tTransaction Type: %s\n", transaction.TransactionType)
    fmt.Printf("\tSending Account: %d\n", transaction.SendingAccount)
    fmt.Printf("\tReceiving Account: %d\n", transaction.ReceivingAccount)
    fmt.Printf("\tTransaction Note: %s\n", transaction.TransactionNote)
  }
}