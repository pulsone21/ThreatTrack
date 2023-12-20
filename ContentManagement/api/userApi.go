package api

import (
	"ContentManagement/api/models/user"
	"ContentManagement/api/stores"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type CreateUserReq struct {
	FirstName string
	LastName  string
	Email     string
}

type UserApi struct {
	Store stores.UserStore
}

func CreateUserApi(db *sql.DB, serv *ApiServer) *UserApi {
	uA := &UserApi{
		Store: *stores.NewUserStore(db),
	}
	uA.createHandles(serv)
	return uA
}

func (s *UserApi) createHandles(ser *ApiServer) {
	//ser.HandleFunc("/user", createHandleFunc(s.api_GetAllUsers)).Methods("GET")
	//ser.HandleFunc("/user", createHandleFunc(s.api_CreateUser)).Methods("POST")
	//ser.HandleFunc("/user/{id}", createHandleFunc(s.api_GetUserByID)).Methods("GET")
	//ser.HandleFunc("/user/{id}", createHandleFunc(s.api_UpdateUser)).Methods("PUT")
	//ser.HandleFunc("/user/{id}", createHandleFunc(s.api_DeleteUser)).Methods("DELETE")
}

func (s *UserApi) api_GetAllUsers(w http.ResponseWriter, r *http.Request) error {
	usrs, err := s.Store.GetAllUser()
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, usrs)
}

func (s *UserApi) api_CreateUser(w http.ResponseWriter, r *http.Request) error {
	var usrReq CreateUserReq
	json.NewDecoder(r.Body).Decode(&usrReq)
	// TODO check if usrReq.email is infact an email
	// TODO Santisation of all the values
	user := user.CreateUser(usrReq.FirstName, usrReq.LastName, usrReq.Email)
	if err := s.Store.CreateUser(user); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, user)
}
func (s *UserApi) api_GetUserByID(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	usr, err := s.Store.GetUserByID(id)
	if err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, usr)
}
func (s *UserApi) api_UpdateUser(w http.ResponseWriter, r *http.Request) error {
	// TODO Implement
	return fmt.Errorf("not implemented")
}

func (s *UserApi) api_DeleteUser(w http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	if err := s.Store.DeleteUser(id); err != nil {
		return err
	}
	return writeJSON(w, http.StatusOK, fmt.Sprintf(`user.User with ID: %s was deleted`, id))
}
