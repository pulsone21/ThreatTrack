package tests

import (
	"ContentManagement/api"
	"ContentManagement/api/models/ioc"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (t *SuiteTest) TestIocHandler() {
	fmt.Println("Starting incident handler testing")
	baseUrl := fmt.Sprintf("%s/incident", t.serverAdress)

	// Get All Incident
	err := getAllIocs(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Creating a Incident
	iocID, err := createIoc(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Get Specific Incident
	if err := getIocByID(baseUrl, iocID); err != nil {
		t.Error(err)
	}

	// Delete Incident
	if err := deleteIoc(baseUrl, iocID); err != nil {
		t.Error(err)
	}

}

func (t *SuiteTest) TestIocTypeHandler() {
	// Get All Incident Types
	// Creating a Incident Type
	// Get specficic Type
	// Delete a Incident Type
	fmt.Println("Starting incidentType handler testing")
	baseUrl := fmt.Sprintf("%s/incidenttype", t.serverAdress)

	// Get All Incident
	err := getAllIocTypes(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Creating a Incident
	iocID, err := createIocType(baseUrl)
	if err != nil {
		t.Error(err)
	}

	// Get Specific Incident
	if err := getIocTypeByID(baseUrl, iocID); err != nil {
		t.Error(err)
	}

	// Delete Incident
	if err := deleteIocType(baseUrl, iocID); err != nil {
		t.Error(err)
	}
}

func getAllIocs(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetAllIocs - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}
	ACTUAL := string(body)
	EXPECTED := "[{\"id\":\"49c9793c-8492-468b-8ae0-64e37eb01fa0\",\"value\":\"youtube.com\",\"iocType\":{\"id\":2,\"name\":\"DOMAIN\"},\"verdict\":\"Neutral\"},{\"id\":\"b7a8ae7e-55ae-4983-bd36-ba26c5320487\",\"value\":\"google.com\",\"iocType\":{\"id\":2,\"name\":\"DOMAIN\"},\"verdict\":\"Neutral\"}]"
	if ACTUAL != EXPECTED {
		return fmt.Errorf("GetAllIocs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL)
	}
	fmt.Println("GetAllIocs - Passed")
	return nil
}

func getAllIocTypes(url string) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetAllIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	ACTUAL := string(body)
	EXPECTED := "[{\"id\":1,\"name\":\"URL\"},{\"id\":2,\"name\":\"DOMAIN\"}]"
	if ACTUAL != EXPECTED {
		return fmt.Errorf("GetAllIocType - FAILED - Expected %v but got %v", EXPECTED, ACTUAL)
	}
	fmt.Println("GetAllIocType - Passed")
	return nil
}

func createIoc(url string) (string, error) {
	inc_marshalled, err := json.Marshal(&api.CreatIocReq{
		Value:       "IncidentTwo",
		IncidentIDs: []string{},
		IOCType:     1,
	})
	if err != nil {
		return "", err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CreateIoc - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var IOC ioc.IOC

	json.NewDecoder(res.Body).Decode(&IOC)
	EXPECTED_STATUS := ioc.Neutral
	if IOC.Verdict != EXPECTED_STATUS {
		return "", fmt.Errorf("CreateIoc - FAILED - expected Ioc Verdict: %s but got %s", EXPECTED_STATUS, IOC.Verdict)
	}
	fmt.Printf("CreateIoc - PASSED - Incident with ID: '%s' was created. \n", IOC.Id.String())
	return IOC.Id.String(), nil
}

func createIocType(url string) (string, error) {
	inc_marshalled, err := json.Marshal(&map[string]any{
		"Name": "TEST_TYPE",
	})
	if err != nil {
		return "", err
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		return "", err
	}

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("CreatIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var iT ioc.IOCType
	json.NewDecoder(res.Body).Decode(&iT)
	fmt.Printf("CreatIocType - PASSED - Incident with ID: '%d' was created. \n", iT.Id)
	return fmt.Sprint(iT.Id), nil
}

func getIocByID(baseUrl, iocID string) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, iocID))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GetIocByID - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var IOC ioc.IOC

	json.NewDecoder(res.Body).Decode(&IOC)

	if IOC.Id.String() != iocID {
		return fmt.Errorf("GetIocByID - FAILED - expected IocType ID: %s but got %s", iocID, IOC.Id.String())
	}
	fmt.Printf("GetIocByID - PASSED - Got Incident with ID: '%s'\n", IOC.Id.String())
	return nil
}

func getIocTypeByID(baseUrl, iocID string) error {
	res, err := http.Get(fmt.Sprintf("%s/%s", baseUrl, iocID))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("getIocTypeByID - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var iT ioc.IOCType

	json.NewDecoder(res.Body).Decode(&iT)

	if fmt.Sprint(iT.Id) != iocID {
		return fmt.Errorf("getIocTypeByID - FAILED - expected IocType ID: %s but got %d", iocID, iT.Id)
	}
	fmt.Printf("getIocTypeByID - PASSED - Got Incident with ID: '%d'\n", iT.Id)
	return nil
}

func deleteIoc(baseUrl, iocID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", baseUrl, iocID), nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("DeleteIoc - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	EXPECTED := fmt.Sprintf(`Incident with ID: %s was deleted`, iocID)
	if ACTUAL["message"] != EXPECTED {
		return fmt.Errorf("DeleteIoc - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["message"])
	}
	fmt.Printf("DeleteIoc - PASSED - '%s'\n", ACTUAL["message"])
	return nil
}

func deleteIocType(baseUrl, iocID string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", baseUrl, iocID), nil)
	if err != nil {
		return err
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("DeleteIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode)
	}
	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	EXPECTED := fmt.Sprintf(`IocType with ID: %s was deleted`, iocID)
	if ACTUAL["message"] != EXPECTED {
		return fmt.Errorf("DeleteIocType - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["message"])
	}
	fmt.Printf("DeleteIocType - PASSED - '%s'\n", ACTUAL["message"])
	return nil
}
