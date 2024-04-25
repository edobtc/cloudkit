package rpc

import (
	"context"
	"fmt"

	formats "github.com/edobtc/cloudkit/events/formats/bitcoin"
)

func GetBlock(ctx context.Context, event formats.BitcoindEvent) (string, error) {
	fmt.Println(event)

	return fmt.Sprintf("get block count for %s!", event.BlockHash), nil
}
