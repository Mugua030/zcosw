package zencos

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

type App struct {
}

func NewApp() *App {
	return &App{}
}

func (a *App) Run() {
}

func (a *App) Stop() {
}
