## Build

>go build
>./gochallenge

## Running

>go run main.go --transactions "path/to/transaction-json-file"

OR

>go run main.go --transactions "local-json-file"
if the transaction json file is in same directory

OR

if built alreay
> ./gochallenge

## Testing

>go test ./...

OR just tests for transaction service

>go test ./transactionservice/...