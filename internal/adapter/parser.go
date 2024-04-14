package adapter

import (
	"fmt"
	"github.com/alserov/parser/internal/usecase"
	"github.com/alserov/parser/internal/utils"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type Parser struct {
	uc *usecase.Parser
}

// Parse godoc
// @Summary      Parse
// @Description  Parse page
// @Tags         parser
// @Accept       json
// @Produce      json
// @Param url query string true "page url"
// @Param tag query string true "tag name to be parsed"
// @Param incl query string false "value that has to be included"
//
// @Success      200  {array}   string
// @Failure      400  {object}  string
// @Failure      404  {object}  string
// @Failure      404  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/parse [get]
func (p *Parser) Parse(c echo.Context) error {
	url := c.QueryParam("url")
	if url == "" {
		return utils.NewError(utils.ErrBadRequest, "url parameter is required")
	}

	tag := c.QueryParam("tag")
	if url == "" {
		return utils.NewError(utils.ErrBadRequest, "tag parameter is required")
	}

	include := c.QueryParam("incl")
	if url == "" {
		return utils.NewError(utils.ErrBadRequest, "include parameter is required")
	}

	data, err := p.uc.Parse(c.Request().Context(), url, tag, strings.Split(include, ",")...)
	if err != nil {
		return fmt.Errorf("uc error: %w", err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"data": data,
	})
}
