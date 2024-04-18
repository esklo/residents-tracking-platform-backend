package model

type GeoPoint struct {
	Lat float64
	Lon float64
}

type GeoBuilding struct {
	Street      string
	HouseNumber string
	Geometry    []GeoPoint
}

type GeoDistrict struct {
	Id   int64
	Name string
}
