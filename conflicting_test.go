package http_routing_benchmark

import "testing"

var conflictingRoutesAPI = []route{
	{"GET", "/foo/:p"},
	{"GET", "/foo/:p/abc"},
	{"GET", "/foo/:p/def"},
	{"GET", "/foo/foo"},
	{"GET", "/foo/foo/:p"},

	{"POST", "/abc"},
	{"POST", "/abc/go"},
	{"POST", "/abc/go1.html"},
	{"POST", "/abc/index.html"},
	{"POST", "/abc/:q"},
	{"POST", "/abc/:q/go"},
	{"POST", "/abc/:q/go1.html"},
	{"POST", "/abc/:q/:w/index.html"},

	{"PUT", "/bar"},
	{"PUT", "/bar/abc"},
	{"PUT", "/bar/:p"},
	{"PUT", "/bar/:p/abc"},

	{"PATCH", "/bbb/bbb"},
	{"PATCH", "/bbb/:p/yyy"},
	{"PATCH", "/bbb/:p/:y"},

	{"HEAD", "/qqq/bbb"},
	{"HEAD", "/qqq/zzz/aaa"},
	{"HEAD", "/qqq/:a"},
	{"HEAD", "/qqq/:a/:b"},
	{"HEAD", "/qqq/:a/:b/:c"},
	{"HEAD", "/qqq/:a/:b/:c/aaa"},

	{"CONNECT", "/ooo/:a"},
	{"CONNECT", "/ooo/aaa"},

	{"DELETE", "/www/:a"},
	{"DELETE", "/www/:a/vvv"},
	{"DELETE", "/www/ccc"},
	{"DELETE", "/www/ccc/vvv"},

	{"OPTIONS", "/hhh-:a"},
	{"OPTIONS", "/hhh-mmm"},
}

func BenchmarkDune_ConflictingRoutesAll(b *testing.B) {
	router := prepareDune(conflictingRoutesAPI)
	benchmarkRoutes(b, router, conflictingRoutesAPI)
}

func BenchmarkGin_ConflictingRoutesAll(b *testing.B) {
	router := prepareGin(conflictingRoutesAPI)
	benchmarkRoutes(b, router, conflictingRoutesAPI)
}

func BenchmarkEcho_ConflictingRoutesAll(b *testing.B) {
	router := prepareEcho(conflictingRoutesAPI)
	benchmarkRoutes(b, router, conflictingRoutesAPI)
}
