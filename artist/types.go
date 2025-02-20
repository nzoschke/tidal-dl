package artist

import "github.com/najemi-software/tidal-dl/v4/artist_type"

type Role struct {
	CategoryId int    `json:"categoryId"`
	Category   string `json:"category"`
}

type Artist struct {
	Id                         int                      `json:"id"`
	Name                       string                   `json:"name"`
	ArtistTypes                []artist_type.ArtistType `json:"artistTypes"`
	Url                        string                   `json:"url"`
	Picture                    string                   `json:"picture"`
	SelectedAlbumCoverFallback interface{}              `json:"selectedAlbumCoverFallback"`
	Popularity                 int                      `json:"popularity"`
	ArtistRoles                []Role                   `json:"artistRoles"`
	Mixes                      map[string]string        `json:"mixes"`
}

type Bio struct {
	Source      string `json:"source"`
	LastUpdated string `json:"lastUpdated"`
	Text        string `json:"text"`
	Summary     string `json:"summary"`
}

type SimilarArtist struct {
	Artist
	Type         interface{} `json:"type"`
	Banner       interface{} `json:"banner"`
	RelationType string      `json:"relationType"`
}
