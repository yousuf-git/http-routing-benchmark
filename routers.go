package http_routing_benchmark

import (
	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
	"github.com/yousuf64/ape"
	"net/http"
	"testing"
)

type route struct {
	method, path string
}

func prepareApe(routes []route, opts ...func(router *ape.Router)) http.Handler {
	h := func(w http.ResponseWriter, r *http.Request, route ape.Route) error {
		return nil
	}

	r := ape.New()
	for _, opt := range opts {
		opt(r)
	}
	for _, route := range routes {
		r.Map([]string{route.method}, route.path, h)
	}
	return r.Serve()
}

func prepareGin(routes []route, opts ...func(router *gin.Engine)) http.Handler {
	h := func(context *gin.Context) {}

	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	for _, opt := range opts {
		opt(r)
	}
	for _, route := range routes {
		r.Handle(route.method, route.path, h)
	}
	return r
}

func prepareEcho(routes []route) http.Handler {
	h := func(context echo.Context) error {
		return nil
	}

	r := echo.New()
	for _, route := range routes {
		r.Add(route.method, route.path, h)
	}
	return r
}

func benchmarkRoutes(b *testing.B, router http.Handler, routes []route) {
	requests := make([]*http.Request, 0, len(routes))
	for _, route := range routes {
		req, _ := http.NewRequest(route.method, route.path, nil)
		requests = append(requests, req)
	}

	w := newFakeWriter()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, req := range requests {
			router.ServeHTTP(w, req)
		}
	}
}

type fakeWriter struct {
	headers http.Header
}

func newFakeWriter() *fakeWriter {
	return &fakeWriter{
		http.Header{},
	}
}

func (m *fakeWriter) Header() (h http.Header) {
	return m.headers
}

func (m *fakeWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *fakeWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *fakeWriter) WriteHeader(int) {}
