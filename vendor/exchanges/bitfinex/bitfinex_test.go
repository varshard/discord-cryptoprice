package bitfinex

import (
	"fmt"
	"testing"
)

func TestGetLastestPrice(t *testing.T) {
	lastestPrice := GetLastestPrice("eth/usd")

	if lastestPrice < 0 {
		t.Errorf("lastestPrice is: %f", lastestPrice)
	} else {
		fmt.Printf("lastestPrice %f", lastestPrice)
	}
}
