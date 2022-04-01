package main

import (
	"crypto-arbitrage/arbitrage"
	"crypto-arbitrage/mailer"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	minARS        float64
	stepARS       float64
	maxARS        float64
	minimumProfit float64

	interval time.Duration

	// sender configs
	email = os.Getenv("CRYPTO_EMAIL")
	pass  = os.Getenv("CRYPTO_APP_PASS")
	to    = os.Getenv("CRYPTO_TO")
)

func main() {
	flag.Float64Var(&minARS, "min-ars", 10_000, "")
	flag.Float64Var(&stepARS, "step-ars", 10_000, "")
	flag.Float64Var(&maxARS, "max-ars", 100_000, "")

	flag.Float64Var(&minimumProfit, "min-profit", 0, "")
	intervalDuration := flag.String("interval", "1m", "")
	flag.Parse()
	var err error
	interval, err = time.ParseDuration(*intervalDuration)
	if err != nil {
		panic(err)
	}

	t := time.NewTicker(interval)
	for {
		if err := iteration(); err != nil {
			log.Println(err)
		}
		<-t.C
	}
}

func iteration() error {
	lemonRate, buenbitRate, err := arbitrage.GetRates()
	if err != nil {
		return err
	}
	// start testing until the first one that generates profit
	var (
		startingARS = minARS
		report      string
		profit      float64
	)
	for startingARS <= maxARS {
		report, profit, err = arbitrage.GenerateReport(lemonRate, buenbitRate, startingARS)
		if err != nil {
			return err
		}

		if profit > startingARS*minimumProfit {
			fmt.Println("### JACKPOT!!! ###")
			fmt.Print(report)
			err := mailer.NewClient(email, pass).Send(to, report)
			if err != nil {
				return err
			}
			fmt.Println("Email sent")
			return nil
		}

		startingARS += stepARS
	}
	fmt.Println("### Found nothing, last report was: ###")
	fmt.Print(report)
	return nil
}
