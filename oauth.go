package plurkgo

import "github.com/dghubble/oauth1"

// Config represents an Plurk consumer's key and secret and the callback URL.
type Config struct {
	*oauth1.Config
}

// Token is an access token which allows a consumer to access resources from Plurk.
type Token struct {
	*oauth1.Token
}

// NewConfig returns a new Config with the given consumer key and secret.
func NewConfig(consumerKey, consumerSecret string) *Config {
	return &Config{
		Config: &oauth1.Config{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			Endpoint: oauth1.Endpoint{
				RequestTokenURL: "https://www.plurk.com/OAuth/request_token",
				AuthorizeURL:    "https://www.plurk.com/OAuth/authorize",
				AccessTokenURL:  "https://www.plurk.com/OAuth/access_token",
			},
		},
	}
}

// NewToken returns a new Token with the given token and token secret.
func NewToken(token, tokenSecret string) *Token {
	return &Token{
		Token: &oauth1.Token{
			Token:       token,
			TokenSecret: tokenSecret,
		},
	}
}
