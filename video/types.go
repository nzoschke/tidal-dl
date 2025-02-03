package video

import "github.com/najemi-software/tidal-dl/v2/common"

type Video struct {
	common.BaseVideo
	Artist common.ItemArtist `json:"artist"`
}

type StreamUrl struct {
	Url          string `json:"url"`
	VideoQuality string `json:"videoQuality"`
}
