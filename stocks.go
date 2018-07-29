package stocks

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Stock represents a stock or other security
type Stock struct {
	Code   string  `json:"symbol"`
	Date   string  `json:"date"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
	Close  float64 `json:"close"`
	Open   float64 `json:"open"`
	Change float64
}

// GetStock returns the most recent information about a given stock
func GetStock(code string) *Stock {
	res, err := http.Get("http://pseapi.com/api/stock/" + code)
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}
	var stocks []Stock
	json.Unmarshal(body, &stocks)
	stock := stocks[len(stocks)-1]
	stock.Change = (stock.Bid - stock.Close) / stock.Bid * 100
	return &stock
}
