package common

type BaseVideo struct {
	Id                     int          `json:"id"`
	Title                  string       `json:"title"`
	VolumeNumber           int          `json:"volumeNumber"`
	TrackNumber            int          `json:"trackNumber"`
	ReleaseDate            string       `json:"releaseDate"`
	ImagePath              interface{}  `json:"imagePath"`
	ImageId                string       `json:"imageId"`
	VibrantColor           string       `json:"vibrantColor"`
	Duration               int          `json:"duration"`
	Quality                string       `json:"quality"`
	StreamReady            bool         `json:"streamReady"`
	AdSupportedStreamReady bool         `json:"adSupportedStreamReady"`
	DJReady                bool         `json:"djReady"`
	StemReady              bool         `json:"stemReady"`
	StreamStartDate        string       `json:"streamStartDate"`
	AllowStreaming         bool         `json:"allowStreaming"`
	Explicit               bool         `json:"explicit"`
	Popularity             int          `json:"popularity"`
	Type                   string       `json:"type"`
	AdsUrl                 interface{}  `json:"adsUrl"`
	AdsPrePaywallOnly      bool         `json:"adsPrePaywallOnly"`
	Artists                []ItemArtist `json:"artists"`
	Album                  interface{}  `json:"album"`
}
