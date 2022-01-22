package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/woses/websiteAPI/lastfm"
	"github.com/woses/websiteAPI/router"
)

func main() {
	endpoint := "http://ws.audioscrobbler.com/2.0/?method=user.getrecenttracks&user=Wosess&api_key=a03ef2452262e8cc102403ea68648a7b&format=json"
	lfm := lastfm.New(endpoint)
	port := "80"

	fmt.Println("Running on Port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router.NewRouter(lfm)))
}
