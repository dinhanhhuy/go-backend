package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func page() string {
	workDir := os.Getenv("WORK_DIR")
	dat, _ := os.ReadFile(workDir + "/" + "index.html")
	return string(dat)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	pageData := page()
	fmt.Fprintf(w, pageData)
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	id := os.Getenv("MOCKER_ID")
	fmt.Println("ID: " + id)
	handleRequests()
}
