package user

import (
	"errors"
	"strings"

	"github.com/10gen/stitch-cli/auth"
)

// Errors related to user configuration.
var (
	ErrNotLoggedIn = errors.New("you are not logged in")
)

// User stores the user's login credentials and some metadata.
type User struct {
	APIKey       string `yaml:"api_key"`
	Username     string `yaml:"username"`
	RefreshToken string `yaml:"refresh_token"`
	AccessToken  string `yaml:"access_token"`
}

// LoggedIn returns a boolean representing whether the user is logged in or not
func (u *User) LoggedIn() bool {
	return auth.ValidAccessToken(u.AccessToken)
}

// TokenIsExpired returns a boolean representing whether or not the token is expired
// or an error if the token is invalid
func (u *User) TokenIsExpired() (bool, error) {
	token, err := auth.NewJWT(u.AccessToken)
	if err != nil {
		return false, err
	}

	return token.Expired(), nil
}

// RedactedAPIKey returns a string representing the user's API key
// with everything but the last portion of the key displayed as "*"
func (u *User) RedactedAPIKey() string {
	apiKeyParts := strings.Split(u.APIKey, "-")
	redactedParts := make([]string, len(apiKeyParts))

	lastIndex := len(apiKeyParts) - 1
	for i := 0; i < lastIndex; i++ {
		redactedParts[i] = strings.Repeat("*", len(apiKeyParts[i]))
	}
	redactedParts[lastIndex] = apiKeyParts[lastIndex]

	return strings.Join(redactedParts, "-")
}
