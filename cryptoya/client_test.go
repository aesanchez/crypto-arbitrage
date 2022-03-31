package cryptoya

import (
	"fmt"
	"log"
	"testing"
)

func TestGetCoinRates(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	r, err := c.GetCoinRates("usdt", "ars")
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range r {
		fmt.Println(k, v)
	}

	fmt.Println("--GetTopSellers--")
	r.GetTopSellers(5)
	fmt.Println("--GetTopBuyers--")
	r.GetTopBuyers(5)
}

func TestGetCoinRateFromExchange(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	rr, err := c.GetCoinRateFromExchange("lemoncash", "usdt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rr)
}
