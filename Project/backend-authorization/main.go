package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	http.HandleFunc("/", HelloServer)
	port := os.Getenv("PORT")
	fmt.Println("Listening on port:", port)
	http.ListenAndServe(":"+port, nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path[1:])
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
