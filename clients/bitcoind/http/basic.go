package http

import (
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/edobtc/cloudkit/config"
)

func NewBasicRpcClient() (*rpcclient.Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         config.Read().Node.Host,
		User:         config.Read().Node.RPCUser,
		Pass:         config.Read().Node.RPCPassword,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true,
	}
	return rpcclient.New(connCfg, nil)
}
