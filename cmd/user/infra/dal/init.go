package dal

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"time"
	"wargaming/common/config"
)

var (
	MongoDB *mongo.Database
)

func Init() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(config.Mongo.Url)
	clientOptions.SetAuth(options.Credential{
		Username: config.Mongo.UserName,
		Password: config.Mongo.Password,
	})
	clientOptions.SetMinPoolSize(uint64(config.Mongo.MinPoolSize))
	clientOptions.SetMaxPoolSize(uint64(config.Mongo.MaxPoolSize))
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		klog.Fatal(err)
	}
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		klog.Fatal(err)
	}

	MongoDB = client.Database(config.Mongo.DB)
}
