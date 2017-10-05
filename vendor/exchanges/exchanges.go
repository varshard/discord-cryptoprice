package exchanges

import (
	"exchanges/bitfinex"
	"exchanges/bx"
	"fmt"
)

// GetLastestPrice from an exchange's API
func GetLastestPrice(exchange string, pair string) string {
	var price float64
	switch exchange {
	case "bx":
		price = bx.GetLastestPrice(pair)
	case "bitfinex":
		price = bitfinex.GetLastestPrice(pair)
	}

	if price > 0 {
		return fmt.Sprintf("%s %s %f", exchange, pair, price)
	}

	return fmt.Sprintf("Exchange %s isn't supported", exchange)
}
