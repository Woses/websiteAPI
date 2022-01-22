package models

type Response struct {
	RecentTracks struct {
		Track []struct {
			Arist struct {
				Text string `json:"#text"`
			} `json:"artist"`
			Name string `json:"name"`
			Attr struct {
				NowPlaying string `json:"nowplaying"`
			} `json:"@attr"`
			Date struct {
				Uts string `json:"uts"`
			} `json:"date"`
		} `json:"track"`
	} `json:"recenttracks"`
}

type ResponeBuilder struct {
	Data    interface{}
	Success bool
	Error   string
}
