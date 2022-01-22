package lastfm

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

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
	r = data
	return
}
