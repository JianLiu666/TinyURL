package v1

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestCreate_OK(t *testing.T) {
	// 1. prepare request body
	reqData := &createReqBody{
		Url:   "https://tinyurl.com/app/",
		Alias: "",
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/create", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusOK {
		t.Fatal("response status code incorrect.")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	respData := &createRespBody{}
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		t.Fatal(err)
	}

	if respData.Origin != reqData.Url {
		t.Fatal("data incorrect.")
	}
}

func TestCreate_BadRequest(t *testing.T) {
	// 1. prepare request body
	reqData := &createReqBody{
		Url:   "",
		Alias: "",
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/create", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatal("response status code incorrect.")
	}
}
