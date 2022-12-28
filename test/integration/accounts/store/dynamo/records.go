package main

import (
	"github.com/edobtc/cloudkit/store/dynamodb"
	"github.com/sirupsen/logrus"
)

const (
	tablename = "edobtc_accounts"
)

type records map[string]interface{}

type Obj struct {
	PK            string `json:"PK"`
	SK            string `json:"SK"`
	Name          string `json:"name,omitempty"`
	Balance       int    `json:"balance,omitempty"`
	ExternalID    string `json:"externalID,omitempty"`
	TransactionID string `json:"transactionID,omitempty"`
}

func jsonTest() {
	w := dynamodb.NewDynamoWriter(tablename)

	ac := Obj{
		PK:         "rza|5555",
		SK:         "name|wutang",
		Name:       "wutang",
		ExternalID: "36",
	}

	err := w.Write(ac)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func runRecordWriteTest() {
	w := dynamodb.NewDynamoWriter(tablename)

	testRecord := records{
		"PK":         "cashapp|ee3345",
		"SK":         "name|edobtc",
		"name":       "edobtc",
		"balance":    35,
		"externalID": "user_45",
	}

	err := w.Write(testRecord)
	if err != nil {
		logrus.Error(err)
		return
	}
}

func main() {
	runRecordWriteTest()
	jsonTest()
}
