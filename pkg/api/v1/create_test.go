package v1

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCreate(t *testing.T) {
	jsonBody, err := json.Marshal(&createReqBody{
		Url:   "https://tinyurl.com/app/",
		Alias: "",
	})
	if err != nil {
		panic(err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	_, err = http.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/create", bodyReader)
	if err != nil {
		panic(err)
	}
}
