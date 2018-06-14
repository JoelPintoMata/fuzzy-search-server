package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/renstrom/fuzzysearch/fuzzy"
)

var port string = os.Getenv("PORT")

var termsMap = map[string][]string{}

func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":" + port, nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	source := r.URL.Query().Get("source")
	if termsMap[source] == nil {
		readSource(source)
	}
	ranks := fuzzy.RankFindFold(strings.ToLower(query), termsMap[source])
	rankAsJSON, err := json.Marshal(ranks)
	if err != nil {
		fmt.Println(err)
		return
	}
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    w.Write(rankAsJSON)
}

func readSource(source string) {
	file, err := os.Open(source + ".txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		term := scanner.Text()
		termsMap[source] = append(termsMap[source], term)
	}
}