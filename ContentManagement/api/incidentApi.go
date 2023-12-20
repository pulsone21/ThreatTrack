package api

import (
	"ContentManagement/api/models/incident"
	"ContentManagement/api/stores"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type IncidentApi struct {
	Store *stores.IncidentStore
}

type CreaIncReq struct {
	Name         string `json:"name"`
	Severity     string `json:"severity"`
	IncidentType int    `json:"incidentType"`
}

type CreIncTypeReq struct {
	Name string `json:"name"`
}
type GetIncTypeReq struct {
	Id   int    `json:"_id"`
	Name string `json:"name"`
}

func CreateIncidentApi(db *sql.DB, ser *ApiServer) *IncidentApi {
	iA := &IncidentApi{
		Store: stores.NewIncidentStore(db),
	}
	iA.createHandles(ser)
	return iA
}

func (s *IncidentApi) createHandles(ser *ApiServer) {

	inc_sR := ser.Router.PathPrefix("/incident").Subrouter()
	inc_sR.HandleFunc("", createHandleFunc(s.api_GetAllIncidents)).Methods(http.MethodGet)
	inc_sR.HandleFunc("", createHandleFunc(s.api_CreateIncident)).Methods(http.MethodPost)
	inc_sR.HandleFunc("/{id}", createHandleFunc(s.api_GetIncidentByID)).Methods(http.MethodGet)
	inc_sR.HandleFunc("/{id}", createHandleFunc(s.api_UpdateIncident)).Methods(http.MethodPut)
	inc_sR.HandleFunc("/{id}", createHandleFunc(s.api_DeleteIncident)).Methods(http.MethodDelete)

	type_sR := ser.Router.PathPrefix("/incidenttype").Subrouter()
	type_sR.HandleFunc("", createHandleFunc(s.api_GetAllIncidentTypes)).Methods(http.MethodGet)
	type_sR.HandleFunc("", createHandleFunc(s.api_CreateIncidentType)).Methods(http.MethodPost)
	type_sR.HandleFunc("/{id}", createHandleFunc(s.api_DeleteIncidentType)).Methods(http.MethodDelete)
	type_sR.HandleFunc("/{id}", createHandleFunc(s.api_GetIncidentTypeById)).Methods(http.MethodGet)

}

func (s *IncidentApi) api_GetAllIncidents(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	incs, err := s.Store.GetAllIncidents()
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	return NewApiResponse(http.StatusOK, r.RequestURI, incs), nil
}

func (s *IncidentApi) api_CreateIncident(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	fmt.Println("Request for Creating Incident")
	var IncReq CreaIncReq
	json.NewDecoder(r.Body).Decode(&IncReq)
	fmt.Println(IncReq)

	iT, err := s.Store.GetIncidentTypeBy("id", fmt.Sprintf("%v", IncReq.IncidentType))
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}

	inc := incident.NewIncident(IncReq.Name, incident.IncidentSeverity(IncReq.Severity), *iT)
	err = s.Store.CreateIncident(inc)
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{inc}), nil
}

func (s *IncidentApi) api_GetIncidentByID(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	id := mux.Vars(r)["id"]
	inc, err := s.Store.GetIncidentByID(id)
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{inc}), nil
}

func (s *IncidentApi) api_UpdateIncident(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	return nil, &ApiError{
		error:      fmt.Errorf("not implemented"),
		StatusCode: http.StatusNotImplemented,
		RequestUrl: r.RequestURI,
	}
}

func (s *IncidentApi) api_DeleteIncident(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	id := mux.Vars(r)["id"]
	if err := s.Store.DeleteIncident(id); err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	res_map := make(map[string]any)
	res_map["Message"] = fmt.Sprintf(`Incident with ID: %s was deleted`, id)
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{res_map}), nil
}

func (s *IncidentApi) api_GetAllIncidentTypes(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	fmt.Println("Get all Incident Types")
	incs, err := s.Store.GetAllIncidentTypes()
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	fmt.Sprintln(incs)
	return NewApiResponse(http.StatusOK, r.RequestURI, incs), nil
}

func (s *IncidentApi) api_CreateIncidentType(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	fmt.Println("Creating new Incident Type")
	var iTR CreIncTypeReq
	json.NewDecoder(r.Body).Decode(&iTR)
	id, err := s.Store.CreateIncidentType(iTR.Name)
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	iT := &incident.IncidentType{
		Id:   id,
		Name: iTR.Name,
	}
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{iT}), nil
}

func (s *IncidentApi) api_GetIncidentTypeById(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	id := mux.Vars(r)["id"]
	iT, err := s.Store.GetIncidentTypeBy("id", id)
	if err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{iT}), nil
}

func (s *IncidentApi) api_DeleteIncidentType(w http.ResponseWriter, r *http.Request) (*ApiResponse, *ApiError) {
	id := mux.Vars(r)["id"]
	if err := s.Store.DeleteIncidentType(id); err != nil {
		return nil, InternalServerError(err, r.RequestURI)
	}
	res_map := make(map[string]any)
	res_map["Message"] = fmt.Sprintf(`IncidentType with ID: %s was deleted`, id)
	return NewApiResponse(http.StatusOK, r.RequestURI, []interface{}{res_map}), nil
}
