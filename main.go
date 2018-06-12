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

var termsArray = []string{}

// initializes this fuzzy logic terms
func init() {
	readTerms()
}

func main() {
	http.HandleFunc("/search", search)
	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	ranks := fuzzy.RankFindFold(strings.ToLower(query), termsArray)
	rankAsJSON, err := json.Marshal(ranks)
	if err != nil {
		fmt.Println(err)
		return
	}
    w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/json")
    w.Write(rankAsJSON)
}

func readTerms() {
	file, err := os.Open("terms-source.txt")
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		term := scanner.Text()
		termsArray = append(termsArray, term)
	}
}