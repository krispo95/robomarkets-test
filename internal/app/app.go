package app

import (
	"fmt"
	"geobase/config"
	"geobase/internal/api"
	"geobase/internal/repository"
	"geobase/internal/usecase"
)

// Run creates app
func Run(cfg *config.Config) {
	repo, err := repository.NewGeoRepository("./docs/geobase.dat")
	if err != nil {
		panic(err)
	}
	uc := usecase.NewGeoUsecase(repo)
	_ = uc
	fmt.Println(111)
	api.StartServer(cfg, uc)
}
