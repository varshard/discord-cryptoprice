package bitfinex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// URL is bitfinex.com api URL
const URL = "https://api.bitfinex.com/v1"

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
		"btc/usd": "btcusd",
		"ltc/usd": "ltcusd",
		"ltc/btc": "ltcbtc",
		"eth/usd": "ethusd",
		"eth/btc": "ethbtc",
		"etc/btc": "etcbtc",
		"etc/usd": "etcusd",
		"rrt/usd": "rrtusd",
		"rrt/btc": "rrtbtc",
		"zec/usd": "zecusd",
		"zec/btc": "zecbtc",
		"xmr/usd": "xmrusd",
		"xmr/btc": "xmrbtc",
		"dsh/usd": "dshusd",
		"dsh/btc": "dshbtc",
		"bcc/btc": "bccbtc",
		"bcu/btc": "bcubtc",
		"bcc/usd": "bccusd",
		"bcu/usd": "bcuusd",
		"xrp/usd": "xrpusd",
		"xrp/btc": "xrpbtc",
		"iot/usd": "iotusd",
		"iot/btc": "iotbtc",
		"iot/eth": "ioteth",
		"eos/usd": "eosusd",
		"eos/btc": "eosbtc",
		"eos/eth": "eoseth",
		"san/usd": "sanusd",
		"san/btc": "sanbtc",
		"san/eth": "saneth",
		"omg/usd": "omgusd",
		"omg/btc": "omgbtc",
		"omg/eth": "omgeth",
		"bch/usd": "bchusd",
		"bch/btc": "bchbtc",
		"bch/eth": "bcheth",
		"neo/usd": "neousd",
		"neo/btc": "neobtc",
		"neo/eth": "neoeth",
		"etp/usd": "etpusd",
		"etp/btc": "etpbtc",
		"etp/eth": "etpeth",
	}
}

// GetLastestPrice get average price of recent trades
func GetLastestPrice(pair string) float64 {
	pairID, exist := pairs[pair]
	if exist {
		url := URL + "/pubticker/" + pairID
		res, reqErr := client.Get(url)
		if reqErr != nil {
			fmt.Println("Can't get price", reqErr)
		}
		defer res.Body.Close()

		var ticker Ticker
		readBodyErr := json.NewDecoder(res.Body).Decode(&ticker)

		if readBodyErr != nil {
			panic(readBodyErr)
		}

		price, convError := strconv.ParseFloat(ticker.LastPrice, 64)
		if convError != nil {
			return 0
		}

		return price
	}

	return -1
}

type Ticker struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}
