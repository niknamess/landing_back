package utils

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"
)

type ServiceResponse[Data any] struct {
	Response Data          `json:"response"`
	Error    *ServiceError `json:"error"`
}

type ServiceError struct {
	Description    string      `json:"description"`
	Reason         ErrorReason `json:"reason"`
	HttpStatusCode int         `json:"-"`
	HttpStatus     int         `json:"-"`
	Err            any         `json:"-"`
}

type ErrorReason string

const (
	NotFound       ErrorReason = "not found"
	BadRequest     ErrorReason = "bad request"
	InternalServer ErrorReason = "internal server"
	AccessDenied   ErrorReason = "access denied"
	TimeOut        ErrorReason = "time out"
)

func (resp *ServiceResponse[Data]) WriteData(data Data) *ServiceResponse[Data] {
	if resp == nil {
		resp = &ServiceResponse[Data]{
			Response: data,
		}
	} else {
		resp.Response = data
	}

	return resp
}

func (resp *ServiceResponse[Data]) WriteError(httpStatus int, reason ErrorReason, desc string, err any) *ServiceResponse[Data] {
	er := NewServiceError(httpStatus, reason, desc, err)

	if resp == nil {
		resp = &ServiceResponse[Data]{
			Error: er,
		}
	} else {
		resp.Error = er
	}

	return resp
}

func (resp *ServiceResponse[Data]) Send(c *gin.Context, log *logrus.Entry, method string) {
	if resp.Error == nil {
		resp.send(c, http.StatusOK)
	} else {
		log.WithField("code", resp.Error.Reason).Error(resp.Error.Err)
		resp.send(c, resp.Error.HttpStatus)
	}

}

func NewServiceError(httpStatus int, reason ErrorReason, desc string, err any) *ServiceError {
	return &ServiceError{
		Description: desc,
		Reason:      reason,
		HttpStatus:  httpStatus,
		Err:         err,
	}
}

func (resp *ServiceResponse[_]) send(c *gin.Context, httpStatus int) {
	c.JSON(httpStatus, resp)
}

func CreateUlid() ulid.ULID {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	id, _ := ulid.New(ms, entropy)

	return id
}
