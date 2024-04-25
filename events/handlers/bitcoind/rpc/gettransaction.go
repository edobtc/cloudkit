package rpc

import (
	"context"
	"fmt"

	formats "github.com/edobtc/cloudkit/events/formats/bitcoin"
)

func GetTransaction(ctx context.Context, event formats.BitcoindEvent) (string, error) {
	fmt.Println(event)

	return fmt.Sprintf("get transaction %s", event.TxID), nil
}
