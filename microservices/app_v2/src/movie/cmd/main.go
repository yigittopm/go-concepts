package main

import (
	"app_v2/src/movie/internal/controller/movie"
	metadatagateway "app_v2/src/movie/internal/gateway/metadata/http"
	ratinggateway "app_v2/src/movie/internal/gateway/rating/http"
	httphandler "app_v2/src/movie/internal/handler/http"
	"log"
	"net/http"
)

func main() {
	log.Println("Starting the movie service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := movie.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)

	http.Handle("/movie", http.HandlerFunc(h.GetMovieDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
