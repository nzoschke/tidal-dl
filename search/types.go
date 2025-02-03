package search

import (
	"github.com/najemi-software/tidal-dl/v2/artist"
	"github.com/najemi-software/tidal-dl/v2/common"
	"github.com/najemi-software/tidal-dl/v2/generics"
	"github.com/najemi-software/tidal-dl/v2/playlist"
)

// TODO: Make TopHit nicer in subsequent versions

type TopHit struct {
	Value interface{} `json:"value"`
	Type  string      `json:"type"`
}

type ArtistsResponse = generics.PaginatedResponse[artist.Artist]
type AlbumsResponse = generics.PaginatedResponse[common.BaseAlbum]
type PlaylistsResponse = generics.PaginatedResponse[playlist.Playlist]
type TracksResponse = generics.PaginatedResponse[common.BaseTrack]
type VideosResponse = generics.PaginatedResponse[common.BaseVideo]

type Response struct {
	Artists   ArtistsResponse   `json:"artists"`
	Albums    AlbumsResponse    `json:"albums"`
	Playlists PlaylistsResponse `json:"playlists"`
	Tracks    TracksResponse    `json:"tracks"`
	Videos    VideosResponse    `json:"videos"`
	TopHit    TopHit            `json:"topHit"`
}

type Types string

const (
	Artists   Types = "ARTISTS"
	Albums    Types = "ALBUMS"
	Playlists Types = "PLAYLISTS"
	Tracks    Types = "TRACKS"
	Videos    Types = "VIDEOS"
)

func typesToStrings(types []Types) []string {
	var result []string
	for _, t := range types {
		result = append(result, string(t))
	}
	return result
}
