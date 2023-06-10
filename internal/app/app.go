package app

import (
	"fmt"

	"github.com/kirakulakov/go-clean-arch/config"
	"github.com/kirakulakov/go-clean-arch/pkg/logger"
	"github.com/kirakulakov/go-clean-arch/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

}
