/*
Copyright 2018 Moshe Tubul
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package SimpleStringService

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-kit/kit/endpoint"
)

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(uppercaseRequest)
		v, err := svc.Uppercase(req.S)
		if err != nil {
			return uppercaseResponse{v, err.Error()}, nil
		}
		return uppercaseResponse{v, ""}, nil
	}
}

func makeLowercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(lowercaseRequest)
		v, err := svc.Lowercase(req.S)
		if err != nil {
			return lowercaseResponse{v, err.Error()}, nil
		}
		return lowercaseResponse{v, ""}, nil
	}
}

func makeCutSubEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(cutsubRequest)
		v, err := svc.CutSub(req.Src, req.Sub)
		if err != nil {
			return cutsubResponse{v, err.Error()}, nil
		}
		return cutsubResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(countRequest)
		v := svc.Count(req.S)
		return countResponse{v}, nil
	}
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request uppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeLowercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request lowercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCutSubRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request cutsubRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request countRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type uppercaseRequest struct {
	S string `json:"s"`
}

type uppercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type lowercaseRequest struct {
	S string `json:"s"`
}

type lowercaseResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type cutsubRequest struct {
	Src string `json:"src"`
	Sub string `json:"sub"`
}

type cutsubResponse struct {
	V   string `json:"v"`
	Err string `json:"err,omitempty"`
}

type countRequest struct {
	S string `json:"s"`
}

type countResponse struct {
	V int `json:"v"`
}
