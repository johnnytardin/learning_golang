package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func hora(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora Certa: %s", s)
}

func main() {
	http.HandleFunc("/hora", hora)
	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
