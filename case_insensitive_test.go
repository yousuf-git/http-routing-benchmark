package http_routing_benchmark

import (
	"github.com/gin-gonic/gin"
	"github.com/yousuf64/shift"
	"testing"
)

// Conflicting routes have been commented out due to Gin not supporting
// case-insensitive matching for conflicting routes.
var caseInsensitiveAPI = []route{
	{"GET", "/users/find"},
	{"GET", "/users/find/:name"},
	//{"GET", "/users/:id/delete"},
	//{"GET", "/users/:id/update"},
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
	//{"POST", "/search/:q"},
	//{"POST", "/search/:q/go"},
	//{"POST", "/search/:q/go1.html"},
	//{"POST", "/search/:q/:w/index.html"},

	{"PUT", "/src/:dest/invalid"},
	//{"PUT", "/src/invalid"},
	{"PUT", "/src1/:dest"},
	{"PUT", "/src1"},

	{"PATCH", "/signal-r/:cmd/reflection"},
	{"PATCH", "/signal-r"},
	{"PATCH", "/signal-r/:cmd"},

	//{"HEAD", "/query/unknown/pages"},
	{"HEAD", "/query/:key/:val/:cmd/single"},
	{"HEAD", "/query/:key"},
	{"HEAD", "/query/:key/:val/:cmd"},
	{"HEAD", "/query/:key/:val"},
	//{"HEAD", "/query/unknown"},
	//{"HEAD", "/query/untold"},

	{"CONNECT", "/questions/:index"},
	{"CONNECT", "/questions"},

	{"DELETE", "/graphql"},
	{"DELETE", "/graph"},
	{"DELETE", "/graphql/cmd"},
	{"DELETE", "/file"},
	{"DELETE", "/file/remove"},

	{"OPTIONS", "/hero-:name"},
}

var caseInsensitiveTestAPI = []route{
	{"GET", "/UsErS/FiNd"},
	{"GET", "/uSErS/fINd/:name"},
	//{"GET", "/USeRs/:id/DeLeTE"},
	//{"GET", "/USERS/:id/UPDATE"},
	{"GET", "/users/GroUpS/:groupId/DuMp"},
	{"GET", "/USerS/grOuPS/:groupId/EXPORT"},
	{"GET", "/uSErs/DELete"},
	{"GET", "/uSeRs/ALL/duMP"},
	{"GET", "/USErs/aLL/exPORT"},
	{"GET", "/USERS/ANy"},

	{"POST", "/SEaRcH"},
	{"POST", "/sEarCh/Go"},
	{"POST", "/SeArCh/GO1.HTML"},
	{"POST", "/sEaRcH/IndEX.htML"},
	//{"POST", "/seArCh/:q"},
	//{"POST", "/sEaRcH/:q/gO"},
	//{"POST", "/SeArCh/:q/Go1.HTml"},
	//{"POST", "/sEArch/:q/:w/iNdEx.HtMl"},

	{"PUT", "/sRc/:dest/iNvaLiD"},
	//{"PUT", "/SRC/iNvAlId"},
	{"PUT", "/SRc1/:dest"},
	{"PUT", "/srC1"},

	{"PATCH", "/sIgNal-R/:cmd/reFlEctIon"},
	{"PATCH", "/SiGnaL-r"},
	{"PATCH", "/SIGNAL-R/:cmd"},

	//{"HEAD", "/QUERY/uNknOwN/PAges"},
	{"HEAD", "/qUErY/:key/:val/:cmd/sInGlE"},
	{"HEAD", "/QuERy/:key"},
	{"HEAD", "/qUeRy/:key/:val/:cmd"},
	{"HEAD", "/QUeRy/:key/:val"},
	//{"HEAD", "/QUERY/UnKnoWn"},
	//{"HEAD", "/qUeRy/unTolD"},

	{"CONNECT", "/QuesTioNs/:index"},
	{"CONNECT", "/qUEstIOnS"},

	{"DELETE", "/grAPhQl"},
	{"DELETE", "/gRapH"},
	{"DELETE", "/GRAPHQL/CMD"},
	{"DELETE", "/FILE"},
	{"DELETE", "/FiLE/ReMOVE"},

	{"OPTIONS", "/HeRO-:name"},
}

func BenchmarkShift_CaseInsensitiveAll(b *testing.B) {
	router := prepareShift(caseInsensitiveAPI, func(router *shift.Router) {
		router.UsePathCorrectionMatch(shift.WithRedirect())
	})
	benchmarkRoutes(b, router, caseInsensitiveTestAPI)
}

func BenchmarkGin_CaseInsensitiveAll(b *testing.B) {
	//defer func() {
	//	p := recover()
	//	if p != nil {
	//		b.Fatalf("gin panics\n")
	//	}
	//}()

	router := prepareGin(caseInsensitiveAPI, func(router *gin.Engine) {
		router.RedirectFixedPath = true
	})
	benchmarkRoutes(b, router, caseInsensitiveTestAPI)
}
