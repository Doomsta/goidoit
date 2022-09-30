package goidoit

import (
	"crypto/tls"
	"net/http"
)

type options func(*Client)

func WithUserCredencials(username, password string) options {
	return func(c *Client) {
		c.username = username
		c.password = password
	}
}

func WithHTTPClient(httpC *http.Client) options {
	return func(c *Client) {
		c.client = httpC
	}
}

func WithInsecure() options {
	return func(c *Client) {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}
