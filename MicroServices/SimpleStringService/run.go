package SimpleStringService

import (
	//"context"
	//"encoding/json"
	//"errors"
	"log"
	"net/http"
	//"strings"

	//"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)


func RunService() {
	svc := stringService{}

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
	log.Fatal(http.ListenAndServe(":8080", nil))
}
