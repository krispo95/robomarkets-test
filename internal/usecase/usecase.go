package usecase

import (
	"bytes"
	"encoding/binary"
	"net"
	"robomarkets-test/internal/entity"
	"robomarkets-test/internal/repository"
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

// FindLocationByName - binary searching location by city name
func (g GeoUseCase) FindLocationByName(name string) *entity.Location {
	var left, right = 0, int(g.repository.GetRecordsNumber())
	for left <= right {
		mid := (left + right) / 2
		// get location index by location id (in cities)
		location := g.repository.GetLocationByLocationIndex(g.repository.GetCityById(mid))
		if location == nil {
			//impossible case
			panic("city is nil")
		}
		// compare alphabetically
		// if name "bigger"(1) - search right, else (-1) - left
		// if equal (0) - found
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

// FindLocationByIP - binary searching location by ip
func (g GeoUseCase) FindLocationByIP(ip uint32) *entity.Location {
	var left, right = 0, int(g.repository.GetRecordsNumber())
	var foundRange *entity.Range
	for left <= right {
		mid := (left + right) / 2
		rng := g.repository.GetRangeById(mid)
		if rng == nil {
			//impossible case
			panic("range is nil")
		}
		// if ip is in range - found
		if rng.IsInside(ip) {
			foundRange = rng
			break
		}
		// if ip is bigger than range - search right, else - left
		if rng.IsBigger(ip) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if foundRange != nil {
		// get location index by location id (from ranges)
		return g.repository.GetLocationByLocationIndex(entity.City(foundRange.LocationIndex))
	}
	return nil
}

func Ip2Uint32(ip string) uint32 {
	var long uint32
	binary.Read(bytes.NewBuffer(net.ParseIP(ip).To4()), binary.BigEndian, &long)
	return long
}
