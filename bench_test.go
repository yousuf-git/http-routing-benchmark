package http_routing_benchmark

import (
	"github.com/gin-gonic/gin"
	dune "github.com/yousuf-git/dune-project"
	"net/http"
	"testing"
)

type testCase struct {
	method string
	path   string
}

var api = map[string]string{
	"/users/find":                   http.MethodGet,
	"/users/find/:name":             http.MethodGet,
	"/users/:id/delete":             http.MethodGet,
	"/users/:id/update":             http.MethodGet,
	"/users/groups/:groupId/dump":   http.MethodGet,
	"/users/groups/:groupId/export": http.MethodGet,
	"/users/delete":                 http.MethodGet,
	"/users/all/dump":               http.MethodGet,
	"/users/all/export":             http.MethodGet,
	"/users/any":                    http.MethodGet,

	"/search":                  http.MethodPost,
	"/search/go":               http.MethodPost,
	"/search/go1.html":         http.MethodPost,
	"/search/index.html":       http.MethodPost,
	"/search/:q":               http.MethodPost,
	"/search/:q/go":            http.MethodPost,
	"/search/:q/go1.html":      http.MethodPost,
	"/search/:q/:w/index.html": http.MethodPost,

	"/src/:dest/invalid": http.MethodPut,
	"/src/invalid":       http.MethodPut,
	"/src1/:dest":        http.MethodPut,
	"/src1":              http.MethodPut,

	"/signal-r/:cmd/reflection": http.MethodPatch,
	"/signal-r":                 http.MethodPatch,
	"/signal-r/:cmd":            http.MethodPatch,

	"/query/unknown/pages":         http.MethodHead,
	"/query/:key/:val/:cmd/single": http.MethodHead,
	"/query/:key":                  http.MethodHead,
	"/query/:key/:val/:cmd":        http.MethodHead,
	"/query/:key/:val":             http.MethodHead,
	"/query/unknown":               http.MethodHead,
	"/query/untold":                http.MethodHead,

	"/questions/:index": http.MethodConnect,
	"/questions":        http.MethodConnect,

	"/graphql":     http.MethodDelete,
	"/graph":       http.MethodDelete,
	"/graphql/cmd": http.MethodDelete,

	"/file":        http.MethodDelete,
	"/file/remove": http.MethodDelete,

	"/hero-:name": http.MethodGet,
}

var apiTests = []testCase{
	{method: http.MethodGet, path: "/users/find"},
	{method: http.MethodGet, path: "/users/find/yousuf"},
	{method: http.MethodGet, path: "/users/john/delete"},
	{method: http.MethodGet, path: "/users/911/update"},
	{method: http.MethodGet, path: "/users/groups/120/dump"},
	{method: http.MethodGet, path: "/users/groups/230/export"},
	{method: http.MethodGet, path: "/users/delete"},
	{method: http.MethodGet, path: "/users/all/dump"},
	{method: http.MethodGet, path: "/users/all/export"},
	{method: http.MethodGet, path: "/users/any"},

	{method: http.MethodPost, path: "/search"},
	{method: http.MethodPost, path: "/search/go"},
	{method: http.MethodPost, path: "/search/go1.html"},
	{method: http.MethodPost, path: "/search/index.html"},
	{method: http.MethodPost, path: "/search/contact.html"},
	{method: http.MethodPost, path: "/search/ducks"},
	{method: http.MethodPost, path: "/search/gophers/go"},
	{method: http.MethodPost, path: "/search/nature/go1.html"},
	{method: http.MethodPost, path: "/search/generics/types/index.html"},

	{method: http.MethodPut, path: "/src/paris/invalid"},
	{method: http.MethodPut, path: "/src/invalid"},
	{method: http.MethodPut, path: "/src1/oslo"},
	{method: http.MethodPut, path: "/src1"},

	{method: http.MethodPatch, path: "/signal-r/protos/reflection"},
	{method: http.MethodPatch, path: "/signal-r"},
	{method: http.MethodPatch, path: "/signal-r/push"},
	{method: http.MethodPatch, path: "/signal-r/connect"},

	{method: http.MethodHead, path: "/query/unknown/pages"},
	{method: http.MethodHead, path: "/query/10/amazing/reset/single"},
	{method: http.MethodHead, path: "/query/911"},
	{method: http.MethodHead, path: "/query/99/sup/update-ttl"},
	{method: http.MethodHead, path: "/query/46/hello"},
	{method: http.MethodHead, path: "/query/unknown"},
	{method: http.MethodHead, path: "/query/untold"},

	{method: http.MethodConnect, path: "/questions/1001"},
	{method: http.MethodConnect, path: "/questions"},

	{method: http.MethodDelete, path: "/graphql"},
	{method: http.MethodDelete, path: "/graph"},
	{method: http.MethodDelete, path: "/graphql/cmd"},
	{method: http.MethodDelete, path: "/file"},
	{method: http.MethodDelete, path: "/file/remove"},

	{method: http.MethodGet, path: "/hero-goku"},
	{method: http.MethodGet, path: "/hero-thor"},
}

func generateRequests(tests []testCase) (requests []*http.Request) {
	for _, t := range tests {
		req, _ := http.NewRequest(t.method, t.path, nil)
		requests = append(requests, req)
	}
	return
}

func duneHandler(http.ResponseWriter, *http.Request, dune.Route) error {
	return nil
}

func BenchmarkDune(b *testing.B) {
	r := dune.New()

	for route, method := range api {
		r.Map([]string{method}, route, duneHandler)
	}

	svr := r.Serve()

	w := newMockWriter()
	requests := generateRequests(apiTests)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, req := range requests {
			svr.ServeHTTP(w, req)
		}
	}
}

func ginHandler(*gin.Context) {

}

func BenchmarkGin(b *testing.B) {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()

	for route, method := range api {
		e.Handle(method, route, ginHandler)
	}

	w := newMockWriter()
	requests := generateRequests(apiTests)

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, req := range requests {
			e.ServeHTTP(w, req)
		}
	}
}

type mockWriter struct {
	headers http.Header
}

func newMockWriter() *mockWriter {
	return &mockWriter{
		http.Header{},
	}
}

func (m *mockWriter) Header() (h http.Header) {
	return m.headers
}

func (m *mockWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockWriter) WriteHeader(int) {}
