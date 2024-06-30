package repository

import (
	"block_chain/types"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) CreateNewWallet(wallet *types.Wallet) error {
	ctx := context.Background()
	wallet.Time = strconv.FormatInt(time.Now().Unix(), 10)

	opt := options.Update().SetUpsert(true)

	filter := bson.M{"privateKey": wallet.PrivateKey}
	update := bson.M{"$set": bson.M{
		"privateKey": wallet.PrivateKey,
		"publicKey":  wallet.PublicKey,
		"time":       wallet.Time,
		"balance":    wallet.Balance,
	}}

	if _, err := r.wallet.UpdateOne(ctx, filter, update, opt); err != nil {
		return err
	} else {
		return nil
	}
}

func (r *Repository) GetWallet(pk string) (*types.Wallet, error) {
	ctx := context.Background()

	filter := bson.M{"privateKey": pk}

	var wallet types.Wallet

	if err := r.wallet.FindOne(ctx, filter, options.FindOne()).Decode(&wallet); err != nil {
		return nil, err
	} else {
		return &wallet, nil
	}
}

func (r *Repository) GetWalletByPublicKey(publicKey string) (*types.Wallet, error) {
	ctx := context.Background()

	filter := bson.M{"publicKey": publicKey}
	fmt.Println(filter)
	var wallet types.Wallet

	if err := r.wallet.FindOne(ctx, filter, options.FindOne()).Decode(&wallet); err != nil {
		return nil, err
	} else {
		return &wallet, nil
	}
}

func (r *Repository) UpsertWalletsWhenTransfer(from, to, fromBalance, toBalance string) error {
	ctx := context.Background()
	opt := options.Update().SetUpsert(true)

	if from != (common.Address{}).String() {
		filter := bson.M{"publicKey": from}
		update := bson.M{"$set": bson.M{
			"balance": fromBalance,
		}}

		_, err := r.wallet.UpdateOne(ctx, filter, update, opt)
		if err != nil {
			return err
		}
	}

	filter := bson.M{"publicKey": to}
	update := bson.M{"$set": bson.M{
		"balance": toBalance,
	}}
	_, err := r.wallet.UpdateOne(ctx, filter, update, opt)
	if err != nil {
		return err
	}

	return nil
}
