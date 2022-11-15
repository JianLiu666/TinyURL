package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	v1 "tinyurl/internal/api/v1"
	"tinyurl/internal/config"
)

func create_200(s *session) (bool, error) {
	// 1. prepare request body
	reqData := &v1.CreateReqBody{
		Url:   s.origin,
		Alias: s.atlas,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return false, err
	}

	domain := fmt.Sprintf("http://%s%s/api/v1/create",
		config.Env().Server.Domain,
		config.Env().Server.Port,
	)
	req, err := http.NewRequest(http.MethodPost, domain, bytes.NewBuffer(reqBody))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("response not equal. (expected: %v, actual: %v)", http.StatusOK, resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	respData := &v1.CreateRespBody{}
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		return false, err
	}

	if respData.Origin != reqData.Url {
		return false, fmt.Errorf("origin url not equal. (expected: %s, actual: %s)", reqData.Url, respData.Origin)
	}

	s.tiny = respData.Tiny

	return true, nil
}

func create_400(s *session) (bool, error) {
	// 1. prepare request body
	reqData := &v1.CreateReqBody{
		Url:   s.origin,
		Alias: s.atlas,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		return false, err
	}

	domain := fmt.Sprintf("http://%s%s/api/v1/create",
		config.Env().Server.Domain,
		config.Env().Server.Port,
	)
	req, err := http.NewRequest(http.MethodPost, domain, bytes.NewBuffer(reqBody))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusBadRequest {
		return false, fmt.Errorf("response not equal. (expected: %v, actual: %v)", http.StatusOK, resp.StatusCode)
	}

	return true, nil
}
