package generics

type PaginatedResponse[T any] struct {
	Limit              int `json:"limit"`
	Offset             int `json:"offset"`
	TotalNumberOfItems int `json:"totalNumberOfItems"`
	Items              []T `json:"items"`
}
