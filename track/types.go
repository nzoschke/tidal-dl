package track

import "github.com/najemi-software/tidal-dl/v2/common"

type Track struct {
	common.BaseTrack
	Artist common.ItemArtist `json:"artist"`
}

type StreamUrl struct {
	Url                   string `json:"url"`
	TrackId               int    `json:"trackId"`
	PlayTimeLeftInMinutes int    `json:"playTimeLeftInMinutes"`
	SoundQuality          string `json:"soundQuality"`
	EncryptionKey         string `json:"encryptionKey"`
	Codec                 string `json:"codec"`
}
