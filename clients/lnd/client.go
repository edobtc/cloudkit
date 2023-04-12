package lnd

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/lightningnetwork/lnd/lnrpc"
	"github.com/lightningnetwork/lnd/lnrpc/routerrpc"
	"github.com/lightningnetwork/lnd/macaroons"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/macaroon.v2"
)

var (
	ErrNoMacaroonProvided = errors.New("no macaroon data provided, provide either a hex or filepath")
)

type ClientOptions struct {
	TLSCertPath  string
	MacaroonPath string
	MacaroonHex  string
	Host         string
}

func NewClient(opt ClientOptions) (lnrpc.LightningClient, error) {
	conn, err := GRPCTransport(opt)
	if err != nil {
		return nil, err
	}

	client := lnrpc.NewLightningClient(conn)

	return client, nil
}

func NewRouterClient(opt ClientOptions) (routerrpc.RouterClient, error) {
	conn, err := GRPCTransport(opt)
	if err != nil {
		return nil, err
	}

	routerClient := routerrpc.NewRouterClient(conn)
	return routerClient, nil
}

func decodeMacaroonHex(mHex string) ([]byte, error) {
	macaroonHex := []byte(mHex)
	macFromHex := make([]byte, hex.DecodedLen(len(macaroonHex)))

	_, err := hex.Decode(macFromHex, macaroonHex)
	if err != nil {
		return macaroonHex, err
	}

	return macaroonHex, nil
}

func loadMacaroon(hexString string, filepath string) (*macaroon.Macaroon, error) {
	var data []byte
	var err error

	if hexString == "" && filepath == "" {
		return nil, ErrNoMacaroonProvided
	}

	if hexString != "" {
		data, err = decodeMacaroonHex(hexString)
		if err != nil {
			return nil, err
		}
	}

	if filepath != "" {
		data, err = ioutil.ReadFile(filepath)
		if err != nil {
			return nil, err
		}
	}

	mac := &macaroon.Macaroon{}
	if err = mac.UnmarshalBinary(data); err != nil {
		return nil, err
	}

	return mac, nil
}

func GRPCTransport(opt ClientOptions) (*grpc.ClientConn, error) {
	tlsCreds, err := credentials.NewClientTLSFromFile(opt.TLSCertPath, "")
	if err != nil {
		fmt.Println("Cannot get node tls credentials", err)
		return nil, err
	}

	mac, err := loadMacaroon(opt.MacaroonHex, opt.MacaroonPath)
	if err != nil {
		return nil, err
	}

	mcCred, err := macaroons.NewMacaroonCredential(mac)
	if err != nil {
		return nil, err
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(tlsCreds),
		grpc.WithBlock(),
		grpc.WithPerRPCCredentials(mcCred),
	}

	conn, err := grpc.Dial(opt.Host, opts...)
	if err != nil {
		return nil, err
	}

	return conn, err
}
