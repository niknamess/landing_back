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
	collNameUser = "user"
)

type UserRepo struct {
	db *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *UserRepo {
	return &UserRepo{
		db: db.Collection(collNameUser),
	}
}

func (ur *UserRepo) getUserBase(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) (*domain.User, error) {
	user := new(domain.User)
	err := ur.db.FindOne(ctx, filter, opts...).Decode(user)
	if err != nil {
		return nil, utils.DropNoDocs(err)
	}

	return user, nil
}
func (ur *UserRepo) getUsersBase(ctx context.Context, filter interface{}, opts ...*options.FindOptions) ([]*domain.User, error) {
	var users []*domain.User
	cursor, err := ur.db.Find(ctx, filter, opts...)
	if err != nil {
		return nil, utils.DropNoDocs(err)
	}
	err = cursor.Decode(users)
	if err != nil {
		return nil, utils.DropNoDocs(err)
	}
	return users, nil
}

func (ur *UserRepo) InsertUser(ctx context.Context, data *domain.User) error {
	log.Println("insert rep")
	_, err := ur.db.InsertOne(ctx, data)
	if err != nil {
		log.Println("insert rep err", err)
		return err
	}
	return nil
}

func (ur *UserRepo) GetDataDriverByEmail(ctx context.Context, email string) (*domain.User, error) {
	return ur.getUserBase(ctx, bson.M{"email": email})
}

func (dr *UserRepo) InsertData(ctx context.Context, data *domain.TestDataRepo) error {
	_, err := dr.db.InsertOne(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
