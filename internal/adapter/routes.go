package adapter

import (
	"github.com/alserov/parser/internal/usecase"
	"github.com/alserov/parser/internal/utils"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Controller struct {
	*Parser
}

func NewController(uc *usecase.Usecase) *Controller {
	return &Controller{
		&Parser{
			uc: uc.Parser,
		},
	}
}

func Setup(r *echo.Echo, ctrl *Controller) {
	r.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := r.Group("/v1")
	v1.Use(utils.HandleError)
	v1.GET("/parse", ctrl.Parse)
}
