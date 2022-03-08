package usecase

import (
	"geobase/internal/entity"
	"geobase/internal/repository"
	"strings"
)

type Usecase interface {
	FindLocationByName(name string) *entity.Location
	FindLocationByIP(ip uint32) *entity.Location
}

type GeoUseCase struct {
	repository repository.Repository
}

func NewGeoUsecase(repo repository.Repository) GeoUseCase {
	return GeoUseCase{
		repository: repo,
	}
}

func (g GeoUseCase) FindLocationByName(name string) *entity.Location {
	var left, right = 0, int(g.repository.GetRecordsNumber())
	for left <= right {
		mid := (left + right) / 2
		location := g.repository.GetLocationByLocationIndex(g.repository.GetCityById(mid))
		if location == nil {
			panic("city is nil")
		}
		switch strings.Compare(name, location.City) {
		case 0:
			return location
		case 1:
			left = mid + 1
		case -1:
			right = mid - 1
		}
	}
	return nil
}

func (g GeoUseCase) FindLocationByIP(ip uint32) *entity.Location {
	var left, right = 0, int(g.repository.GetRecordsNumber())
	var foundRange *entity.Range
	for left <= right {
		mid := (left + right) / 2
		rng := g.repository.GetRangeById(mid)
		if rng == nil {
			panic("range is nil")
		}
		if rng.IsInside(ip) {
			foundRange = rng
			break
		}
		if rng.IsBigger(ip) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if foundRange != nil {
		return g.repository.GetLocationByLocationIndex(entity.City(foundRange.LocationIndex))
	}
	return nil
}
