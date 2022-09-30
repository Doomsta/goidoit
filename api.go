package goidoit

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Client struct {
	url      string
	apiKey   string
	username string
	password string

	client *http.Client

	CMDB    *CmdbService
	Idoit   *IdoitService
	Console *ConsoleService
}

type SearchResult struct {
	DocumentId string `json:"documentId"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	Type       string `json:"type"`
	Link       string `json:"link"`
	Score      string `json:"score"`
	Status     string `json:"status"`
}

type GenericResponse[T any] struct {
	ID      int    `json:"id"`
	JsonRpc string `json:"jsonrpc"`
	Result  T      `json:"result"`
}

func NewClient(url, apiKey string, options ...options) *Client {
	c := &Client{
		url:    url,
		apiKey: apiKey,
		client: http.DefaultClient,
	}
	common := &service{client: c}

	c.CMDB = (*CmdbService)(common)
	c.Idoit = (*IdoitService)(common)
	c.Console = (*ConsoleService)(common)

	for _, opt := range options {
		opt(c)
	}

	return c
}

type service struct {
	client *Client
}

type Request struct {
	Version string      `json:"version"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	ID      int         `json:"id"`
}

func (c *Client) Request(ctx context.Context, method string, parameters any) ([]byte, *http.Response, error) {
	params, err := buildParams(c, parameters)
	if err != nil {
		return nil, nil, err
	}

	data := Request{
		Version: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return nil, nil, err
	}

	req, err := http.NewRequest("POST", c.url, bytes.NewBuffer(dataJSON))
	if err != nil {
		return nil, nil, err
	}
	req.Header.Add("content-type", "application/json")
	req.Header.Add("accept", "application/json")

	if c.username != "" && c.password != "" {
		req.Header["X-RPC-Auth-Username"] = []string{"admin"}
		req.Header["X-RPC-Auth-Password"] = []string{"admin"}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	b, err := io.ReadAll(resp.Body)
	return b, resp, err
}

func parse[T any](data []byte, _ *http.Response, _ error) (T, error) {
	d := GenericResponse[T]{}
	log.Println(string(data))
	if err := json.Unmarshal(data, &d); err != nil {
		return d.Result, err
	}
	return d.Result, nil
}

func buildParams(c *Client, parameters any) (any, error) {
	jsonParameters, err := json.Marshal(parameters)
	if err != nil {
		return nil, err
	}

	params := map[string]interface{}{}
	err = json.Unmarshal(jsonParameters, &params)
	if err != nil {
		return nil, err
	}

	jsonAPIkey, err := json.Marshal(struct {
		APIkey string `json:"apikey"`
	}{c.apiKey})
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonAPIkey, &params)
	return params, err
}
