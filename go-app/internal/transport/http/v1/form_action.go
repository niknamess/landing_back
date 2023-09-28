package v1

import (
	soberdriverdata "sober_driver/pkg/domain"

	"github.com/ArenAzibekyan/logrus-helper/logger"
	"github.com/gin-gonic/gin"
)

// @Summary записать в форму
// @Schemes
// @Description запись в форму инпуты
// @Tags action
// @Accept application/json
// @Produce application/json
// @Param request body soberdriverdata.FormToAction true "Форма для заполнения"
// @Success 200 {object} soberdriverdata.RegisterResponse
// @Router /api/v1/form/insert  [POST]
func (hr *Handler) insertForm(c *gin.Context) {
	log := logger.Default()
	data := new(soberdriverdata.FormToAction)
	err := c.ShouldBindJSON(data)
	if err != nil {
		return
	}

	resp := hr.Service.FormAction.InsertForm(c.Request.Context(), log, data)
	resp.Send(c, log, "InsertForm")
}

// @Summary получить формы по user id
// @Schemes
// @Description получение формы
// @Tags action
// @Accept application/json
// @Produce application/json
// @Param request body soberdriverdata.GetFormsRequest true "Ввести UserID"
// @Success 200 {object} soberdriverdata.FormsResponse
// @Router /api/v1/form  [GET]
func (hr *Handler) getForms(c *gin.Context) {
	log := logger.Default()
	data := new(soberdriverdata.GetFormsRequest)
	err := c.ShouldBindJSON(data)
	if err != nil {
		return
	}

	resp := hr.Service.FormAction.GetForms(c.Request.Context(), log, data)
	resp.Send(c, log, "GetForms")
}
