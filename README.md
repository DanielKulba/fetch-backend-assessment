# Fetch Receipt Processor Assessment

## Running the code

Run the following in the receipt-processor-challenge directory:

```sh
go run main.go
```

This will expose port 8080, allowing the user to make requests to `localhost:8080/receipts/...` using Postman or curl for testing.

## Project Structure

- main.go
  - Creater router using gin and configure paths/handlers
  - Initialize mock database
- controllers/handlers.go
  - Contains handler functions that are responsible for
- db/mockDatabase.go
  - An in-memory data store (map) that contains receipts
  - Key: unique id
  - Value: Receipt object containing receipt information
  - Provides methods for saving and retrieving receipts
- models/models.go
  - Contains definitions of structs used throughout the code
  - These are used for parsing the JSON in request as well as building responses
- tools/calculate.go
  - Used for calculating the point value of a receipt
  - Most rules are broken out into helper methods

## Unit Testing

Use the following command in the receipt-processor-challenge directory to execute unit tests:

```sh
go test ./...
```

For the sake of brevity, unit tests are only provided for the tools and db packages. The bulk of the testing focuses on the processing of receipts and the implementation of the provided rules for calculatting points. In a production environment, it would be wise to include unit tests for all packages and have more complete code coverage.

## Using Go

Before this assessment, I did not have any prior experience using Go. Despite this, I thought this would be a good opportunity to give it a shot and challenge myself. In doing so, I tried my best to adhere to Go's best practices, project structure, etc. Go definetely has it quirks, but I was able to pick it up quickly with the help of [this great tutorial](https://www.youtube.com/watch?v=8uiZC0l4Ajw). I was pleasantly surprised with Go's simple yet effective syntax as well as the open source nature of external libraries. I could definetely see myself using and mastering Go in the future, whether that is at Fetch or for my own personal projects!
