package plurkgo

import "net/url"

// OAuth1 is the authorization method used by Plurk API 2.0.
type OAuth1 struct {
	consumer *Consumer
	token    *Token
}

// NewOAuth1 returns a new Client with the given OAuth config and token.
func NewOAuth1(consumer *Consumer, token *Token) *OAuth1 {
	return &OAuth1{
		consumer: consumer,
		token:    token,
	}
}

// GetAuthURL for user to authorize the consumer.
func (o *OAuth1) GetAuthURL() (*url.URL, error) {
	return o.consumer.getAuthURL()
}

// GetToken with the verifier from user.
func (o *OAuth1) GetToken(verifier string) error {
	if token, err := o.consumer.getToken(verifier); err == nil {
		o.token = token
	} else {
		return err
	}
	return nil
}
