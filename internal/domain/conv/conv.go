package conv

import (
	"sober_driver/internal/domain"
	soberdriverdata "sober_driver/pkg/domain"
)

func FioToDomain(data *soberdriverdata.Fio) *domain.Fio {
	return &domain.Fio{
		Name:       data.Name,
		Surname:    data.Surname,
		Patronymic: data.Patronymic,
	}
}

func FormToPkg(data *domain.FormInput) *soberdriverdata.FormToAction {
	return &soberdriverdata.FormToAction{
		UserID:      data.UserID,
		WhereFrom:   data.WhereFrom,
		WhereTo:     data.WhereTo,
		Description: data.Description,
		Time:        data.Time,
	}
}
