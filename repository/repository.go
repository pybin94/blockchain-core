package repository

import (
	"block_chain/config"
	"context"

	"github.com/inconshreveable/log15"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
	config *config.Config
	db     *mongo.Database
	log    log15.Logger
}

func NewRepository(config *config.Config) (*Repository, error) {
	r := &Repository{
		config: config,
		log:    log15.New("module", "repository"),
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
		r.db = r.client.Database(mongoConfig.DB)

		// TODO => 컬랙션 연결

		r.log.Info("success to connect Repository", "uri", mongoConfig.Uri, "db", mongoConfig.DB)
		return r, nil
	}
}
