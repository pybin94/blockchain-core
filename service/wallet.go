package service

import (
	"block_chain/types"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func (s *Service) newWallet() (string, string, error) {
	p256 := elliptic.P256()
	if private, err := ecdsa.GenerateKey(p256, rand.Reader); err != nil {
		return "", "", err
	} else if private == nil {
		return "", "", errors.New("Pk is Nil")
	} else {
		privateKeyBytes := crypto.FromECDSA(private)
		privateKey := hexutil.Encode(privateKeyBytes)
		againPrivateKey, err := crypto.HexToECDSA(privateKey[2:])
		if err != nil {
			return "", "", err
		}
		cPublicKey := againPrivateKey.Public()
		publicKeyECDSA, ok := cPublicKey.(*ecdsa.PublicKey)

		if !ok {
			msg := fmt.Sprintf("error casting public key to ECDSA")
			panic(msg)
		}

		publicKey := crypto.PubkeyToAddress(*publicKeyECDSA)

		fmt.Println()
		fmt.Println("publicKey:", hexutil.Encode(publicKey[:]))

		return privateKey, hexutil.Encode(publicKey[:]), nil
	}
}

func (s *Service) MakeWallet() *types.Wallet {
	var wallet types.Wallet
	var err error

	if wallet.PrivateKey, wallet.PublicKey, err = s.newWallet(); err != nil {
		panic(err)
	}

	if err = s.repository.CreateNewWallet(&wallet); err != nil {
		return nil
	} else {
		return &wallet
	}
}

func (s *Service) GetWallet(pk string) (*types.Wallet, error) {
	if wallet, err := s.repository.GetWallet(pk); err != nil {
		return nil, err
	} else {
		fmt.Println(wallet)
		return wallet, nil
	}
}
