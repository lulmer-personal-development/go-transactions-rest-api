package model

type Transaction struct {
  ID             int `json:"id"`
  Amount         int `json:"amount"`
	MessageType    string `json:"message_type"`
	CreatedAt      string `json:"created_at"`
  TransactionID  int `json:"transaction_id"`
  PAN            int `json:"pan"`
  TransactionCategory string `json:"transaction_category"`
  PostedTimeStamp string `json:"posted_timestamp"`
  TransactionType string `json:"transaction_type"`
  SendingAccount  int   `json:"sending_account"`
  ReceivingAccount int `json:"receiving_account"`
  TransactionNote string `json:"transaction_note"`
}

// a struct for outputting the json in a manner that can format its display
type TransactionDisplay struct {
  ID             int `json:"id"`
  Amount         int `json:"amount"`
	MessageType    string `json:"message_type"`
	CreatedAt      string `json:"created_at"`
  TransactionID  int `json:"transaction_id"`
  PAN            string `json:"pan"`
  TransactionCategory string `json:"transaction_category"`
  PostedTimeStamp string `json:"posted_timestamp"`
  TransactionType string `json:"transaction_type"`
  SendingAccount  int   `json:"sending_account"`
  ReceivingAccount int `json:"receiving_account"`
  TransactionNote string `json:"transaction_note"`
}
