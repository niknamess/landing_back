package service

import (
	"context"
	"net/http"
	"sober_driver/internal/domain"
	"sober_driver/internal/domain/conv"
	"sober_driver/internal/repository"
	repo "sober_driver/internal/repository"
	soberdriverdata "sober_driver/pkg/domain"
	"sober_driver/pkg/utils"

	"github.com/sirupsen/logrus"
)

type userAction struct {
	repo *repository.Repository
}

type IUserAction interface {
	Register(ctx context.Context, log *logrus.Entry, data *soberdriverdata.RegisterRequest) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse])
	CheckRegister(ctx context.Context, log *logrus.Entry, data *soberdriverdata.RegisterCheckRequest) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse])
	InsertDataTest(ctx context.Context, log *logrus.Entry, data *domain.TestDataService) (resp utils.ServiceResponse[string])
}

func newUserAction(repo *repo.Repository) IUserAction {
	return &userAction{
		repo: repo,
	}
}

func (ua *userAction) Register(ctx context.Context, log *logrus.Entry, data *soberdriverdata.RegisterRequest) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse]) {
	log.Println("2")

	user, err := ua.repo.User.GetDataDriverByEmail(ctx, data.Email)
	if err != nil {
		return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Get user by email err", err)
	}

	if user == nil {
		err := ua.repo.User.InsertUser(ctx, &domain.User{
			ID:       utils.CreateUlid().String(),
			Fio:      conv.FioToDomain(data.Fio),
			Email:    data.Email,
			Password: data.Password,
			Driver:   data.Driver,
			Customer: data.Customer,
		})
		if err != nil {
			log.Println("not insert", err)
			return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Can't insert user", err)
		}
		log.Println("InsertUser")

		return *resp.WriteData(&soberdriverdata.RegisterResponse{
			Status: "user registered",
		})
	}

	return *resp.WriteData(&soberdriverdata.RegisterResponse{
		Status: "user already registered",
	})
}

func (ua *userAction) CheckRegister(ctx context.Context, log *logrus.Entry, data *soberdriverdata.RegisterCheckRequest) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse]) {
	user, err := ua.repo.User.GetDataDriverByEmail(ctx, data.Email)
	if err != nil {
		return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Get user by email err", err)
	}

	if user == nil {
		return *resp.WriteData(&soberdriverdata.RegisterResponse{
			Status: "user not registered",
		})
	}

	return *resp.WriteData(&soberdriverdata.RegisterResponse{
		Status: "user already registered",
	})
}

func (ua *userAction) InsertDataTest(ctx context.Context, log *logrus.Entry, data *domain.TestDataService) (resp utils.ServiceResponse[string]) {
	err := ua.repo.User.InsertData(ctx, &domain.TestDataRepo{
		ID:         data.ID,
		DriverName: data.DriverName,
		CarsName:   data.CarsName,
	})
	if err != nil {
		return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Insert data err", err)
	}

	return *resp.WriteData("Success")
}
