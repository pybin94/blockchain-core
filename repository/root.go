package repository

import (
	"block_chain/config"
	"context"
	"time"

	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client

	wallet *mongo.Collection
	tx     *mongo.Collection

	log log15.Logger
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		log: log15.New("module", "repository"),
	}

	var err error
	ctx := context.Background()

	mongoConfig := config.Mongo

	if r.client, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoConfig.Uri)); err != nil {
		r.log.Error("failed to connect to mongo", "uri", mongoConfig.Uri)
		return nil, err
	} else if err := r.client.Ping(ctx, nil); err != nil {
		r.log.Error("failed to ping to mongo", "uri", mongoConfig.Uri)
		return nil, err
	} else {
		db := r.client.Database(config.Mongo.DB, nil)

		r.wallet = db.Collection("wallet")
		r.tx = db.Collection("wallet")

		r.log.Info("Success To Connect Repository", "info", time.Now().Unix(), "repository", config.Mongo.Uri)
		return r, nil
	}
}
