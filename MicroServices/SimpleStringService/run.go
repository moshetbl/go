package SimpleStringService

import (
	"os"
	"net/http"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)


func RunService() {
	logger := log.NewLogfmtLogger(os.Stderr)
	var svc StringService
	svc = stringService{}
	svc = loggingMiddleware{logger, svc}

	uppercaseHandler := httptransport.NewServer(
		makeUppercaseEndpoint(svc),
		decodeUppercaseRequest,
		encodeResponse,
	)

	lowercaseHandler := httptransport.NewServer(
		makeLowercaseEndpoint(svc),
		decodeLowercaseRequest,
		encodeResponse,
	)

	cutsubHandler := httptransport.NewServer(
		makeCutSubEndpoint(svc),
		decodeCutSubRequest,
		encodeResponse,
	)

	countHandler := httptransport.NewServer(
		makeCountEndpoint(svc),
		decodeCountRequest,
		encodeResponse,
	)


	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/lowercase", lowercaseHandler)
	http.Handle("/cutsub", cutsubHandler)
	http.Handle("/count", countHandler)
	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", nil))
}
