package lnd

import (
	"bytes"
	"html/template"
)

type ConfigValues struct {
	Bitcoind Bitcoind
}

type Bitcoind struct {
	Host    string
	RPCUser string
	RPCPass string
}

func RenderTemplate() ([]byte, error) {
	var data bytes.Buffer
	tmpl, _ := template.ParseFiles("lnd/config.tmpl")
	err := tmpl.Execute(&data, ConfigValues{
		Bitcoind: Bitcoind{
			Host:    "test",
			RPCUser: "test",
			RPCPass: "test",
		},
	})

	return data.Bytes(), err
}
