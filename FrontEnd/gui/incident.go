package gui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

type IncidentHandler struct {
	backendAdress string
	templatePath  string
}

type incTableViewData struct {
	Incidents []Incident
}

type Incident struct {
	Id           string           `json:"id"`
	Name         string           `json:"name"`
	Severity     IncidentSeverity `json:"severity"`
	IncidentType IncidentType     `json:"type"`
}

type IncidentType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type IncidentSeverity string

const (
	Low      IncidentSeverity = "Low"
	Medium   IncidentSeverity = "Medium"
	High     IncidentSeverity = "High"
	Critical IncidentSeverity = "Critical"
)

func (iH *IncidentHandler) createHandles(s *Server) {
	s.Router.HandleFunc("/incidentTable/", s.createHandleFunc(iH.serveIncidentTable)).Methods("GET")
	s.Router.HandleFunc("/incident/summary", s.createHandleFunc(iH.serveIncidentPage)).Methods("GET")
	s.Router.HandleFunc("/incident/worklog", s.createHandleFunc(iH.serveIncidentWorklog)).Methods("GET")
	s.Router.HandleFunc("/incident/planing", s.createHandleFunc(iH.serveIncidentPlaning)).Methods("GET")
	s.Router.HandleFunc("/incident/iocView", s.createHandleFunc(iH.serveIncidentiocView)).Methods("GET")
}

func CreateIncidentHandler(ser *Server, backendBase string) *IncidentHandler {
	wd, _ := os.Getwd()
	iH := &IncidentHandler{
		backendAdress: fmt.Sprintf("%s/incident", backendBase),
		templatePath:  "templates/incident",
	}
	fmt.Printf("%s/%s\n", wd, iH.templatePath)
	iH.createHandles(ser)
	return iH
}

func (iH *IncidentHandler) serveIncidentTable(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("\nrequesting backend with %s \n", iH.backendAdress)
	res, err := http.Get(iH.backendAdress)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		// Handle the Error
		return http.ErrAbortHandler
	}
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentTable.html", iH.templatePath))
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	var incs []Incident

	if err = json.Unmarshal(resbody, &incs); err != nil {
		return err
	}
	return tmpl.Execute(w, incTableViewData{
		Incidents: incs,
	})
}

func (iH *IncidentHandler) serveIncidentPage(w http.ResponseWriter, r *http.Request) error {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incident.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentWorklog(w http.ResponseWriter, r *http.Request) error {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentWorklogs.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentPlaning(w http.ResponseWriter, r *http.Request) error {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentPlaning.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}

func (iH *IncidentHandler) serveIncidentiocView(w http.ResponseWriter, r *http.Request) error {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	var inc Incident
	if err = json.NewDecoder(res.Body).Decode(&inc); err != nil {
		return err
	}

	tmpl, err := template.ParseFiles(fmt.Sprintf("%s/incidentIOCView.html", iH.templatePath))
	if err != nil {
		return err
	}
	return tmpl.Execute(w, inc)
}
