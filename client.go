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

// Client extands http.Client for Plurk API.
type Client struct {
	*http.Client
}

const (
	plurkBaseURL = "https://www.plurk.com"
)

// NewClient returns a new Client with the given OAuth config and token.
func NewClient(config *Config, token *Token) *Client {
	return &Client{
		Client: config.Config.Client(oauth1.NoContext, token.Token),
	}
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
	if resp, err := c.Get(requestURL); err != nil {
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
