package bx

import (
	"encoding/json"
	"fmt"
	"math"
	"testing"
)

func TestGetLastestPrice(t *testing.T) {
	lastestPrice := GetLastestPrice("THB/ETH")

	if lastestPrice < 0 {
		t.Errorf("lastestPrice is: %E", lastestPrice)
	} else {
		fmt.Printf("lastestPrice %E", lastestPrice)
	}
}

func TestTradesUnmashal(t *testing.T) {
	tradesJSON := []byte(`{
    "trades": [
        {
            "trade_id": "1863532",
            "rate": "126099.00000000",
            "amount": "1.04804715",
            "trade_date": "2017-09-23 19:13:36",
            "order_id": "5617141",
            "trade_type": "buy",
            "reference_id": "0",
            "seconds": 197
        },
        {
            "trade_id": "1863545",
            "rate": "125100.00000000",
            "amount": "0.00342419",
            "trade_date": "2017-09-23 19:14:08",
            "order_id": "5617125",
            "trade_type": "sell",
            "reference_id": "0",
            "seconds": 165
        },
        {
            "trade_id": "1863550",
            "rate": "125999.00000000",
            "amount": "0.80000000",
            "trade_date": "2017-09-23 19:14:31",
            "order_id": "5617163",
            "trade_type": "buy",
            "reference_id": "0",
            "seconds": 142
        },
        {
            "trade_id": "1863551",
            "rate": "125999.00000000",
            "amount": "1.44290000",
            "trade_date": "2017-09-23 19:14:31",
            "order_id": "5617166",
            "trade_type": "buy",
            "reference_id": "0",
            "seconds": 142
        },
        {
            "trade_id": "1863552",
            "rate": "126000.00000000",
            "amount": "0.01278283",
            "trade_date": "2017-09-23 19:14:31",
            "order_id": "5617161",
            "trade_type": "buy",
            "reference_id": "0",
            "seconds": 142
        },
        {
            "trade_id": "1863556",
            "rate": "126000.00000000",
            "amount": "0.02657581",
            "trade_date": "2017-09-23 19:14:37",
            "order_id": "5617167",
            "trade_type": "sell",
            "reference_id": "0",
            "seconds": 136
        },
        {
            "trade_id": "1863567",
            "rate": "126000.00000000",
            "amount": "0.01050364",
            "trade_date": "2017-09-23 19:15:02",
            "order_id": "5617167",
            "trade_type": "sell",
            "reference_id": "0",
            "seconds": 111
        },
        {
            "trade_id": "1863568",
            "rate": "126000.00000000",
            "amount": "0.48289507",
            "trade_date": "2017-09-23 19:15:02",
            "order_id": "5617167",
            "trade_type": "sell",
            "reference_id": "0",
            "seconds": 111
        },
        {
            "trade_id": "1863585",
            "rate": "125005.00000000",
            "amount": "0.01000000",
            "trade_date": "2017-09-23 19:15:26",
            "order_id": "5617188",
            "trade_type": "sell",
            "reference_id": "0",
            "seconds": 87
        },
        {
            "trade_id": "1863598",
            "rate": "126000.00000000",
            "amount": "0.00021343",
            "trade_date": "2017-09-23 19:16:03",
            "order_id": "5617177",
            "trade_type": "buy",
            "reference_id": "0",
            "seconds": 50
        }
    ],
    "lowask": [
        {
            "order_id": "5617195",
            "rate": "125999.21000000",
            "amount": "0.03345971",
            "date_added": "2017-09-23 19:16:10",
            "order_type": "sell",
            "display_vol1": "4,215.90 THB",
            "display_vol2": "0.03345971 BTC"
        },
        {
            "order_id": "5617177",
            "rate": "126000.00000000",
            "amount": "0.96000493",
            "date_added": "2017-09-23 19:15:02",
            "order_type": "sell",
            "display_vol1": "120,960.62 THB",
            "display_vol2": "0.96000493 BTC"
        },
        {
            "order_id": "5617198",
            "rate": "126000.00000000",
            "amount": "0.01000000",
            "date_added": "2017-09-23 19:16:19",
            "order_type": "sell",
            "display_vol1": "1,260.00 THB",
            "display_vol2": "0.01000000 BTC"
        },
        {
            "order_id": "5617199",
            "rate": "126000.00000000",
            "amount": "0.00342419",
            "date_added": "2017-09-23 19:16:31",
            "order_type": "sell",
            "display_vol1": "431.45 THB",
            "display_vol2": "0.00342419 BTC"
        },
        {
            "order_id": "5617081",
            "rate": "126100.00000000",
            "amount": "0.02791948",
            "date_added": "2017-09-23 19:07:45",
            "order_type": "sell",
            "display_vol1": "3,520.65 THB",
            "display_vol2": "0.02791948 BTC"
        },
        {
            "order_id": "5617039",
            "rate": "126190.00000000",
            "amount": "0.03730168",
            "date_added": "2017-09-23 19:04:08",
            "order_type": "sell",
            "display_vol1": "4,707.10 THB",
            "display_vol2": "0.03730168 BTC"
        },
        {
            "order_id": "5617033",
            "rate": "126198.00000000",
            "amount": "0.03741703",
            "date_added": "2017-09-23 19:04:00",
            "order_type": "sell",
            "display_vol1": "4,721.95 THB",
            "display_vol2": "0.03741703 BTC"
        },
        {
            "order_id": "5616837",
            "rate": "126199.00000000",
            "amount": "2.55452329",
            "date_added": "2017-09-23 18:43:19",
            "order_type": "sell",
            "display_vol1": "322,378.28 THB",
            "display_vol2": "2.55452329 BTC"
        },
        {
            "order_id": "5605152",
            "rate": "126298.00000000",
            "amount": "0.08044160",
            "date_added": "2017-09-22 13:01:20",
            "order_type": "sell",
            "display_vol1": "10,159.61 THB",
            "display_vol2": "0.08044160 BTC"
        },
        {
            "order_id": "5617062",
            "rate": "126298.00000000",
            "amount": "0.03524288",
            "date_added": "2017-09-23 19:06:08",
            "order_type": "sell",
            "display_vol1": "4,451.11 THB",
            "display_vol2": "0.03524288 BTC"
        }
    ],
    "highbid": [
        {
            "order_id": "5617193",
            "rate": "125100.00",
            "amount": "0.00798957",
            "date_added": "2017-09-23 19:16:05",
            "order_type": "buy",
            "display_vol1": "999.50 THB",
            "display_vol2": "0.00798957 BTC"
        },
        {
            "order_id": "5617188",
            "rate": "125005.00",
            "amount": "0.26019018",
            "date_added": "2017-09-23 19:15:22",
            "order_type": "buy",
            "display_vol1": "32,525.07 THB",
            "display_vol2": "0.26019018 BTC"
        },
        {
            "order_id": "5617135",
            "rate": "125000.00",
            "amount": "0.37540649",
            "date_added": "2017-09-23 19:12:26",
            "order_type": "buy",
            "display_vol1": "46,925.81 THB",
            "display_vol2": "0.37540649 BTC"
        },
        {
            "order_id": "5617165",
            "rate": "125000.00",
            "amount": "0.28852767",
            "date_added": "2017-09-23 19:14:26",
            "order_type": "buy",
            "display_vol1": "36,065.96 THB",
            "display_vol2": "0.28852767 BTC"
        },
        {
            "order_id": "5617104",
            "rate": "124994.00",
            "amount": "1.99763169",
            "date_added": "2017-09-23 19:09:02",
            "order_type": "buy",
            "display_vol1": "249,691.98 THB",
            "display_vol2": "1.99763169 BTC"
        },
        {
            "order_id": "5617059",
            "rate": "124992.00",
            "amount": "1.28576371",
            "date_added": "2017-09-23 19:05:48",
            "order_type": "buy",
            "display_vol1": "160,710.18 THB",
            "display_vol2": "1.28576371 BTC"
        },
        {
            "order_id": "5616428",
            "rate": "124501.00",
            "amount": "0.08011984",
            "date_added": "2017-09-23 17:49:28",
            "order_type": "buy",
            "display_vol1": "9,975.00 THB",
            "display_vol2": "0.08011984 BTC"
        },
        {
            "order_id": "5616664",
            "rate": "124300.00",
            "amount": "0.40124698",
            "date_added": "2017-09-23 18:20:52",
            "order_type": "buy",
            "display_vol1": "49,875.00 THB",
            "display_vol2": "0.40124698 BTC"
        },
        {
            "order_id": "5617035",
            "rate": "124200.00",
            "amount": "0.04015700",
            "date_added": "2017-09-23 19:04:04",
            "order_type": "buy",
            "display_vol1": "4,987.50 THB",
            "display_vol2": "0.04015700 BTC"
        },
        {
            "order_id": "5617043",
            "rate": "124101.00",
            "amount": "0.04018904",
            "date_added": "2017-09-23 19:04:24",
            "order_type": "buy",
            "display_vol1": "4,987.50 THB",
            "display_vol2": "0.04018904 BTC"
        }
    ],
    "user_orders": []
}`)
	var trades BXTrades
	err := json.Unmarshal(tradesJSON, &trades)
	if err != nil {
		fmt.Printf(err.Error())
	}
	avg := trades.AVGTrade()
	if math.IsNaN(avg) {
		t.Error("avg is NaN")
	} else {
		fmt.Printf("avg is %E", avg)
	}
}
