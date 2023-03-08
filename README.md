# Prime Checker API - Developed by Sinisa Mitrovic

This is a simple REST API implemented using the Fiber Go framework
that checks whether a given list of integers are prime or not.

## API Endpoints

POST /prime

Accepts a JSON array of integers in the request body, and returns a JSON array of booleans indicating whether each number is prime or not. If any element in the input array is not a valid integer, the API returns an error message in the response.

### Request Format

```json
[2, 3, 4, 5]
```

### Response Format

```json
[true, true, false, true]
```

If any element in the input array is not a valid integer, the API returns an error message in the response:

### Request Format

```json
[2, 3, "nan", 5]
```

### Response Format

```json
{
  "error": "Element on index 2 is not valid"
}
```

## Error Codes

400 Bad Request: if the request body is not a valid JSON array of integers.
500 Internal Server Error: if there is an error in the server while processing the request.

## Local Development

To run the application locally, follow these steps:

Clone the repository and navigate to the project directory:

git clone https://github.com/mitrovicsinisaa/prime.git
cd prime

Install the dependencies:
go mod download

Run the application (on port 8080):

```
go run main.go --address=:8080
```

This will start the application on http://localhost:8080.

## Testing

To run the tests for the application, use the following command:

go test ./...
