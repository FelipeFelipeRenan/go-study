package main

import (
	"log"
	"movieexample/movie/internal/controller/movie"
	metadatagateway "movieexample/movie/internal/gateway/metadata/http"
	ratinggateway "movieexample/movie/internal/gateway/rating/http"
	httphandler "movieexample/movie/internal/handler/http"
	"net/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhosy:8081")
	ratingGateway := ratinggateway.New("localhost:8082")

	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil{
		panic(err)
	}
}