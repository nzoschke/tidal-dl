package auth

type LoginLink struct {
	DeviceCode              string `json:"deviceCode"`
	UserCode                string `json:"userCode"`
	VerificationUri         string `json:"verificationUri"`
	VerificationUriComplete string `json:"verificationUriComplete"`
	ExpiresIn               int    `json:"expiresIn"`
	Interval                int    `json:"interval"`
}

type User struct {
	UserId             int     `json:"userId"`
	Email              string  `json:"email"`
	CountryCode        string  `json:"countryCode"`
	FullName           *string `json:"fullName"`
	FirstName          *string `json:"firstName"`
	LastName           *string `json:"lastName"`
	NickName           *string `json:"nickName"`
	Username           string  `json:"username"`
	Address            *string `json:"address"`
	PostalCode         *string `json:"postalcode"`
	UseState           *string `json:"useState"`
	PhoneNumber        *string `json:"phoneNumber"`
	Birthday           *int64  `json:"birthday"`
	ChannelId          int     `json:"channelId"`
	ParentId           int     `json:"parentId"`
	AcceptedEULA       bool    `json:"acceptedEULA"`
	Created            int     `json:"created"`
	Updated            int     `json:"updated"`
	FacebookUId        *int    `json:"facebookUId"`
	AppleUId           *string `json:"appleUId"`
	GoogleUId          *int    `json:"googleUId"`
	AccountLinkCreated bool    `json:"accountLinkCreated"`
	EmailVerified      bool    `json:"emailVerified"`
	NewUser            bool    `json:"newUser"`
}

type LoginLinkError struct {
	Status           int    `json:"status"`
	ErrorType        string `json:"error"`
	SubStatus        int    `json:"sub_status"`
	ErrorDescription string `json:"error_description"`
}

func (e *LoginLinkError) Error() string {
	return e.ErrorDescription
}

type GrantResponse struct {
	Scope        string `json:"scope"`
	User         User   `json:"user"`
	ClientName   string `json:"clientName"`
	TokenType    string `json:"token_type"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	UserId       int    `json:"user_id"`
}
