package api

import (
		"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type jsonResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    any 		`json:"data"`
}

var validate = validator.New()

func (app *Config) validateBody(c *gin.Context, data any) error {
	if err := c.BindJSON(data); err != nil {
		return err
	}
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}



func (app *Config) WriteJSON(c *gin.Context, status int, data any) {
	var payload jsonResponse
	payload.Error = false
	payload.Message = "success"
	payload.Data = data
	c.JSON(status, payload)
}

func (app *Config) errorJSON(c *gin.Context, err error, status ...int) {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}
	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()
	payload.Data = nil
	c.JSON(statusCode, payload)
}