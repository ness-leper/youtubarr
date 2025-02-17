package youtube

import (
	"context"
	"fmt"
	env "github.com/youtubarr/environment"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
)

func Search(query string) [10]string {
	apiKey := env.YoutubeApi().Value

	// Create YouTube service.
	ctx := context.Background()
	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	// Create the search call.
	call := service.Search.List([]string{"snippet"}).
		Q(query).
		Type("video").
		MaxResults(10) // Limit to 5 results for brevity.  Adjust as needed

  // Execute the search.
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}

	// Print the results.
	fmt.Println("Search Results:")
	ids := [10]string{}
	count := 0
	for _, item := range response.Items {
		switch item.Id.Kind {
		case "youtube#video":
			fmt.Printf("Video: %s (%s)\n", item.Snippet.Title, item.Id.VideoId)
			ids[count] = item.Id.VideoId
			count += 1
		}
	}

	return ids
}
