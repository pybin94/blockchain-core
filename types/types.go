package types

type Wallet struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Balance    string `json:"balance"`
	Time       string `json:"time"`
}

type Block struct {
	Time         int64          `json:"time"`
	Hash         string         `json:"hash"`
	PrevHash     string         `json:"prevHash"`
	Nonce        int64          `json:"nonce"`
	Height       int64          `json:"height"`
	Transactions []*Transaction `json:"transaction"`
}
type Transaction struct {
	Block   int64  `json:"block"`
	Time    int64  `json:"time"`
	From    string `json:"from"`
	To      string `json:"To"`
	Amount  string `json:"amount"`
	Message string `json:"message"`
	Tx      string `json:"tx"`
}
