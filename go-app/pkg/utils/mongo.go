package utils

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo"
)

func DropNoDocs(err error) error {
	if err == mongo.ErrNoDocuments {
		return nil
	}
	return err
}

type DB struct {
	Client mongo.Client
	DB     *mongo.Database
}

type DBconfig struct {
	AppName string
	DBName  string
	URL     string
}

func Open(ctx context.Context, conf *DBconfig) (*DB, error) {
	opt := options.Client().SetAppName(conf.AppName).ApplyURI(conf.URL).SetMonitor(otelmongo.NewMonitor())
	cli, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Println("err conect db")
		return nil, err
	}
	log.Println("connect to db")

	err = cli.Ping(ctx, nil)
	if err != nil {
		log.Println("err ping")
		return nil, err
	}

	db := cli.Database(conf.DBName)

	response := &DB{
		Client: *cli,
		DB:     db,
	}

	return response, nil
}
