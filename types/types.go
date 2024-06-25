package types

type Wallet struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
	Time       string `json:"time"`
}

type Block struct {
	Time         int64           `json:"time"`
	Hash         []byte          `json:"hash"`
	PrevHash     []byte          `json:"prevHash"`
	Nonce        int64           `json:"nonce"`
	Height       int64           `json:"height"`
	Transactions []*Transactions `json:"transactions"`
}
type Transactions struct {
	Block   int64  `json:"block"`
	Time    int64  `json:"time"`
	From    string `json:"from"`
	To      string `json:"To"`
	Amount  string `json:"amount"`
	Message string `json:"message"`
	Tx      string `json:"tx"`
}
