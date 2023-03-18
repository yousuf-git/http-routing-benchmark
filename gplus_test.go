package http_routing_benchmark

import (
	"testing"
)

// http://developer.github.com/v3/
var gplusAPI = []route{
	// People
	{"GET", "/people/:userId"},
	{"GET", "/people"},
	{"GET", "/activities/:activityId/people/:collection"},
	{"GET", "/people/:userId/people/:collection"},
	{"GET", "/people/:userId/openIdConnect"},

	// Activities
	{"GET", "/people/:userId/activities/:collection"},
	{"GET", "/activities/:activityId"},
	{"GET", "/activities"},

	// Comments
	{"GET", "/activities/:activityId/comments"},
	{"GET", "/comments/:commentId"},

	// Moments
	{"POST", "/people/:userId/moments/:collection"},
	{"GET", "/people/:userId/moments/:collection"},
	{"DELETE", "/moments/:id"},
}

func BenchmarkShift_GPlusAll(b *testing.B) {
	router := prepareShift(gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}

func BenchmarkGin_GPlusAll(b *testing.B) {
	router := prepareGin(gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}

func BenchmarkEcho_GPlusAll(b *testing.B) {
	router := prepareEcho(gplusAPI)
	benchmarkRoutes(b, router, gplusAPI)
}
