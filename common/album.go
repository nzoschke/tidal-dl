package common

type BaseAlbum struct {
	Id                     int           `json:"id"`
	Title                  string        `json:"title"`
	Duration               int           `json:"duration"`
	StreamReady            bool          `json:"streamReady"`
	AdSupportedStreamReady bool          `json:"adSupportedStreamReady"`
	DJReady                bool          `json:"djReady"`
	StemReady              bool          `json:"stemReady"`
	StreamStartDate        string        `json:"streamStartDate"`
	AllowStreaming         bool          `json:"allowStreaming"`
	PremiumStreamingOnly   bool          `json:"premiumStreamingOnly"`
	NumberOfTracks         int           `json:"numberOfTracks"`
	NumberOfVideos         int           `json:"numberOfVideos"`
	NumberOfVolumes        int           `json:"numberOfVolumes"`
	ReleaseDate            string        `json:"releaseDate"`
	Copyright              string        `json:"copyright"`
	Type                   string        `json:"type"`
	Version                interface{}   `json:"version"`
	Url                    string        `json:"url"`
	Cover                  string        `json:"cover"`
	VibrantColor           string        `json:"vibrantColor"`
	VideoCover             interface{}   `json:"videoCover"`
	Explicit               bool          `json:"explicit"`
	Upc                    string        `json:"upc"`
	Popularity             int           `json:"popularity"`
	AudioQuality           string        `json:"audioQuality"`
	AudioModes             []string      `json:"audioModes"`
	MediaMetadata          MediaMetadata `json:"mediaMetadata"`
	Artists                []ItemArtist  `json:"artists"`
}
