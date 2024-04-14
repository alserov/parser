package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/alserov/parser/internal/adapter"
	"github.com/alserov/parser/internal/config"
	"github.com/alserov/parser/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func MustStart(cfg *config.Config) {

	uc := usecase.NewUsecase()

	ctrl := adapter.NewController(uc)

	r := echo.New()
	adapter.Setup(r, ctrl)

	ch := make(chan os.Signal, 1)
	go run(ch, r, fmt.Sprintf(":%d", cfg.Port))
	<-ch
}

func run(ch chan<- os.Signal, r *echo.Echo, addr string) {
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	if err := r.Start(addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic("failed to start server: " + err.Error())
	}
	r.Shutdown(context.Background())
}
