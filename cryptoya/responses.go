package cryptoya

import (
	"fmt"
	"sort"
)

type GetAllCoinRatesResponse map[string]CoinResponse

type CoinResponse struct {
	// Precio de compra reportado por el exchange, sin sumar comisiones.
	Ask float64 `json:"ask"`
	// Precio de compra final incluyendo las comisiones de transferencia y trade.
	TotalAsk float64 `json:"totalAsk"`
	// Precio de venta reportado por el exchange, sin restar comisiones.
	Bid float64 `json:"bid"`
	// Precio de venta final incluyendo las comisiones de transferencia y trade.
	TotalBid float64 `json:"totalBid"`
	// Timestamp del momento en que fue actualizada esta cotización.
	Time int `json:"time"`
}

func (g GetAllCoinRatesResponse) GetTopSellers(limit int) {
	type wrapper struct {
		CoinResponse
		Exchange string
	}
	a := []wrapper{}
	for k, v := range g {
		a = append(a, wrapper{CoinResponse: v, Exchange: k})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].TotalAsk < a[j].TotalAsk
	})
	i := 0
	for _, c := range a {
		if i == limit {
			break
		}
		fmt.Println(c.Exchange, c.CoinResponse)
		i++
	}
}

func (g GetAllCoinRatesResponse) GetTopBuyers(limit int) {
	type wrapper struct {
		CoinResponse
		Exchange string
	}
	a := []wrapper{}
	for k, v := range g {
		a = append(a, wrapper{CoinResponse: v, Exchange: k})
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i].TotalBid > a[j].TotalBid
	})
	i := 0
	for _, c := range a {
		if i == limit {
			break
		}
		fmt.Println(c.Exchange, c.CoinResponse)
		i++
	}
}
