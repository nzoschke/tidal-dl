package user

type User struct {
	Id        int     `json:"id"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
}
