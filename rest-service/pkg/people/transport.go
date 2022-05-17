package people

import (
	"context"
	"encoding/json"
	"net/http"

	kitTransport "github.com/go-kit/kit/transport/http"
	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"

	"github.com/stackpath/backend-developer-tests/rest-service/pkg/common"
)

func AddRouteToHandler(endpoints Endpoints, r *mux.Router, options []kitTransport.ServerOption) {
	r.Methods(http.MethodGet).Path("/people").Handler(
		kitTransport.NewServer(
			endpoints.GetPeople,
			decodeEmptyRequest,
			common.EncodeResponse,
			options...,
		),
	)

	r.Methods(http.MethodGet).Path("/people/{personID}").Handler(
		kitTransport.NewServer(
			endpoints.GetPersonByID,
			decodeGetPersonByIDRequest,
			common.EncodeResponse,
			options...,
		),
	)

	r.Methods(http.MethodPost).Path("/people").Handler(
		kitTransport.NewServer(
			endpoints.FindPerson,
			decodeGetPeopleRequest,
			common.EncodeResponse,
			options...,
		),
	)
}
func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}

func decodeGetPeopleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var getPeopleRequest GetPeopleRequest
	if err := json.NewDecoder(r.Body).Decode(&getPeopleRequest); err != nil {
		return nil, common.NewMSError(common.ErrJSONInvalid.Error(), common.BadRequest, common.DecodeGetPeopleRequest, common.TransportLevel, err)
	}

	if !((getPeopleRequest.FirstName != "" && getPeopleRequest.LastName != "") || getPeopleRequest.Phone != "") {
		return nil, common.NewMSError(common.ErrInvalidRequest.Error(), common.BadRequest, common.DecodeGetPeopleRequest, common.TransportLevel, common.ErrInvalidRequest)
	}

	return getPeopleRequest, nil
}

func decodeGetPersonByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	str := vars["personID"]
	id, err := uuid.FromString(str)
	if err != nil {
		return nil, common.NewMSError(common.ErrInvalidUUID.Error(), common.BadRequest, common.DecodeGetPersonByIDRequest, common.TransportLevel, err)
	}

	req := GetPersonByIDRequest{
		ID: id,
	}
	return req, nil
}
