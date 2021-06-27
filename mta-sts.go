package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 - not found", http.StatusNotFound)
}
func stsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "version: STSv1")
	fmt.Fprintln(w, "mode:", strings.TrimSpace(os.Getenv("STS_MODE")))
	fmt.Fprintln(w, "max_age:", strings.TrimSpace(os.Getenv("STS_MAX_AGE")))

	servers := strings.Split(os.Getenv("STS_SERVERS"), ",")
	for i := range servers {
		fmt.Fprintln(w, "mx:", strings.TrimSpace(servers[i]))
	}
}

func main() {
	http.HandleFunc("/", notFound)
	http.HandleFunc("/.well-known/mta-sts.txt", stsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
