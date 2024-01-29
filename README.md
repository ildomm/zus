[![Go Report Card](https://goreportcard.com/badge/github.com/ildomm/zus?cache=v1)]

# Title: ZUS

## Description
This API allows you to perform actions related to tokens and their hashes. Here's a brief overview of the available endpoints:

### API Endpoints

#### Create a New Hash
- Endpoint: POST /hash
- Description: Generates a new hash for a given token.
- Parameters:
  - token (required): The token for which a hash will be generated.
- Responses:
  - 200: The operation was successful. Returns the generated hash and the creation timestamp.
  Default: An error occurred. Returns the error code and message.
  
#### Get All Received Hashes
- Endpoint: GET /hashes
- Description: Retrieves all the received hashes.
- Responses:
  - 200: The operation was successful. Returns a list of tokens along with their associated hashes and creation timestamps.
Default: An error occurred. Returns the error code and message.

#### Get Hash Info
- Endpoint: GET /hashes/{id}
- Description: Retrieves information about a specific hash.
- Parameters:
  - id (required): The ID of the hash for which information is requested.
- Responses:
  - 200: The operation was successful. Returns details about the token, including its associated hash and creation timestamp.
Default: An error occurred. Returns the error code and message.

        
## Prerequisites

* Go 1.9+
* PostgreSQL 9.6+
* GNU Make 3.81+

## Local installation

```
make deps
make setupdb
make build
```

## Configuration

Configuration:
```
cp config.yml.example config.yml
```

Database configuration:
```
make setupdb
```

## Validation

Testing:
```
make test
```

## Running

Locally:
```
make build
./build/zus
```
