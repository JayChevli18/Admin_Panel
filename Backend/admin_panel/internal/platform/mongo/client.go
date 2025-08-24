package mongo

import (
	context "context"
	time "time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Client struct {
	DB *mongo.Database
}

func NewMongoClient(mongoURI, mongoDB string, logger *zap.Logger) (*Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		logger.Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	if err := mongoClient.Ping(ctx, nil); err != nil {
		logger.Error("Failed to ping MongoDB", zap.Error(err))
		return nil, err
	}

	logger.Info("Connected to MongoDB", zap.String("db", mongoDB))

	db := mongoClient.Database(mongoDB)

	return &Client{DB: db}, nil
}
