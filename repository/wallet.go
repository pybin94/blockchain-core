package repository

import (
	"block_chain/types"
	"context"
	"strconv"
	"time"

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
		"pubilcKey":  wallet.PublicKey,
		"time":       wallet.Time,
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
