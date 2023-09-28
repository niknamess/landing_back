package v1

import (
	soberdriverdata "sober_driver/pkg/domain"

	"github.com/ArenAzibekyan/logrus-helper/logger"
	"github.com/gin-gonic/gin"
)

// @Summary регистрация пользователя
// @Schemes
// @Description регистрация с проверкой
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param request body soberdriverdata.RegisterRequest true "Ввести данные"
// @Success 200 {object} soberdriverdata.RegisterResponse
// @Router /api/v1/register [POST]
func (hr *Handler) register(c *gin.Context) {
	log := logger.Default()
	data := new(soberdriverdata.RegisterRequest)
	err := c.ShouldBindJSON(data)
	if err != nil {
		log.Println(err)
		return
	}
	resp := hr.Service.UserAction.Register(c.Request.Context(), log, data)
	resp.Send(c, log, "Register")
}

// @Summary проверка регистрации пользователя
// @Schemes
// @Description проверка регистрации пользователя
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param request body soberdriverdata.RegisterCheckRequest true "Ввести Email"
// @Success 200 {object} soberdriverdata.RegisterResponse
// @Router /api/v1/register/check [GET]
func (hr *Handler) checkRegister(c *gin.Context) {
	log := logger.Default()
	data := new(soberdriverdata.RegisterCheckRequest)
	err := c.ShouldBindJSON(data)
	if err != nil {

		return
	}

	resp := hr.Service.UserAction.CheckRegister(c.Request.Context(), log, data)
	resp.Send(c, log, "CheckRegister")
}
