# Fuzzy search REST server

## Run
```bash
$ go run main.go
```

## Tests
```bash
$ go test -v
```

## Endpoints
Verb | URI                         | Parameters | Example
---------------------------------------------------------------------------------------
GET  | <server>:<port>/convert/csv | q=<query>  | http://localhost:8080/search?q=leiden 