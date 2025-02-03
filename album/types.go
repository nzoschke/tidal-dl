package album

import (
	"github.com/najemi-software/tidal-dl/v2/common"
)

type Album struct {
	common.BaseAlbum
	Artist common.ItemArtist `json:"artist"`
}
