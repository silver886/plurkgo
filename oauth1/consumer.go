package oauth1

import (
	"net/url"

	"github.com/dghubble/oauth1"
)

// Consumer represents an Plurk consumer's key and secret and the callback URL.
type Consumer struct {
	config *oauth1.Config

	requestToken  string
	requestSecret string
}

// NewConsumer returns a new Consumer with the given consumer key and secret.
func NewConsumer(consumerKey, consumerSecret string) *Consumer {
	return &Consumer{
		config: &oauth1.Config{
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

func (c *Consumer) getAuthURL() (*url.URL, error) {
	if rt, rs, err := c.config.RequestToken(); err != nil {
		return nil, err
	} else if authURL, err := c.config.AuthorizationURL(rt); err != nil {
		return nil, err
	} else {
		c.requestToken = rt
		c.requestSecret = rs
		return authURL, nil
	}
}

func (c *Consumer) getToken(verifier string) (*Token, error) {
	if at, as, err := c.config.AccessToken(c.requestToken, c.requestSecret, verifier); err != nil {
		return nil, err
	} else {
		return NewToken(at, as), nil
	}
}
