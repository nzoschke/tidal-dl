package artist

type Types string

const (
	ARTIST Types = "ARTIST"
)

type Role struct {
	CategoryId int    `json:"categoryId"`
	Category   string `json:"category"`
}

type Artist struct {
	Id                         int               `json:"id"`
	Name                       string            `json:"name"`
	ArtistTypes                []string          `json:"artistTypes"`
	Url                        string            `json:"url"`
	Picture                    string            `json:"picture"`
	SelectedAlbumCoverFallback interface{}       `json:"selectedAlbumCoverFallback"`
	Popularity                 int               `json:"popularity"`
	ArtistRoles                []Role            `json:"artistRoles"`
	Mixes                      map[string]string `json:"mixes"`
}
