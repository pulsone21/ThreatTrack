package gui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
)

type IndicatorHandler struct {
	backendAdress string
}

type indTableViewData struct {
	Indicators []Indicator
}

type Indicator struct {
	Id      string        `json:"id"`
	Value   string        `json:"value"`
	Type    IndicatorType `json:"iocType"`
	Verdict Verdict       `json:"verdict"`
}

type IndicatorType struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Verdict string

const (
	Benigne   Verdict = "Benigne"
	Malicious Verdict = "Malicious"
	Neutral   Verdict = "Neutral"
)

func (iH *IndicatorHandler) createHandles(s *Server) {
	s.Router.HandleFunc("/indicatorTable/", s.createHandleFunc(iH.serveIncidentTable)).Methods("GET")
	s.Router.HandleFunc("/indicator/", s.createHandleFunc(iH.serveIndicatorPage)).Methods("GET")
}

func CreateIndicatorHandler(ser *Server, backendBase string) *IndicatorHandler {
	iH := &IndicatorHandler{
		backendAdress: fmt.Sprintf("%s/ioc/", backendBase),
	}
	iH.createHandles(ser)
	return iH
}

func (iH *IndicatorHandler) serveIncidentTable(w http.ResponseWriter, r *http.Request) error {
	fmt.Printf("\nrequesting backend with %s \n", iH.backendAdress)
	res, err := http.Get(iH.backendAdress)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("backend request failed: %v", res.StatusCode)
	}
	resbody, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	tmpl, err := template.ParseFiles("./templates/indicatorTable.html")
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	var inds []Indicator

	if err = json.Unmarshal(resbody, &inds); err != nil {
		return err
	}
	fmt.Println(inds)

	return tmpl.Execute(w, indTableViewData{
		Indicators: inds,
	})
}

func (iH *IndicatorHandler) serveIndicatorPage(w http.ResponseWriter, r *http.Request) error {
	incId := r.URL.Query().Get("id")
	url := fmt.Sprintf("%s/%s", iH.backendAdress, incId)
	fmt.Printf("\nrequesting backend with %s \n", url)

	res, err := http.Get(url)
	if err != nil {

		log.Fatalln(err.Error())
		return err
	}

	var ind Indicator
	if err = json.NewDecoder(res.Body).Decode(&ind); err != nil {
		log.Fatalln(err.Error())
		return err
	}
	fmt.Println(ind)
	tmpl, err := template.ParseFiles("./templates/indicator.html")
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return tmpl.Execute(w, ind)
}
