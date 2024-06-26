package service

import (
	"block_chain/types"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

type PowWork struct {
	Block      *types.Block `json:"block"`
	Target     *big.Int     `json:"target"`
	Difficulty int64        `json:"difficulty"`
}

func (s *Service) NewPow(b *types.Block) *PowWork {
	t := new(big.Int).SetInt64(1)
	t.Lsh(t, uint(256-s.difficulty))
	return &PowWork{Block: b, Target: t, Difficulty: s.difficulty}
}

func (p *PowWork) RunMining() (int64, []byte) {
	var iHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		// fmt.Println("nonce: ", nonce)
		d := p.makeHash(nonce)
		hash = sha256.Sum256(d)

		fmt.Printf("\r%x", hash)

		iHash.SetBytes(hash[:])

		if iHash.Cmp(p.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	return int64(nonce), hash[:]
}

func (p *PowWork) makeHash(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			p.Block.PrevHash,
			HashTransactions(p.Block),
			intToHex(p.Difficulty),
			intToHex(int64(nonce)),
		},
		[]byte{},
	)
}

func intToHex(number int64) []byte {
	// 300 -> 0x012C
	// binary.BigEndian -> 01
	b := new(bytes.Buffer)

	if err := binary.Write(b, binary.BigEndian, number); err != nil {
		panic(err)
	} else {
		return b.Bytes()
	}
}
