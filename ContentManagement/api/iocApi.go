package api

import (
	"ContentManagement/api/models/ioc"
	"ContentManagement/api/stores"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type IocApi struct {
	Store *stores.IocStore
}
type CreatIocReq struct {
	Value       string   `json:"value"`
	IncidentIDs []string `json:"incidentIDs"`
	IOCType     int      `json:"IOCType"`
}

func CreateIocApi(db *sql.DB, ser *ApiServer) *IocApi {
	iA := &IocApi{
		Store: stores.NewIocStore(db),
	}
	iA.createHandles(ser)
	fmt.Println("Created IOC Api")
	return iA
}

func (api *IocApi) createHandles(ser *ApiServer) {

	inc_sR := ser.Router.PathPrefix("/ioc").Subrouter()
	inc_sR.HandleFunc("", createHandleFunc(api.api_GetAllIocs)).Methods(http.MethodGet)
	inc_sR.HandleFunc("", createHandleFunc(api.api_CreateIoc)).Methods(http.MethodPost)
	inc_sR.HandleFunc("/{id}", createHandleFunc(api.api_GetIocByID)).Methods(http.MethodGet)
	inc_sR.HandleFunc("/{id}", createHandleFunc(api.api_UpdateIoc)).Methods(http.MethodPut)
	inc_sR.HandleFunc("/{id}", createHandleFunc(api.api_DeleteIoc)).Methods(http.MethodDelete)

	type_sR := ser.Router.PathPrefix("/ioctype").Subrouter()
	type_sR.HandleFunc("", createHandleFunc(api.api_GetAllIocTypes)).Methods(http.MethodGet)
	type_sR.HandleFunc("", createHandleFunc(api.api_CreateIocType)).Methods(http.MethodPost)
	type_sR.HandleFunc("/{id}", createHandleFunc(api.api_DeleteIocType)).Methods(http.MethodDelete)
	type_sR.HandleFunc("/{id}", createHandleFunc(api.api_GetIocTypeById)).Methods(http.MethodGet)

	sev_sR := ser.Router.PathPrefix("/relations").Subrouter()
	sev_sR.HandleFunc("", createHandleFunc(api.api_GetAllIocIncidents)).Methods(http.MethodGet)
	sev_sR.HandleFunc("/?id={id}&type={type}", createHandleFunc(api.api_DeleteRelationBy)).Methods(http.MethodDelete)
	sev_sR.HandleFunc("/?id={id}&type={type}", createHandleFunc(api.api_GetIocIncidentBy)).Methods(http.MethodPost)
}

// API FUNCTIONS FOR IOCS

func (s *IocApi) api_GetAllIocs(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Requesting All IOCs API Handler")
	iocs, err := s.Store.GetAllIocs()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, iocs)
}

func (s *IocApi) api_CreateIoc(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Request for Creating IOC")
	var iocReq CreatIocReq
	json.NewDecoder(r.Body).Decode(&iocReq)
	iT, err := s.Store.GetIocTypeBy("id", fmt.Sprint(iocReq.IOCType))
	if err != nil {
		fmt.Println("error in get IOCType reqeust")
		return err
	}
	ioc := ioc.NewIOC(iocReq.Value, *iT)
	_, err = s.Store.CreateIOC(*ioc, iocReq.IncidentIDs)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, ioc)
}

func (s *IocApi) api_GetIocByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	ioc, err := s.Store.GetGetIocById(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, ioc)
}

func (s *IocApi) api_UpdateIoc(w http.ResponseWriter, r *http.Request) error {
	return fmt.Errorf("not implemented")
}

func (s *IocApi) api_DeleteIoc(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := s.Store.DeleteIoc(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, fmt.Sprintf("Deleted IOC with ID: %s", id))
}

// API FUNCTIONS FOR IOC TYPES

func (s *IocApi) api_GetAllIocTypes(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Get all ioc Types")
	iTs, err := s.Store.GetAllIocTypes()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, iTs)
}

func (s *IocApi) api_CreateIocType(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("Creating new Incident Type")
	var iTR CreIncTypeReq
	err := json.NewDecoder(r.Body).Decode(&iTR)
	if err != nil {
		return err
	}
	iT, err := s.Store.CreateIOCType(iTR.Name)
	if err != nil {
		return nil
	}
	return writeJSON(w, http.StatusOK, iT)
}

func (s *IocApi) api_GetIocTypeById(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	iT, err := s.Store.GetIocTypeBy("id", id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, iT)
}

func (s *IocApi) api_DeleteIocType(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	err := s.Store.DeleteIOCType(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, fmt.Sprintf("Deleted IOCType with ID: %s", id))
}

// API FUNCTIONS FOR IOC_INCIDENT RELEATIONS

func (s *IocApi) api_GetAllIocIncidents(w http.ResponseWriter, r *http.Request) error {
	res, err := s.Store.GetAllRelations()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, res)
}

func (s *IocApi) api_GetIocIncidentBy(w http.ResponseWriter, r *http.Request) error {
	qval := r.URL.Query()
	id := qval.Get("id")
	ty := qval.Get("type")
	res, err := s.Store.GetRelationBy(ty, id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, res)
}

func (s *IocApi) api_DeleteRelationBy(w http.ResponseWriter, r *http.Request) error {
	qval := r.URL.Query()
	id := qval.Get("id")
	ty := qval.Get("type")
	err := s.Store.DeleteRelationsBy(ty, id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, fmt.Sprintf("Delete all Relations from %s with id: %s", ty, id))

}
