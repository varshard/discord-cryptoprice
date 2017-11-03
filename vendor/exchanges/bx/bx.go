package bx

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// URL is bx.in.th api URL
const URL = "https://bx.in.th/api"

var (
	pairs   map[string]string
	timeout int32
	client  *http.Client
)

func init() {
	timeout, err := strconv.Atoi(os.Getenv("TIMEOUT"))
	if err != nil {
		timeout = 30
	}

	client = &http.Client{Timeout: time.Duration(timeout) * time.Second}
	pairs = map[string]string{
		"thb/btc": "1",
		"btc/ltc": "2",
		"btc/nmc": "3",
		"btc/xpm": "7",
		"btc/eth": "20",
		"thb/eth": "21",
		"thb/omg": "26",
		"thb/das": "22",
		"thb/evx": "28",
	}
}

// GetLastestPrice get average price of recent trades
func GetLastestPrice(pair string) float64 {
	pairID, exist := pairs[pair]
	if exist {
		url := URL + "/trade/?pairing=" + pairID
		res, reqErr := client.Get(url)
		if reqErr != nil {
			fmt.Println("Can't get price", reqErr)
		}
		defer res.Body.Close()

		var trades BXTrades
		readBodyErr := json.NewDecoder(res.Body).Decode(&trades)

		if readBodyErr != nil {
			panic(readBodyErr)
		}

		return trades.AVGTrade()
	}

	return -1
}

type BXTrades struct {
	Trades []BXTrade `json:"trades"`
}

func (trades BXTrades) AVGTrade() float64 {
	total := 0.0
	totalAmount := 0.0

	for _, trade := range trades.Trades {
		rate, err := strconv.ParseFloat(trade.Rate, 32)
		if err != nil {
			rate = 0
		}

		amount, err := strconv.ParseFloat(trade.Amount, 32)
		if err != nil {
			amount = 0
		}

		totalAmount += amount
		total += rate * amount
	}

	return total / totalAmount
}

type BXTrade struct {
	TradeID     string `json:"trade_id"`
	Rate        string `json:"rate"`
	Amount      string `json:"amount"`
	TradeDate   string `json:"trade_date"`
	OrderID     string `json:"order_id"`
	TradeType   string `json:"trade_type"`
	ReferenceID string `json:"reference_id"`
	Seconds     int    `json:"seconds"`
}
