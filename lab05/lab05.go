package main

import (
	"log"
	"net/http"
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"
	"html/template"
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

func getTemplatePath(filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/" + filename
}

func errorPage(w http.ResponseWriter) {
	errtmpl := template.Must(template.ParseFiles(getTemplatePath("error.html")))
	errtmpl.Execute(w, nil)
}

func YouTubePage(w http.ResponseWriter, r *http.Request) {
	// TODO: Get API token from .env file
	// TODO: Get video ID from URL query `v`
	// TODO: Get video information from YouTube API
	// TODO: Parse the JSON response and store the information into a struct
	// TODO: Display the information in an HTML page through `template`
	err := godotenv.Load(".env")
	if err != nil {
		errorPage(w)
		return
	}
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	videoID := r.URL.Query().Get("v")
	if videoID == "" {
		errorPage(w)
		return
	}
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		errorPage(w)
		return
	}
	call := service.Videos.List([]string{"snippet", "statistics"}).Id(videoID)
	response, err := call.Do()
	if err != nil {
		errorPage(w)
		return
	}
	publishedAt := response.Items[0].Snippet.PublishedAt
	title := response.Items[0].Snippet.Title
	channelTitle := response.Items[0].Snippet.ChannelTitle
	likeCount := response.Items[0].Statistics.LikeCount
	viewCount := response.Items[0].Statistics.ViewCount
	commentCount := response.Items[0].Statistics.CommentCount
	video := video_info{
		id:            videoID,
		date:          publishedAt,
		title:         title,
		channel:       channelTitle,
		like_count:    int(likeCount),
		view_count:    int(viewCount),
		comment_count: int(commentCount),
	}
	fmt.Fprintf(w, "Published Date: %s\nTitle: %s\nChannel: %s\nLikes: %d\nViews: %d\nComments: %d",
		video.date, video.title, video.channel, video.like_count, video.view_count, video.comment_count)
}

func main() {
	http.HandleFunc("/", YouTubePage)
	log.Fatal(http.ListenAndServe(":8085", nil))
}