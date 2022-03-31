package cryptoya

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	URL = "https://criptoya.com/api"
)

type Client struct {
	c *http.Client
}

func NewClient() (*Client, error) {
	return &Client{c: http.DefaultClient}, nil
}

func (c *Client) GetCoinRates(coin string, fiat string) (GetAllCoinRatesResponse, error) {
	resp, err := c.c.Get(fmt.Sprintf("%s/%s/%s/1", URL, coin, fiat))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var o GetAllCoinRatesResponse
	err = json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return o, nil
}

func (c *Client) GetCoinRateFromExchange(exchange string, coin string) (*CoinResponse, error) {
	resp, err := c.c.Get(fmt.Sprintf("%s/%s/%s", URL, exchange, coin))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var o CoinResponse
	err = json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}

	return &o, nil
}
