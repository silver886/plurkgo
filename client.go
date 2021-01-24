package plurkgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/dghubble/oauth1"
)

// Client extands http.Client for Plurk API 2.0.
type Client struct {
	client *http.Client
	oauth  *OAuth1
}

const (
	plurkBaseURL = "https://www.plurk.com"
)

// NewClient returns a new Client with the given OAuth config and token.
func NewClient(oauth *OAuth1) *Client {
	return &Client{
		client: oauth.consumer.config.Client(oauth1.NoContext, oauth.token.token),
		oauth:  oauth,
	}
}

// GetAuthURL for user to authorize the consumer.
func (c *Client) GetAuthURL() (*url.URL, error) {
	return c.oauth.GetAuthURL()
}

// GetToken with the verifier from user.
func (c *Client) GetToken(verifier string) error {
	return c.oauth.GetToken(verifier)
}

// Call API endpoint with given parameter and response object.
func (c *Client) Call(endpoint string, parameter map[string][]string, response interface{}) {
	c.call(endpoint, parameter, response)
}

func (c *Client) call(endpoint string, parameter map[string][]string, response interface{}) {
	requestURL := plurkBaseURL + endpoint
	if parameter != nil {
		requestURL += "?" + url.Values(parameter).Encode()
	}
	if resp, err := c.client.Get(requestURL); err != nil {
		log.Fatalln(err)
	} else if resp.StatusCode/100 != 2 {
		resp.Body.Close()
		log.Fatalln(resp.Status)
	} else if body, err := ioutil.ReadAll(resp.Body); err != nil {
		resp.Body.Close()
		log.Fatalln(err)
	} else {
		if b, ok := response.(bool); ok && b {
			var responseObject map[string]interface{}
			if err := json.Unmarshal(body, &responseObject); err != nil {
				resp.Body.Close()
				log.Fatalln(err)
			} else if responseJSON, err := json.MarshalIndent(responseObject, "", "  "); err != nil {
				resp.Body.Close()
				log.Fatalln(err)
			} else {
				fmt.Println(string(responseJSON))
			}
		} else if err := json.Unmarshal(body, response); err != nil {
			resp.Body.Close()
			log.Fatalln(err)
		}
		resp.Body.Close()
	}
}
