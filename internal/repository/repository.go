package repository

import (
	"os"
	"robomarkets-test/internal/entity"
)

type Repository interface {
	GetLocationByLocationIndex(locID entity.City) *entity.Location
	GetCityById(id int) entity.City
	GetLocationById(id int) *entity.Location
	GetRecordsNumber() int32
	GetRangeById(id int) *entity.Range
}

type GeoRepo struct {
	header    entity.Header
	ranges    []entity.Range
	cities    []entity.City
	locations []entity.Location
}

func NewGeoRepository(filePath string) (*GeoRepo, error) {
	geoRepo := GeoRepo{}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	geoRepo.header, err = ParseHeader(file)
	if err != nil {
		return nil, err
	}

	geoRepo.ranges, err = ParseRanges(file, int(geoRepo.GetRecordsNumber()))
	if err != nil {
		return nil, err
	}

	geoRepo.locations, err = ParseLocations(file, int(geoRepo.GetRecordsNumber()))
	if err != nil {
		return nil, err
	}

	geoRepo.cities, err = ParseCities(file, int(geoRepo.GetRecordsNumber()))
	if err != nil {
		return nil, err
	}

	return &geoRepo, nil
}

func (g *GeoRepo) GetLocationByLocationIndex(locID entity.City) *entity.Location {
	id := locID / entity.CitySizeBts
	return &g.locations[id]
}

func (g *GeoRepo) GetCityById(id int) entity.City {
	return g.cities[id]
}

func (g *GeoRepo) GetLocationById(id int) *entity.Location {
	return &g.locations[id]
}

func (g *GeoRepo) GetRecordsNumber() int32 {
	return g.header.Records
}

func (g *GeoRepo) GetRangeById(id int) *entity.Range {
	return &g.ranges[id]
}
