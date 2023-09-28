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

type formAction struct {
	repo *repository.Repository
}

type IFormAction interface {
	InsertForm(ctx context.Context, log *logrus.Entry, data *soberdriverdata.FormToAction) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse])
	GetForms(ctx context.Context, log *logrus.Entry, data *soberdriverdata.GetFormsRequest) (resp utils.ServiceResponse[*soberdriverdata.FormsResponse])
}

func newFormAction(repo *repo.Repository) IFormAction {
	return &formAction{
		repo: repo,
	}
}

func (fa *formAction) InsertForm(ctx context.Context, log *logrus.Entry, data *soberdriverdata.FormToAction) (resp utils.ServiceResponse[*soberdriverdata.RegisterResponse]) {

	err := fa.repo.Form.InsertData(ctx, &domain.FormInput{
		UserID:      data.UserID,
		WhereFrom:   data.WhereFrom,
		WhereTo:     data.WhereTo,
		Description: data.Description,
		Time:        data.Time,
	})
	if err != nil {
		return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Insert form err", err)
	}

	return *resp.WriteData(&soberdriverdata.RegisterResponse{
		Status: "form registered",
	})
}

func (fa *formAction) GetForms(ctx context.Context, log *logrus.Entry, data *soberdriverdata.GetFormsRequest) (resp utils.ServiceResponse[*soberdriverdata.FormsResponse]) {
	var formsToReq []*soberdriverdata.FormToAction

	forms, err := fa.repo.Form.GetFormsById(ctx, data.UserID)
	if err != nil {
		return *resp.WriteError(http.StatusInternalServerError, utils.InternalServer, "Get forms err", err)
	}

	for _, form := range forms {
		formsToReq = append(formsToReq, conv.FormToPkg(form))
	}

	return *resp.WriteData(&soberdriverdata.FormsResponse{
		Forms: formsToReq,
	})
}
