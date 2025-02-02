package common

type BaseTrack struct {
	Id                     int               `json:"id"`
	Title                  string            `json:"title"`
	Duration               int               `json:"duration"`
	ReplayGain             float64           `json:"replayGain"`
	Peak                   float64           `json:"peak"`
	AllowStreaming         bool              `json:"allowStreaming"`
	StreamReady            bool              `json:"streamReady"`
	AdSupportedStreamReady bool              `json:"adSupportedStreamReady"`
	DJReady                bool              `json:"djReady"`
	StemReady              bool              `json:"stemReady"`
	StreamStartDate        string            `json:"streamStartDate"`
	PremiumStreamingOnly   bool              `json:"premiumStreamingOnly"`
	TrackNumber            int               `json:"trackNumber"`
	VolumeNumber           int               `json:"volumeNumber"`
	Version                interface{}       `json:"version"`
	Popularity             int               `json:"popularity"`
	Copyright              string            `json:"copyright"`
	BPM                    int               `json:"bpm"`
	Url                    string            `json:"url"`
	ISRC                   string            `json:"isrc"`
	Editable               bool              `json:"editable"`
	Explicit               bool              `json:"explicit"`
	AudioQuality           string            `json:"audioQuality"`
	AudioModes             []string          `json:"audioModes"`
	MediaMetadata          MediaMetadata     `json:"mediaMetadata"`
	Artists                []ItemArtist      `json:"artists"`
	Album                  TrackAlbum        `json:"album"`
	Mixes                  map[string]string `json:"mixes"`
}
