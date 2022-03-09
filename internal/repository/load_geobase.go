package repository

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"math"
	"robomarkets-test/internal/entity"
)

func ParseHeader(file io.Reader) (entity.Header, error) {
	header := entity.Header{}

	headerBts := make([]byte, 60, 60)
	n, err := file.Read(headerBts)
	if n != 60 {
		return entity.Header{}, errors.New("header EOF")
	}
	if err != nil {
		return entity.Header{}, err
	}

	header.Version = ParseInt32(headerBts[:4])
	headerBts = headerBts[4:]

	header.Name = string(bytes.Trim(headerBts[:32], "\x00"))
	headerBts = headerBts[32:]

	header.Timestamp = ParseUint64(headerBts[:8])
	headerBts = headerBts[8:]

	header.Records = ParseInt32(headerBts[:4])
	headerBts = headerBts[4:]

	header.OffsetRanges = ParseUint32(headerBts[:4])
	headerBts = headerBts[4:]

	header.OffsetCities = ParseUint32(headerBts[:4])
	headerBts = headerBts[4:]

	header.OffsetLocations = ParseUint32(headerBts[:4])
	headerBts = headerBts[4:]
	return header, err
}

func ParseRanges(file io.Reader, count int) ([]entity.Range, error) {
	ips := make([]entity.Range, 0, count)

	for len(ips) < count {
		ipBts := make([]byte, 12, 12)
		n, err := file.Read(ipBts)
		if n != 12 {
			return nil, errors.New("ip EOF")
		}
		if err != nil {
			return nil, err
		}
		ip := entity.Range{}

		ip.IpFrom = ParseUint32(ipBts[:4])
		ipBts = ipBts[4:]

		ip.IpTo = ParseUint32(ipBts[:4])
		ipBts = ipBts[4:]

		ip.LocationIndex = ParseUint32(ipBts[:4])
		ipBts = ipBts[4:]
		ips = append(ips, ip)
	}

	return ips, nil
}

func ParseCities(file io.Reader, count int) ([]entity.City, error) {

	cities := make([]entity.City, 0, count)

	for len(cities) < count {
		bts := make([]byte, 4, 4)
		n, err := file.Read(bts)
		if n != 4 {
			return nil, errors.New("idx EOF")
		}
		if err != nil {
			return nil, err
		}

		id := ParseUint32(bts)
		cities = append(cities, entity.City(id))
	}

	return cities, nil
}

func ParseLocations(file io.Reader, count int) ([]entity.Location, error) {
	locations := make([]entity.Location, 0, count)

	for len(locations) < count {
		bts := make([]byte, 96, 96)
		n, err := file.Read(bts)
		if n != 96 {
			return nil, errors.New("addr EOF")
		}
		if err != nil {
			return nil, err
		}
		addr := entity.Location{}

		addr.Country = string(bytes.Trim(bts[:8], "\x00"))
		bts = bts[8:]

		addr.Region = string(bytes.Trim(bts[:12], "\x00"))
		bts = bts[12:]

		addr.Postal = string(bytes.Trim(bts[:12], "\x00"))
		bts = bts[12:]

		addr.City = string(bytes.Trim(bts[:24], "\x00"))
		bts = bts[24:]

		addr.Organization = string(bytes.Trim(bts[:32], "\x00"))
		bts = bts[32:]

		addr.Latitude = ParseFloat32(bts[:4])
		bts = bts[4:]

		addr.Longitude = ParseFloat32(bts[:4])
		bts = bts[4:]

		locations = append(locations, addr)
	}

	return locations, nil
}

func ParseInt32(bts []byte) int32 {
	return int32(binary.LittleEndian.Uint32(bts))
}

func ParseUint32(bts []byte) uint32 {
	return binary.LittleEndian.Uint32(bts)
}

func ParseUint64(bts []byte) uint64 {
	return binary.LittleEndian.Uint64(bts)
}

func ParseFloat32(bts []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(bts))
}
