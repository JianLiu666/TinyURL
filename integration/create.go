package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"tinyurl/config"
	v1 "tinyurl/pkg/api/v1"
)

func create_ok(s *session) {
	// 1. prepare request body
	reqData := &v1.CreateReqBody{
		Url:   s.origin,
		Alias: s.atlas,
	}
	reqBody, err := json.Marshal(reqData)
	if err != nil {
		log.Fatal(err)
	}

	domain := fmt.Sprintf("http://%s%s/api/v1/create",
		config.Env().Server.Domain,
		config.Env().Server.Port,
	)
	req, err := http.NewRequest(http.MethodPost, domain, bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	// 2. send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusOK {
		log.Fatal("response status code incorrect.")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	respData := &v1.CreateRespBody{}
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		log.Fatal(err)
	}

	if respData.Origin != reqData.Url {
		log.Fatal("data incorrect.")
	}

	s.tiny = respData.Tiny
}
