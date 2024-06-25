package service

import (
	"block_chain/types"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func (s *Service) CreateBlock(txs []*types.Transactions, prevHash []byte, height int64) *types.Block {
	var pHash []byte

	if latesBlock, err := s.repository.GetLatestBlock(); err != nil {
		if err == mongo.ErrNoDocuments {
			s.log.Info("Genesis Block Will Be Created")
			newBlock := createBlickInner(txs, pHash, height)

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

func createBlickInner(txs []*types.Transactions, prevHash []byte, height int64) *types.Block {
	return &types.Block{
		Time:         time.Now().Unix(),
		Hash:         []byte{},
		Transactions: txs,
		PrevHash:     prevHash,
		Nonce:        0,
		Height:       height,
	}
}
