package track

import "github.com/najemi-software/tidal-dl/common"

type Track struct {
	common.BaseTrack
	Artist common.ItemArtist `json:"artist"`
}
