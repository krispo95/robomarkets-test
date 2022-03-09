package usecase

import (
	"robomarkets-test/internal/repository"
	"testing"
)

func Benchmark_TestGeoUseCase_FindLocationByName(b *testing.B) {
	repo, err := repository.NewGeoRepository("../../docs/geobase.dat")
	if err != nil {
		panic(err)
	}
	uc := NewGeoUsecase(repo)
	loc := uc.FindLocationByName("cit_Osumi")
	if loc == nil || loc.City != "cit_Osumi" {
		panic("not valid answer")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uc.FindLocationByName("cit_Osumi")
	}
}

func Benchmark_TestGeoUseCase_FindLocationByIP(b *testing.B) {
	repo, err := repository.NewGeoRepository("../../docs/geobase.dat")
	if err != nil {
		panic(err)
	}
	uc := NewGeoUsecase(repo)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		uc.FindLocationByIP(repository.ParseUint32([]byte("123.234.123.234")))
	}
}
