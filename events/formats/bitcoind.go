package formats

type BitcoindEvent struct {
	BlockHash string
	TxID      string
	Raw       bool
	Decorate  bool
}
