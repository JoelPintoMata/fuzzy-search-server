# Fuzzy search REST server

## Run
```bash
$ go run main.go
```

## Tests
```bash
$ go test -v
```

## Example
Verb: GET
Parameters: 
    source=<source>: reads a source file named `<source>.txt`
    q=<query>: the query
Example: http://localhost:8080/search?source=stations&q=leiden 