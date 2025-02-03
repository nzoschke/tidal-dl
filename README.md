# Welcome to tidal-dl

A library for interacting with the tidal api, without the need for registering an app.

Authentication works via a device code, the API will show up as an android automotive device when connected to the TIDAL
account.

## Installation

```bash
go get github.com/najemi-software/tidal-dl/v2@latest
```

## Authentication

You can use our prebuilt functions to authenticate, or you can write your own one if you require custom functionality
for interactive login.

### Interactive login

Interactive login gives the user a link with a short code, when the user visits the link, they are prompted to log in,
the code will be pre-entered. Once the user continues, access will automatically be granted. The software will have to
poll an endpoint to check the status of the login.

#### Using our interactive login

```go
package main

import (
	"fmt"
	"github.com/najemi-software/tidal-dl/v2"
	"github.com/najemi-software/tidal-dl/v2/credentials"
)

func main() {
	err := tidal_dl.InteractiveLogin()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(credentials.AccessToken)
}
```

#### Writing your own interactive login
Here is our interactive login function, you can use it as a base for your own implementation.
```go
package main

import (
	"errors"
	"fmt"
	"github.com/najemi-software/tidal-dl/v2/auth"
	"github.com/najemi-software/tidal-dl/v2/credentials"
	"github.com/najemi-software/tidal-dl/v2/session"
	"time"
)

func main() {
	link, err := auth.GetLoginLink()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Visit the following link to login: https://" + link.VerificationUriComplete)

	startTime := time.Now()
	for startTime.Add(time.Second * time.Duration(link.ExpiresIn-3)).After(time.Now()) {
		status, err := auth.GetLoginLinkStatus(link.DeviceCode)
		if err != nil {
			var loginLinkError *auth.LoginLinkError
			if !errors.As(err, &loginLinkError) {
				panic(err.Error())
			}

			if loginLinkError.ErrorType == "authorization_pending" {
				time.Sleep(time.Second * time.Duration(link.Interval))
				continue
			}

			panic(err.Error())
		}

		credentials.TokenType = status.TokenType
		credentials.AccessToken = status.AccessToken
		credentials.RefreshToken = status.RefreshToken
		credentials.Expires = time.Now().Add(time.Second * time.Duration(status.ExpiresIn))

		loginSession, err := session.Get()
		if err != nil {
			panic(err.Error())
		}

		credentials.SessionId = loginSession.SessionId
		credentials.CountryCode = loginSession.CountryCode

		return
	}

	panic("Login link expired")
}
```

### Login with a refresh token

A refresh token holds enough weight that it can be used for login by itself, if you are storing sessions, this would be preferred.

```go
package main

import (
	"fmt"
	"github.com/najemi-software/tidal-dl/v2"
	"github.com/najemi-software/tidal-dl/v2/credentials"
)

func main() {
	refreshToken := "example"

	err := tidal_dl.RefreshTokenLogin(refreshToken)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(credentials.AccessToken)
}
```

### Login with an access token

If you do not have a refresh token or the user's attention, you can authenticate using an access token. Beware that when this expires, it will cause requests to the API to fail.

```go
package main

import (
	"fmt"
	"github.com/najemi-software/tidal-dl/v2"
	"github.com/najemi-software/tidal-dl/v2/credentials"
)

func main() {
	tokenType := "Bearer"
	accessToken := "example"

	err := tidal_dl.AccessTokenLogin(tokenType, accessToken)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println(credentials.AccessToken)
}
```
