package rpc

import (
	"context"
	"fmt"

	"github.com/edobtc/cloudkit/events/formats"
)

func GetTransaction(ctx context.Context, event formats.BitcoindEvent) (string, error) {
	// Global Count
	fmt.Println(event)

	return fmt.Sprintf("get transaction %s", event.TxID), nil
}
