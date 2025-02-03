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

type PlaybackInfo struct {
	VideoId           int    `json:"videoId"`
	StreamType        string `json:"streamType"`
	AssetPresentation string `json:"assetPresentation"`
	VideoQuality      string `json:"videoQuality"`
	ManifestMimeType  string `json:"manifestMimeType"`
	ManifestHash      string `json:"manifestHash"`
	Manifest          string `json:"manifest"`
}

type Manifest struct {
	MimeType string   `json:"mimeType"`
	Urls     []string `json:"urls"`
}
