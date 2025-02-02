package common

type TrackAlbum struct {
	Id           int         `json:"id"`
	Title        string      `json:"title"`
	Cover        string      `json:"cover"`
	VibrantColor string      `json:"vibrantColor"`
	VideoCover   interface{} `json:"videoCover"`
}
