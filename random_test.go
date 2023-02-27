package http_routing_benchmark

import "testing"

var randomAPI = []route{
	{"GET", "/users/find"},
	{"GET", "/users/find/:name"},
	{"GET", "/users/:id/delete"},
	{"GET", "/users/:id/update"},
	{"GET", "/users/groups/:groupId/dump"},
	{"GET", "/users/groups/:groupId/export"},
	{"GET", "/users/delete"},
	{"GET", "/users/all/dump"},
	{"GET", "/users/all/export"},
	{"GET", "/users/any"},

	{"POST", "/search"},
	{"POST", "/search/go"},
	{"POST", "/search/go1.html"},
	{"POST", "/search/index.html"},
	{"POST", "/search/:q"},
	{"POST", "/search/:q/go"},
	{"POST", "/search/:q/go1.html"},
	{"POST", "/search/:q/:w/index.html"},

	{"PUT", "/src/:dest/invalid"},
	{"PUT", "/src/invalid"},
	{"PUT", "/src1/:dest"},
	{"PUT", "/src1"},

	{"PATCH", "/signal-r/:cmd/reflection"},
	{"PATCH", "/signal-r"},
	{"PATCH", "/signal-r/:cmd"},

	{"HEAD", "/query/unknown/pages"},
	{"HEAD", "/query/:key/:val/:cmd/single"},
	{"HEAD", "/query/:key"},
	{"HEAD", "/query/:key/:val/:cmd"},
	{"HEAD", "/query/:key/:val"},
	{"HEAD", "/query/unknown"},
	{"HEAD", "/query/untold"},

	{"CONNECT", "/questions/:index"},
	{"CONNECT", "/questions"},

	{"DELETE", "/graphql"},
	{"DELETE", "/graph"},
	{"DELETE", "/graphql/cmd"},
	{"DELETE", "/file"},
	{"DELETE", "/file/remove"},

	{"OPTIONS", "/hero-:name"},
}

func BenchmarkApe_RandomAll(b *testing.B) {
	router := prepareApe(randomAPI)
	benchmarkRoutes(b, router, randomAPI)
}

func BenchmarkGin_RandomAll(b *testing.B) {
	router := prepareGin(randomAPI)
	benchmarkRoutes(b, router, randomAPI)
}

func BenchmarkEcho_RandomAll(b *testing.B) {
	router := prepareEcho(randomAPI)
	benchmarkRoutes(b, router, randomAPI)
}
