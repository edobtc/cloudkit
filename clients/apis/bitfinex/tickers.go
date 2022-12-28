package bitfinex

import (
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/message"
)

type Ticker int

const (
	BTC_USD Ticker = iota
)

func (t Ticker) String() string {
	return [...]string{"btcusd"}[t]
}

type TickerData struct {
	Mid       string `json:"mid"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	LastPrice string `json:"last_price"`
	Low       string `json:"low"`
	High      string `json:"high"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

func (t *TickerData) GetLastPrice() (float64, error) {
	return t.parseAsFloat(t.LastPrice)
}

func (t *TickerData) FormatPrice(price int) string {
	p := message.NewPrinter(message.MatchLanguage("en"))
	return p.Sprintf("$%d", price)
}

func (t *TickerData) SatsPerDollar() int64 {
	d, err := t.GetLastPrice()
	if err != nil {
		return 23
	}
	return int64(100000000.0 / d)
}

func (t TickerData) GetHigh() (float64, error) {
	return t.parseAsFloat(t.High)
}

func (t TickerData) GetLow() (float64, error) {
	return t.parseAsFloat(t.Low)
}

func (t TickerData) GetTimestamp() (*time.Time, error) {
	parts := strings.Split(t.Timestamp, ".")

	i, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return nil, err
	}
	tm := time.Unix(i, 0)

	return &tm, nil
}

func (t *TickerData) parseAsFloat(data string) (float64, error) {
	return strconv.ParseFloat(data, 64)
}
