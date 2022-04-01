package arbitrage

import (
	"crypto-arbitrage/cryptoya"
	"fmt"
	"time"
)

var (
	lemonBuyUSDTFee         = 0.01
	lemonUSDTTransactionFee = 0.5
)

const reportTemplate = `Report: %v
Lemon:
	Starting ARS: %f
	Change fee: %f %% 
	Transaction fee: %f USDT

	USDT rate: %f
	Real USDT rate: %f
	UDSTs: %f
	USDTs after transaction: %f
Buenbit:
	Received USDTs: %f
	USDT rate: %f
	Final ARS: %f

	Profit: %f
	Profit (%%): %f

`

func GetRates() (*cryptoya.CoinResponse, *cryptoya.CoinResponse, error) {
	c, err := cryptoya.NewClient()
	if err != nil {
		return nil, nil, err
	}
	// lemon side
	lemonRate, err := c.GetCoinRateFromExchange("lemoncash", "usdt")
	if err != nil {
		return nil, nil, err
	}

	// buenbit side
	buenbitRate, err := c.GetCoinRateFromExchange("buenbit", "usdt")
	if err != nil {
		return nil, nil, err
	}

	return lemonRate, buenbitRate, nil
}

func GenerateReport(lemonRate, buenbitRate *cryptoya.CoinResponse, startingARS float64) (string, float64, error) {
	lemonBuyUSDTRate := lemonRate.Ask
	// we need to compute the real value given it's fee
	reaLemonBuyUSDTRate := lemonBuyUSDTRate / (1 - lemonBuyUSDTFee)
	lemonUSDTs := startingARS / reaLemonBuyUSDTRate

	receivedUSDTs := lemonUSDTs - lemonUSDTTransactionFee
	buenbitSellUSDTRate := buenbitRate.Bid
	finalARS := receivedUSDTs * buenbitSellUSDTRate

	profit := finalARS - startingARS

	report := fmt.Sprintf(reportTemplate,
		time.Now(), startingARS, lemonBuyUSDTFee*100, lemonUSDTTransactionFee, lemonBuyUSDTRate, reaLemonBuyUSDTRate, lemonUSDTs, receivedUSDTs,
		receivedUSDTs, buenbitSellUSDTRate, finalARS, profit, 100*profit/startingARS)

	return report, profit, nil
}
