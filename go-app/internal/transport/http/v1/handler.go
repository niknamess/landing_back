package v1

import (
	"sober_driver/internal/service"

	"github.com/gin-gonic/gin"
)

const (
	URLV1Registration      = "/register"
	URLV1Testreq           = "/register/:id"
	URLV1CheckRegistration = "/register/check"
	URLV1InsertForm        = "/form/insert"
	URLV1GetForms          = "/form"

	URLV1Login = "/login"

	/////Test
	URLV1Ruchka = "/test/test"
	//test
	URLV1Allalbums = "/albums"
	URLV1AlbumByID = "/albums/:id"
	URLV1Albums    = "/albums"
)

type Handler struct {
	Service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (hr *Handler) Init(api *gin.RouterGroup) {

	api.POST(URLV1Registration, hr.register)
	api.POST(URLV1InsertForm, hr.insertForm).Use()
	api.GET(URLV1CheckRegistration, hr.checkRegister).Use()
	api.GET(URLV1GetForms, hr.getForms).Use()

}
