package router

import (
	"net/http"

	kitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/common"
	"github.com/stackpath/backend-developer-tests/rest-service/pkg/people"
)

type Endpoints struct {
	People people.Endpoints
}

type StatusResponse struct {
	Status string `json:"status"`
}

func NewHandler(endpoints Endpoints, mwf ...mux.MiddlewareFunc) http.Handler {
	r := mux.NewRouter()
	r.Use(mwf...)
	r.StrictSlash(true)

	options := []kitTransport.ServerOption{
		kitTransport.ServerErrorEncoder(common.EncodeError),
	}

	people.AddRouteToHandler(endpoints.People, r, options)

	return r
}
