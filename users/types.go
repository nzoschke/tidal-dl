package users

type User struct {
	Id        int         `json:"id"`
	FirstName interface{} `json:"firstName"`
	LastName  interface{} `json:"lastName"`
}
