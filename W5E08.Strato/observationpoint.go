package main

import (
	"math"
	"net/http"
	"strings"
)

var pointcounter Counter

type ObservationPoint struct {
	Id           int
	Name         string
	Latitude     float64 // -90 90
	Longitude    float64 // -180 180
	Observations []*Observation
}

var observationpoints = make([]*ObservationPoint, 0)

type ListPointHelper struct {
	ObservationPoints []*ObservationPoint
}

func newPointHelper() *ListPointHelper {
	return &ListPointHelper{ObservationPoints: observationpoints}
}

func (o *ObservationPoint) IsValid(r *http.Request) bool {
	if o.Name == "" {
		return false
	}

	// Using ',' instead of '.'
	if strings.Contains(r.FormValue("Longitude"), ",") ||
		strings.Contains(r.FormValue("Latitude"), ",") {

		return false
	}

	// Probably got non-numerical data.
	if (!strings.HasPrefix(r.FormValue("Latitude"), "0") && o.Latitude == 0.0) ||
		(!strings.HasPrefix(r.FormValue("Longitude"), "0") && o.Longitude == 0.0) {
		return false
	}

	// Not valid coordinate
	if math.Abs(o.Latitude) > 90 || math.Abs(o.Longitude) > 180 {
		return false
	}

	return true

}
