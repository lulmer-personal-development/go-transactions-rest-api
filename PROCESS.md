## Staring Up

The first thing is actually just getting a handle on my development environment. That is always the first step as nothing can be done until you are familiar with the tools. In the 'real world' I would be using my already installed IDE and the setup I was comforatable with.

My initial journey started with learning the console and shell. I tried running the existing main.go and then using curl to reach out to 'http://localhost:8000/transactions' and observing the returned json and the text printing in the console. After that, I felt comfortable enough to proceed.


## Initial Look
There are a lot of things that need to be done with this code. I noticed early on that the decorators strings on the Transaction structure are incorrect and confusing. Both TransactionID and CreatedAt had the decorator of 'created_at' while the MessageType member had the somewhat confusing name of 'conversation_type'

Looking through the problem I realized there are number of requirements that you would likely not find in the 'real world'. Some highlights:

- Transactions would not be stored in a json file. They would be stored in a database. Either relational or a key/value. Especially not for returning data for REST API endpoints

- I do not believe json files would be included in such a manner as a command line option. The option would likely be an identifier telling the program to look for a file without the user knowing the important files, allowing for important data to be hidden

- There should be a limiting factor to the number of transaction records returned to prevent the response from overloading the memory or network

- Finally, there is no way this one file approach would fly. As it was I had to change the package in go.mod to mymain to get the unit tests working. It should have been set up as a different package with 'go mod init challengepackage' or something to that effect and the functions would go into subdirectories with their own folder

## The design

The design for this was actually pretty straight forward. Parse the option for the transaction file and then assign it to a global that the handlers could use and then set up the handlers.

The handlers themselves were pretty similar. They both follow similar steps:

1) Open the transaction json file

2) Unmarshal the json into the array of transaction structures

3) Do any necessary manipulation of the data. A sort in the case of the sorted endpoints

4) Convert the transactions to a similar structure that I call `TransactionDisplay` which allows for hiding the PAN with *'s except the last 4 digits

5) Marshal json and write the response

## Implmentation steps

1) Read From file

The first step was to have GetTransactions reading from a file. I chose to move the mock data into a file called `transactions.json`. This was rather time consuming task as it turned out.

Once that was complete I added some code to open the file and unmarshal it before into a structure before marshalling again to send the response. I used the standard go library `io/ioutil` for the reading as it is already part of Go.

2) Take In The Option From Command line

The next task was to take in the option from the command line. I decided to use the flags library instead of using looking at os.Args. This looks cleaner to me and allows for some error checking. It also allows for expansion in the future as new options could be added and we would not need to worry about flag order.

Once the value was read in it was simply a matter of assigning the value to a global called `TransactionFile`. Personally this is where I ended up really not liking the one file approach. Globals have a number of issues and I believe a better approach would have been to create an interface for the handlers and then pass in the value to the struct that implemented that interface. I technically could still have done that with the one file approach but it kind of defeats the purpose of an interface with the implementation is in the same file.

Lastly there was the simple matter of changing the file opening in `GetTransactions` to use `TransactionFile` instead of a hard coded value. It was also at this point I decided to create another json file for the purposes of writing unit tests to confirm its operations with the files given to it.

3) Hiding the PAN

Next was changing the output to hide most of the PAN values behind *'s. This unfortunately necessitated creating a new stuct for outputting that I called `TransactionDisplay`. The structures are identical except that PAN is a string as opposed to an int. I was unable to find anyway to write an int with digits hidden in such a fashion. This necessitated a conversion function that turned `Transaction` in a `TransactionDisplay` and called a separate `HidePan` function to turned the PAN int into a PAN string.

The `HidePan` function was actually extremely simple. Convert the int to a string (I chose `fmt.Sprintf` so as not to import another go library) then build a new string using the PAN string length and final 4 characters.

4) The Sorted Transaction Endpoint

Next up was the endpoint to return transactions ordered by a descending `posted_timestamp`. As GO already has a library specifically for sorting I decided to utilize that instead of making my own custom sorting function. It ended up being something of a carbon copy of `GetTransactions` and I could abstract out some of the code but time started to become a factor.

5) Unit Tests

I decided not to make unit tests for `CreateTransactionDisplays` and `PrintTransactionDisplays` functions as they really did not do anything in terms of manipulation. I did make tests for `GetTransactionsFromFile` and `HidePan` to confirm their functionality.

I also made tests for the handlers. I could not remember how to unit test handlers off the top of my head so I had to look it up. Thank you `https://www.thepolyglotdeveloper.com/2017/02/unit-testing-golang-application-includes-http` for the help on this one.

It should be noted that I had to change the package name in the go.mod file to get this working. I chose to call it `gochallenge`. Simply calling it main is not good practice and makes it difficult to split up functionality into files

6) Splitting the Functionality between multiple files

Initially I had convinced myself that I needed to keep as much as possible inside the main.go file. But after renaming the module and making a main_test.go file, I decided to just say screw it and go all in on multiple files. I moved the structs into a model directory containing model.go and moved the handler functions and there supports into the service directory. I made an interface for the main function to utilize in setting up the endpoints and then made an implementation of the service called TransactionService with a creation function that takes in the transaction file name. I then moved all the unit tests into the service file. I feel this makes the whole program a lot more readable as the one file solution was actually quite messy. This is more how such a program would be completed in the 'real world' as well.