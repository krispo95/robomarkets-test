package app

import (
	"robomarkets-test/config"
	"robomarkets-test/internal/api"
	"robomarkets-test/internal/repository"
	"robomarkets-test/internal/usecase"
)

// Run creates app
func Run(cfg *config.Config) {
	repo, err := repository.NewGeoRepository("./docs/geobase.dat")
	if err != nil {
		panic(err)
	}
	uc := usecase.NewGeoUsecase(repo)
	_ = uc
	api.StartServer(cfg, uc)
}
