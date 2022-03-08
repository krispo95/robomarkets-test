package entity

type Header struct {
	Version         int32
	Name            string
	Timestamp       uint64
	Records         int32
	OffsetRanges    uint32
	OffsetCities    uint32
	OffsetLocations uint32
}
