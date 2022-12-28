package bitfinex

import (
	"context"
	"encoding/json"
	"time"

	redis "github.com/go-redis/redis/v8"
)

const (
	KeyStringBase = "prices/bitfinex/btcusd"

	DefaultCacheTimeout = time.Second * 3
)

func (b *BitfinexClient) AddCacheConnection(conn *redis.Client) {
	b.cacheConn = conn
}

func (b *BitfinexClient) FetchCached() (*TickerData, error) {
	ticker := TickerData{}
	ctx := context.Background()

	strData := b.cacheConn.Get(ctx, KeyStringBase)

	if strData.Err() == redis.Nil {
		return nil, nil
	}

	data, err := strData.Bytes()
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &ticker)

	// TODO turn into time.Time
	// if ticker.Timestamp.Add(AllowedInterval) {
	// return nil, ErrCacheEntryExpired
	// }

	return &ticker, nil
}

func (b *BitfinexClient) WriteCached(t *TickerData) error {
	d, err := json.Marshal(t)
	if err != nil {
		return err
	}

	ctx := context.Background()
	// ctx = context.WithTimeout(ctx, b.CacheTimeout)
	strCmd := b.cacheConn.Set(ctx, KeyStringBase, string(d), DefaultCacheTimeout)

	return strCmd.Err()
}
