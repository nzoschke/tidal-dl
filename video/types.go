package video

import "github.com/najemi-software/tidal-dl/common"

type Video struct {
	common.BaseVideo
	Artist common.ItemArtist `json:"artist"`
}
