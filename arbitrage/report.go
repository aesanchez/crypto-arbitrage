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

const reportTemplate = `
Report: %v
Lemon:
	Starting ARS: %f
	Change fee: %f
	Transaction fee: %f

	USDT rate: %f
	Real USDT rate: %f
	UDSTs: %f
	USDTs after transaction: %f
Buenbit:
	Received USDTs: %f
	USDT rate: %f
	Final ARS: %f

	Profit: %f 
`

func GenerateReport(startingARS float64) (string, float64, error) {
	c, err := cryptoya.NewClient()
	if err != nil {
		return "", 0, err
	}

	// lemon side
	lemonRate, err := c.GetCoinRateFromExchange("lemoncash", "usdt")
	if err != nil {
		return "", 0, err
	}

	lemonBuyUSDTRate := lemonRate.Ask
	// we need to compute the real value given it's fee
	reaLemonBuyUSDTRate := lemonBuyUSDTRate / (1 - lemonBuyUSDTFee)
	lemonUSDTs := startingARS / reaLemonBuyUSDTRate

	// buenbit side
	buenbitRate, err := c.GetCoinRateFromExchange("buenbit", "usdt")
	if err != nil {
		return "", 0, err
	}

	receivedUSDTs := lemonUSDTs - lemonUSDTTransactionFee
	buenbitSellUSDTRate := buenbitRate.Bid
	finalARS := receivedUSDTs * buenbitSellUSDTRate

	profit := finalARS - startingARS

	report := fmt.Sprintf(reportTemplate,
		time.Now(), startingARS, lemonBuyUSDTFee, lemonUSDTTransactionFee, lemonBuyUSDTRate, reaLemonBuyUSDTRate, lemonUSDTs, receivedUSDTs,
		receivedUSDTs, buenbitSellUSDTRate, finalARS, profit)

	return report, profit, nil
}
