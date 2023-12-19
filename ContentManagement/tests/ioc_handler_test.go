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

func (t *SuiteTest) Test1GetAllIocs() {
	url := fmt.Sprintf("%s/ioc", t.ApiServer)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}

	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("GetAllIocs - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}
	ACTUAL := string(body)
	EXPECTED := "[{\"id\":\"49c9793c-8492-468b-8ae0-64e37eb01fa0\",\"value\":\"youtube.com\",\"iocType\":{\"id\":2,\"name\":\"DOMAIN\"},\"verdict\":\"Neutral\"},{\"id\":\"b7a8ae7e-55ae-4983-bd36-ba26c5320487\",\"value\":\"google.com\",\"iocType\":{\"id\":2,\"name\":\"DOMAIN\"},\"verdict\":\"Neutral\"}]"
	t.Assertions.Equal(EXPECTED, ACTUAL, fmt.Sprintf("GetAllIocs - FAILED - Expected %v but got %v", EXPECTED, ACTUAL))
}

func (t *SuiteTest) Test1getAllIocTypes() {
	url := fmt.Sprintf("%s/ioctype", t.ApiServer)
	res, err := http.Get(url)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("GetAllIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
		return
	}

	ACTUAL := string(body)
	EXPECTED := "[{\"id\":1,\"name\":\"URL\"},{\"id\":2,\"name\":\"DOMAIN\"}]"
	t.Assertions.Equal(EXPECTED, ACTUAL, fmt.Sprintf("GetAllIocType - FAILED - Expected %v but got %v", EXPECTED, ACTUAL))
}

func (t *SuiteTest) Test2CreateIoc() {
	inc_marshalled, err := json.Marshal(&api.CreatIocReq{
		Value:       "Test.com",
		IncidentIDs: []string{},
		IOCType:     1,
	})
	if err != nil {
		t.Error(err)
		return
	}
	url := fmt.Sprintf("%s/ioc", t.ApiServer)
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("CreateIoc - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	var IOC ioc.IOC

	json.NewDecoder(res.Body).Decode(&IOC)
	EXPECTED_STATUS := ioc.Neutral
	t.Assertions.Equal(EXPECTED_STATUS, IOC.Verdict, fmt.Sprintf("CreateIoc - FAILED - expected Ioc Verdict: %s but got %s", EXPECTED_STATUS, IOC.Verdict))
}

func (t *SuiteTest) Test2CreateIocType() {
	EXPECTED := "TEST_TYPE"
	inc_marshalled, err := json.Marshal(&map[string]any{
		"Name": EXPECTED,
	})
	if err != nil {
		t.Error(err)
		return
	}
	url := fmt.Sprintf("%s/ioctype", t.ApiServer)
	res, err := http.Post(url, "application/json", bytes.NewReader(inc_marshalled))
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("CreatIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	defer res.Body.Close()
	var iT ioc.IOCType
	json.NewDecoder(res.Body).Decode(&iT)

	t.Assertions.Equal(EXPECTED, iT.Name, fmt.Sprintf("CreatIocType - FAILED - expected name was: '%s' but got %s\n", EXPECTED, iT.Name))
}

func (t *SuiteTest) Test3GetIocByID() {
	url := fmt.Sprintf("%s/ioc", t.ApiServer)
	iocID := "49c9793c-8492-468b-8ae0-64e37eb01fa0"
	res, err := http.Get(fmt.Sprintf("%s/%s", url, iocID))
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("GetIocByID - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	defer res.Body.Close()
	var IOC ioc.IOC
	json.NewDecoder(res.Body).Decode(&IOC)
	t.Assertions.Equal(iocID, fmt.Sprint(IOC.Id), fmt.Sprintf("GetIocByID - FAILED - expected IocType ID: %s but got %s", iocID, fmt.Sprint(IOC.Id)))
}

func (t *SuiteTest) Test3GetIocTypeByID() {
	url := fmt.Sprintf("%s/ioctype", t.ApiServer)
	iocID := "49c9793c-8492-468b-8ae0-64e37eb01fa0"
	res, err := http.Get(fmt.Sprintf("%s/%s", url, iocID))
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("getIocTypeByID - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	defer res.Body.Close()
	var iT ioc.IOCType

	json.NewDecoder(res.Body).Decode(&iT)
	t.Assertions.Equal(iocID, fmt.Sprint(iT.Id), fmt.Sprintf("getIocTypeByID - FAILED - expected IocType ID: %s but got %d", iocID, iT.Id))
}

func (t *SuiteTest) Test4DeleteIoc() {
	url := fmt.Sprintf("%s/ioc", t.ApiServer)
	iocID := "49c9793c-8492-468b-8ae0-64e37eb01fa0"
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", url, iocID), nil)
	if err != nil {
		t.Error(err)
		return
	}

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("DeleteIoc - FAILED - expected status code 200 but got %d\n", res.StatusCode))
	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	fmt.Print(ACTUAL)
	EXPECTED := fmt.Sprintf(`Incident with ID: %s was deleted`, iocID)
	t.Assertions.Equal(EXPECTED, ACTUAL["Message"], fmt.Sprintf("DeleteIoc - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["Message"]))
}

func (t *SuiteTest) Test4deleteIocType() {
	url := fmt.Sprintf("%s/ioctype", t.ApiServer)
	iocID := "1"
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", url, iocID), nil)
	if err != nil {
		t.Error(err)
		return
	}
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		t.Error(err)
		return
	}
	t.Assertions.Equal(http.StatusOK, res.StatusCode, fmt.Sprintf("DeleteIocType - FAILED - expected status code 200 but got %d\n", res.StatusCode))

	defer res.Body.Close()
	var ACTUAL map[string]any
	json.NewDecoder(res.Body).Decode(&ACTUAL)
	EXPECTED := fmt.Sprintf(`IocType with ID: %s was deleted`, iocID)
	t.Assertions.Equal(EXPECTED, ACTUAL["Message"], fmt.Sprintf("DeleteIocType - FAILED - expected DeleteMessage: %s but got %s", EXPECTED, ACTUAL["Message"]))
}
