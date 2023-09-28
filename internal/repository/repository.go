package repository

import (
	"context"
	"sober_driver/internal/domain"
	mongorepo "sober_driver/internal/repository/mongodb"

	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	Form IForm
	User IUser
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		User: mongorepo.NewUserRepo(db),
		Form: mongorepo.NewFormRepo(db),
	}
}

type IForm interface {
	InsertData(ctx context.Context, data *domain.FormInput) error
	GetFormsById(ctx context.Context, userID string) ([]*domain.FormInput, error)
}

type IUser interface {
	InsertUser(ctx context.Context, data *domain.User) error
	GetDataDriverByEmail(ctx context.Context, email string) (*domain.User, error)
	InsertData(ctx context.Context, data *domain.TestDataRepo) error
}
