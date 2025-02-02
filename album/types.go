package album

import (
	"github.com/najemi-software/tidal-dl/common"
)

type Album struct {
	common.BaseAlbum
	Artist common.ItemArtist `json:"artist"`
}
