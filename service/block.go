package service

import (
	"block_chain/types"
	"bytes"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateBlock(txs []*types.Transaction, prevHash []byte, height int64) *types.Block {
	var pHash []byte

	if latesBlock, err := s.repository.GetLatestBlock(); err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Info("Genesis Block Will Be Created")
			genesisMessage := "This Is First Genesis Block"
			tx := createTransaction(genesisMessage, "publicKey", "", "", 1)

			newBlock := createBlickInner([]*types.Transaction{tx}, pHash, height)

			pow := s.NewPow(newBlock)
			newBlock.Nonce, newBlock.Hash = pow.RunMining()

			return newBlock
		} else {
			s.log.Crit("Failed To Get Latest Block")
			panic(err)
		}
	} else {
		pHash = latesBlock.Hash

		newBlock := createBlickInner(txs, pHash, height)
		pow := s.NewPow(newBlock)

		newBlock.Nonce, newBlock.Hash = pow.RunMining()

		return newBlock
	}
}

func createBlickInner(txs []*types.Transaction, prevHash []byte, height int64) *types.Block {
	return &types.Block{
		Time:         time.Now().Unix(),
		Hash:         []byte{},
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Height:       height,
	}
}

func createTransaction(message, from, to, amount string, block int64) *types.Transaction {
	data := struct {
		Message string `json:"message"`
		From    string `json:"from"`
		To      string `json:"to"`
		Amount  string `json:"amount"`
	}{
		Message: message,
		From:    from,
		To:      to,
		Amount:  amount,
	}

	dataToSign := fmt.Sprintf("%x\n", data)

	pk := ""

	if ecdsaPrivateKey, err := crypto.HexToECDSA(pk); err != nil {
		panic(err)
	} else if r, s, err := ecdsa.Sign(rand.Reader, ecdsaPrivateKey, []byte(dataToSign)); err != nil {
		panic(err)
	} else {
		signature := append(r.Bytes(), s.Bytes()...)

		return &types.Transaction{
			Block:   block,
			Time:    time.Now().Unix(),
			From:    from,
			To:      to,
			Amount:  amount,
			Message: message,
			Tx:      hex.EncodeToString(signature),
		}
	}

}

func HashTransactions(b *types.Block) []byte {
	var txHashes [][]byte

	for _, tx := range b.Transactions {
		var encoded bytes.Buffer
		enc := gob.NewEncoder((&encoded))

		if err := enc.Encode(tx); err != nil {
			panic(err)
		} else {
			txHashes = append(txHashes, encoded.Bytes())
		}
	}

	tree := NewMerkleTree(txHashes)

	return tree.RootNode.Data
}
