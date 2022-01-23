package models

type Response struct {
	RecentTracks struct {
		Track []struct {
			Arist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Image []struct {
				Size string `json:"size"`
				Text string `json:"#text"`
			} `json:"image"`
			Name string `json:"name"`
			Attr struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr"`
			Date struct {
				Uts string `json:"uts"`
			} `json:"date"`
			Link string `json:"url"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type ResponeBuilder struct {
	Data    interface{}
	Success bool
	Error   string
}
