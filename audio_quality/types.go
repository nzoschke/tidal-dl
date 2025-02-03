package audio_quality

type AudioQuality string

const (
	Normal AudioQuality = "LOW"
	High   AudioQuality = "HIGH"
	HiFi   AudioQuality = "LOSSLESS"
	Master AudioQuality = "HI_RES"
	Max    AudioQuality = "HI_RES_LOSSLESS"
)
