package transactionservice

import (
  "fmt"
  "testing"
  "reflect"
  "encoding/json"
  "net/http"
  "net/http/httptest"

  "github.com/gorilla/mux"

  "gochallenge/transactionservice/model"
)

var expectedFromFile []model.Transaction = []model.Transaction{
  {
    ID: 1,
    Amount: 1000,
    MessageType: "Debit",
    CreatedAt: "2020-06-11T19:11:24+00:00",
    TransactionID: 101,
    PAN: 1234567890123456,
    TransactionCategory: "Automotive",
    PostedTimeStamp: "2020-06-11T19:11:24+00:00",
    TransactionType: "POS", 
    SendingAccount: 12345, 
    ReceivingAccount: 1234567, 
    TransactionNote: "Test Business Inc.",
  },
  {
    ID: 2,
    Amount: 2000,
    MessageType: "Credit",
    CreatedAt: "2021-11-11T23:15:34+00:00",
    TransactionID: 256,
    PAN: 9876543210987654,
    TransactionCategory: "Health Services",
    PostedTimeStamp: "2021-11-11T23:15:34+00:00",
    TransactionType: "POS", 
    SendingAccount: 67890, 
    ReceivingAccount: 7891234, 
    TransactionNote: "Another Business",
  },
  {
    ID: 3,
    Amount: 3000,
    MessageType: "Debit",
    CreatedAt: "2019-09-23T06:45:12+00:00",
    TransactionID: 420,
    PAN: 1357924680654321,
    TransactionCategory: "Food And Beverage",
    PostedTimeStamp: "2019-09-23T06:45:12+00:00",
    TransactionType: "POS", 
    SendingAccount: 54321, 
    ReceivingAccount: 1357924, 
    TransactionNote: "Business MK. III",
  },
}

var expectedDisplays []model.TransactionDisplay = []model.TransactionDisplay{
    {
    ID: 1,
    Amount: 1000,
    MessageType: "Debit",
    CreatedAt: "2020-06-11T19:11:24+00:00",
    TransactionID: 101,
    PAN: "************3456",
    TransactionCategory: "Automotive",
    PostedTimeStamp: "2020-06-11T19:11:24+00:00",
    TransactionType: "POS", 
    SendingAccount: 12345, 
    ReceivingAccount: 1234567, 
    TransactionNote: "Test Business Inc.",
  },
  {
    ID: 2,
    Amount: 2000,
    MessageType: "Credit",
    CreatedAt: "2021-11-11T23:15:34+00:00",
    TransactionID: 256,
    PAN: "************7654",
    TransactionCategory: "Health Services",
    PostedTimeStamp: "2021-11-11T23:15:34+00:00",
    TransactionType: "POS", 
    SendingAccount: 67890, 
    ReceivingAccount: 7891234, 
    TransactionNote: "Another Business",
  },
  {
    ID: 3,
    Amount: 3000,
    MessageType: "Debit",
    CreatedAt: "2019-09-23T06:45:12+00:00",
    TransactionID: 420,
    PAN: "************4321",
    TransactionCategory: "Food And Beverage",
    PostedTimeStamp: "2019-09-23T06:45:12+00:00",
    TransactionType: "POS", 
    SendingAccount: 54321, 
    ReceivingAccount: 1357924, 
    TransactionNote: "Business MK. III",
  },
}

var expectedSortedDisplays []model.TransactionDisplay = []model.TransactionDisplay{
  {
    ID: 2,
    Amount: 2000,
    MessageType: "Credit",
    CreatedAt: "2021-11-11T23:15:34+00:00",
    TransactionID: 256,
    PAN: "************7654",
    TransactionCategory: "Health Services",
    PostedTimeStamp: "2021-11-11T23:15:34+00:00",
    TransactionType: "POS", 
    SendingAccount: 67890, 
    ReceivingAccount: 7891234, 
    TransactionNote: "Another Business",
  },  
  {
    ID: 1,
    Amount: 1000,
    MessageType: "Debit",
    CreatedAt: "2020-06-11T19:11:24+00:00",
    TransactionID: 101,
    PAN: "************3456",
    TransactionCategory: "Automotive",
    PostedTimeStamp: "2020-06-11T19:11:24+00:00",
    TransactionType: "POS", 
    SendingAccount: 12345, 
    ReceivingAccount: 1234567, 
    TransactionNote: "Test Business Inc.",
  },
  {
    ID: 3,
    Amount: 3000,
    MessageType: "Debit",
    CreatedAt: "2019-09-23T06:45:12+00:00",
    TransactionID: 420,
    PAN: "************4321",
    TransactionCategory: "Food And Beverage",
    PostedTimeStamp: "2019-09-23T06:45:12+00:00",
    TransactionType: "POS", 
    SendingAccount: 54321, 
    ReceivingAccount: 1357924, 
    TransactionNote: "Business MK. III",
  },
}

func setup() *TransactionService {
  return NewTransactionInterface("test_transactions.json")
}

func Test_GetTransactionsFromFile(t *testing.T) {
  ts := setup()
  err := ts.GetTransactionsFromFile()
  
  if err != nil {
    t.Error("Expected not error. Got: ", err.Error())
  } else {
    t.Log("Received no error")
  }

  for i, transaction := range ts.Transactions {
    if reflect.DeepEqual(transaction, expectedFromFile[i]) {
      t.Log("Transactions are as expected")
    } else {
      t.Error("Transactions not as expected. Expected: ", expectedFromFile[i], " Actual: ", transaction)
    }
  }
}

func Test_HidePan(t *testing.T) {
  ts := setup()

  pan := 1234567890123456
  expected := "************3456"

  actual := ts.HidePan(pan)
  if actual != expected {
    t.Error("Output Not Expected. Expected: ", expected, "Got: ", actual)
  } else {
    t.Log("HidePan works as expected")
  }
}

func makeTestRouter() *mux.Router {
  ts := setup()

  testRouter := mux.NewRouter()
  testRouter.HandleFunc("/transactions", ts.GetTransactions).Methods("GET")
  testRouter.HandleFunc("/transactions/sorted", ts. GetTransactionsNewestToOldest).Methods("GET")
  return testRouter
}

func parseResponseJson(resp *httptest.ResponseRecorder) []model.TransactionDisplay{
  var responseTransactionDisplays []model.TransactionDisplay
  err := json.NewDecoder(resp.Body).Decode(&responseTransactionDisplays)
  if err != nil {
    fmt.Println("Couldn't decode")
  }
  return responseTransactionDisplays;
}

func Test_GetTransactions(t *testing.T) {
  testRouter := makeTestRouter()
  testRequest, err := http.NewRequest("GET", "/transactions", nil)
  if err != nil {
    t.Error("Error making request")
    return
  }
  testResponse := httptest.NewRecorder()
  testRouter.ServeHTTP(testResponse, testRequest)
  if testResponse.Code != 200 {
    t.Error("Test GetTransactions: received improper response code")
  }

  actual := parseResponseJson(testResponse)
  if !reflect.DeepEqual(expectedDisplays, actual) {
    t.Error(actual)
    t.Error("Error in returned transactions")
  }
}

func Test_GetTransactionsNewestToOldest(t *testing.T) {
  testRouter := makeTestRouter()
  testRequest, err := http.NewRequest("GET", "/transactions/sorted", nil)
  if err != nil {
    t.Error("Error making request")
    return
  }
  testResponse := httptest.NewRecorder()

  testRouter.ServeHTTP(testResponse, testRequest)
  if testResponse.Code != 200 {
    t.Error("Test GetTransactions: received improper response code")
  }

  actual := parseResponseJson(testResponse)
  if !reflect.DeepEqual(expectedSortedDisplays, actual) {
      t.Error("Error in transactions")
  }
}