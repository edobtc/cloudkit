package rpc

import (
	"context"
	"fmt"

	"github.com/edobtc/cloudkit/events/formats"
)

func GetBlock(ctx context.Context, event formats.BitcoindEvent) (string, error) {
	fmt.Println(event)

	return fmt.Sprintf("get block count for %s!", event.BlockHash), nil
}
