package plurkgo

import (
	"github.com/dghubble/oauth1"
)

// Token is an access token which allows a consumer to access resources from Plurk.
type Token struct {
	token *oauth1.Token
}

// NewToken returns a new Token with the given token and token secret.
func NewToken(token, tokenSecret string) *Token {
	return &Token{
		token: &oauth1.Token{
			Token:       token,
			TokenSecret: tokenSecret,
		},
	}
}
