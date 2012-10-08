package main

import (
	"net/http"
	"strconv"
	"time"
)

var observationcounter Counter

type Observation struct {
	Id        int
	Timestamp time.Time
	Celsius   int
	Point     *ObservationPoint
}

var observations = make([]*Observation, 0)

func GetObservations(page int) []*Observation {
	// non-trivial case (last page or later)
	if len(observations) < page*5 {
		if len(observations) < 5 {
			return observations
		}
		// Trying to get a page that's out of bounds
		// XXX: Still messes prev/next buttons
		if len(observations) < (page-1)*5 {
			return observations[:5]
		}
		// Just return whatever is on the last page (n<=5).
		return observations[(page-1)*5:]
	}
	// trivial case, non-last page.
	return observations[(page-1)*5 : page*5]
}

// Struct that holds information we need to fill the template
type ListObservationHelper struct {
	Page, MaxPage     int
	Observations      []*Observation
	ObservationPoints []*ObservationPoint
}

func newLOHelper(r *http.Request) *ListObservationHelper {
	var page int
	var err error

	r.ParseForm()

	maxpage := ((len(observations) - 1) / 5) + 1
	pagenum := r.FormValue("pageNumber")
	if pagenum == "" {
		page = maxpage
	} else {
		page, err = strconv.Atoi(pagenum)
		if err != nil || page <= 0 {
			page = 1
		}
	}

	l := &ListObservationHelper{
		Page:              page,
		MaxPage:           ((len(observations) - 1) / 5) + 1,
		Observations:      GetObservations(page),
		ObservationPoints: observationpoints,
	}

	// Fix so that the next/prev buttons wont bork when trying to reach page > maxpage.
	if l.Page > l.MaxPage {
		l.Page = 1
	}

	return l
}

func (l ListObservationHelper) NotLastPage() bool {
	if l.Page < l.MaxPage {
		return true
	}
	return false
}

func (l ListObservationHelper) NotFirstPage() bool {
	if l.Page != 1 {
		return true
	}
	return false
}

func (l ListObservationHelper) NextPage() int {
	return l.Page + 1
}

func (l ListObservationHelper) LastPage() int {
	return l.Page - 1
}
