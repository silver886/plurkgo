package plurkgo

import (
	"net/url"

	"github.com/dghubble/oauth1"
)

// Consumer represents an Plurk consumer's key and secret and the callback URL.
type Consumer struct {
	*oauth1.Config

	requestToken  string
	requestSecret string
}

// Token is an access token which allows a consumer to access resources from Plurk.
type Token struct {
	*oauth1.Token
}

// NewConsumer returns a new Consumer with the given consumer key and secret.
func NewConsumer(consumerKey, consumerSecret string) *Consumer {
	return &Consumer{
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

// GetAuthURL for user to authorize the consumer.
func (c *Consumer) GetAuthURL() (*url.URL, error) {
	if rt, rs, err := c.Config.RequestToken(); err != nil {
		return nil, err
	} else if authURL, err := c.Config.AuthorizationURL(rt); err != nil {
		return nil, err
	} else {
		c.requestToken = rt
		c.requestSecret = rs
		return authURL, nil
	}
}

// GetToken with the verifier from user.
func (c *Consumer) GetToken(verifier string) (*Token, error) {
	if at, as, err := c.Config.AccessToken(c.requestToken, c.requestSecret, verifier); err != nil {
		return nil, err
	} else {
		return NewToken(at, as), nil
	}
}
