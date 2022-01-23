package lastfm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/woses/websiteAPI/models"
)

type Lastfm struct {
	Apikey   string
	Endpoint string
}

func New(endpoint string) Lastfm {
	return Lastfm{
		Endpoint: endpoint,
	}
}

func (l Lastfm) GetData() (r models.Response, e error) {
	resp, err := http.Get(l.Endpoint)
	if err != nil {
		e = err
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		e = err
		return
	}

	var data models.Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		e = err
		return
	}

	for _, track := range data.RecentTracks.Track {
		track.Link = strings.Replace(track.Link, "\\", "", -1)
		for _, image := range track.Image {
			image.Text = strings.Replace(image.Text, "\\", "", -1)
		}
	}

	r = data
	return
}
