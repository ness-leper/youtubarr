package youtube

import (
	"context"
	"flag"
	"fmt"
	"log"
	env "github.com/youtubarr/environment"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func Test() {
	query := flag.String("q", "Golang tutorial", "Search query")
	apiKey := env.YoutubeApi().Value
	flag.Parse()

	// Create YouTube service.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// Create the search call.
	call := service.Search.List([]string{"snippet"}).
		Q(*query).
		MaxResults(10) // Limit to 5 results for brevity.  Adjust as needed

		// Execute the search.
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Print the results.
	fmt.Println("Search Results:")
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			fmt.Printf("Video: %s (%s)\n", item.Snippet.Title, item.Id.VideoId)
		case "youtube#channel":
			fmt.Printf("Channel: %s (%s)\n", item.Snippet.Title, item.Id.ChannelId)
		case "youtube#playlist":
			fmt.Printf("Playlist: %s (%s)\n", item.Snippet.Title, item.Id.PlaylistId)
		}
		fmt.Println("------------------------------------")
	}

}
