package main

import (
	"crypto-arbitrage/arbitrage"
	"crypto-arbitrage/mailer"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	startingARS   float64 = 25_000
	minimumProfit         = 0.001 // 0.1%

	interval = 5 * time.Second

	// sender configs
	email = os.Getenv("CRYPTO_EMAIL")
	pass  = os.Getenv("CRYPTO_APP_PASS")
	to    = os.Getenv("CRYPTO_TO")
)

func main() {
	t := time.NewTicker(interval)
	for {
		<-t.C
		report, profit, err := arbitrage.GenerateReport(startingARS)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Print(report)
		if profit > startingARS*minimumProfit {
			// jackpot!
			err := mailer.NewClient(email, pass).Send(to, report)
			if err != nil {
				log.Println(err)
				continue
			}
		}
	}
}
