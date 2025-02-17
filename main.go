package main

import (
	"flag"
	"fmt"

	yt "github.com/youtubarr/youtube"
)

func main(){
	query := flag.String("q", "Golang tutorial", "Search query")
  flag.Parse()
  ids := yt.Search(*query)
  fmt.Println(ids)
}
