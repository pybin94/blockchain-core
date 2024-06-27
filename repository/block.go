package repository

import (
	"block_chain/types"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) GetLatestBlock() (*types.Block, error) {
	ctx := context.Background()

	var block types.Block
	opt := options.FindOne().SetSort(bson.M{"time": -1})

	if err := r.block.FindOne(ctx, bson.M{}, opt).Decode(&block); err != nil {
		return nil, err
	} else {
		return &block, nil
	}
}

func (r *Repository) SaveBlock(newBlock *types.Block) error {
	ctx := context.Background()

	filter := bson.M{"hash": newBlock.Hash}
	update := bson.M{"$set": bson.M{
		"time":         newBlock.Time,
		"hash":         newBlock.Hash,
		"prevHash":     newBlock.PrevHash,
		"nonce":        newBlock.Nonce,
		"height":       newBlock.Height,
		"Transactions": newBlock.Transactions,
	}}

	if _, err := r.block.UpdateOne(ctx, filter, update, options.Update().SetUpsert(true)); err != nil {
		return err
	} else {
		return nil
	}
}
