package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
)

// TODO: Please create a struct to include the information of a video
type video_info struct{
	id	string
	date	string
	title	string
	channel	string
	like_count	int
	view_count	int
	comment_count	int
}

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
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	video := video_info{id: r.URL.Query().Get("v"), date: "", title: "", channel: "", like_count: 0, view_count: 0, comment_count: 0}
	if video.id == "" {
		http.Error(w, "Missing 'v' query parameter (video ID)", http.StatusBadRequest)
		return
	}
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		http.Error(w, "Error creating YouTube client", http.StatusInternalServerError)
		return
	}
	call := service.Videos.List([]string{"snippet", "statistics"}).Id(video.id)
	response, err := call.Do()
	if err != nil {
		http.Error(w, "Error calling YouTube API", http.StatusInternalServerError)
		return
	}
	publishedAt := response.Items[0].Snippet.PublishedAt
	title := response.Items[0].Snippet.Title
	channelTitle := response.Items[0].Snippet.ChannelTitle
	likeCount := response.Items[0].Statistics.LikeCount
	viewCount := response.Items[0].Statistics.ViewCount
	commentCount := response.Items[0].Statistics.CommentCount
	video.date = publishedAt
	video.title = title
	video.channel = channelTitle
	video.like_count = int(likeCount)
	video.view_count = int(viewCount)
	video.comment_count = int(commentCount)
	fmt.Fprintf(w, "Published Date: %s\nTitle: %s\nChannel: %s\nLikes: %d\nViews: %d\nComments: %d",
		video.date, video.title, video.channel, video.like_count, video.view_count, video.comment_count)
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}