package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/woses/websiteAPI/lastfm"
	"github.com/woses/websiteAPI/models"
)

func NewRouter(lfm lastfm.Lastfm) http.Handler {
	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET"},
		AllowCredentials: true,
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("woses"))
	})

	r.Get("/spotifyRecent", func(w http.ResponseWriter, r *http.Request) {
		data, err := lfm.GetData()
		if err != nil {
			fmt.Errorf("%e", err)
			SendResponse(w, nil, false, "Error while getting Data")
		} else {
			SendResponse(w, data.RecentTracks.Track, true, "")
		}
	})

	return r
}

func SendResponse(w http.ResponseWriter, data interface{}, success bool, er string) {
	response := models.ResponeBuilder{
		Data:    data,
		Success: success,
		Error:   er,
	}
	json, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error"))
	}
	_, err = w.Write([]byte(json))
	if err != nil {
		w.Write([]byte("Error"))
	}
}
