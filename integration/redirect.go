package integration

import (
	"fmt"
	"log"
	"net/http"
)

func redirect_ok(s *session) {
	// 1. prepare request data
	domain := fmt.Sprintf("http://%s", s.tiny)
	req, err := http.NewRequest(http.MethodGet, domain, nil)
	if err != nil {
		log.Fatal(err)
	}

	// 2. send request
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusFound {
		log.Fatal("response status code incorrect.")
	}

	location, err := resp.Location()
	if err != nil {
		log.Fatal(err)
	}

	if location.String() != s.origin {
		log.Fatalf("url mismath: %s -> %s", s.origin, location.String())
	}
}
