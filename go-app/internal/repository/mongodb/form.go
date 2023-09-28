package db

import (
	"context"
	"log"
	"sober_driver/internal/domain"
	"sober_driver/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	collNameDriver = "form"
)

type FormRepo struct {
	db *mongo.Collection
}

func NewFormRepo(db *mongo.Database) *FormRepo {
	return &FormRepo{
		db: db.Collection(collNameDriver),
	}
}

func (fr *FormRepo) getFormsBase(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*domain.FormInput, error) {
	forms := []*domain.FormInput{}
	cursor, err := fr.db.Find(ctx, filter, opts...)
	if err != nil {
		return nil, utils.DropNoDocs(err)
	}

	if err = cursor.All(context.TODO(), &forms); err != nil {
		log.Fatal(err)
	}

	return forms, nil
}

func (fr *FormRepo) InsertData(ctx context.Context, data *domain.FormInput) error {
	_, err := fr.db.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}

func (fr *FormRepo) GetFormsById(ctx context.Context, userID string) ([]*domain.FormInput, error) {
	return fr.getFormsBase(ctx, bson.M{"user_id": userID})
}
