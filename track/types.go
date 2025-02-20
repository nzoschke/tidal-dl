package track

import (
	"github.com/najemi-software/tidal-dl/v3/audio_quality"
	"github.com/najemi-software/tidal-dl/v3/common"
	"github.com/najemi-software/tidal-dl/v3/encryption_type"
)

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

type PlaybackInfo struct {
	TrackId            int                        `json:"trackId"`
	AssetPresentation  string                     `json:"assetPresentation"`
	AudioMode          string                     `json:"audioMode"`
	AudioQuality       audio_quality.AudioQuality `json:"audioQuality"`
	ManifestMimeType   string                     `json:"manifestMimeType"`
	ManifestHash       string                     `json:"manifestHash"`
	Manifest           string                     `json:"manifest"`
	AlbumReplayGain    float64                    `json:"albumReplayGain"`
	AlbumPeakAmplitude float64                    `json:"albumPeakAmplitude"`
	TrackReplayGain    float64                    `json:"trackReplayGain"`
	TrackPeakAmplitude float64                    `json:"trackPeakAmplitude"`
	BitDepth           *int                       `json:"bitDepth"`
	SampleRate         *int                       `json:"sampleRate"`
}

type Manifest struct {
	MimeType       string                         `json:"mimeType"`
	Codecs         string                         `json:"codecs"`
	EncryptionType encryption_type.EncryptionType `json:"encryptionType"`
	Urls           []string                       `json:"urls"`
}
