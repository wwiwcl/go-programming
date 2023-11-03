package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

// TODO: Please create a struct to include the information of a video

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	// TODO: Get video ID from URL query `v`
	// TODO: Get video information from YouTube API
	// TODO: Parse the JSON response and store the information into a struct
	// TODO: Display the information in an HTML page through `template`
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	value := os.Getenv("YOUTUBE_API_KEY")
	fmt.Fprintf(w, value)
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}