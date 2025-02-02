package playlist

type Creator struct {
	Id *int `json:"id,omitempty"`
}

type PromotedArtist struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    string `json:"type"`
	Picture string `json:"picture"`
}

type Playlist struct {
	Uuid            string           `json:"uuid"`
	Title           string           `json:"title"`
	NumberOfTracks  int              `json:"numberOfTracks"`
	NumberOfVideos  int              `json:"numberOfVideos"`
	Creator         Creator          `json:"creator"`
	Description     string           `json:"description"`
	Duration        int              `json:"duration"`
	LastUpdated     string           `json:"lastUpdated"`
	Created         string           `json:"created"`
	Type            string           `json:"type"`
	PublicPlaylist  bool             `json:"publicPlaylist"`
	Url             string           `json:"url"`
	Image           string           `json:"image"`
	Popularity      int              `json:"popularity"`
	SquareImage     string           `json:"squareImage"`
	CustomImageUrl  interface{}      `json:"customImageUrl"`
	PromotedArtists []PromotedArtist `json:"promotedArtists"`
	LastItemAddedAt string           `json:"lastItemAddedAt"`
}
