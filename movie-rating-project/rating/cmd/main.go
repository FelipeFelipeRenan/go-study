package main

import (
	"log"
	"movieexample/rating/internal/controller/rating"
	httphandler "movieexample/rating/internal/handler/http"
	"movieexample/rating/internal/repository/memory"
	"net/http"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	ctrl := rating.New(repo)
	h := httphandler.New(ctrl)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil{
		panic(err)
	}
}