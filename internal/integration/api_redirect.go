package integration

import (
	"fmt"
	"net/http"
)

func (t *tester) redirect_302(s *session) (bool, error) {
	// 1. prepare request data
	domain := fmt.Sprintf("http://%s", s.tiny)
	req, err := http.NewRequest(http.MethodGet, domain, nil)
	if err != nil {
		return false, err
	}

	// 2. send request
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusFound {
		return false, fmt.Errorf("response not equal. (expected: %v, actual: %v)", http.StatusFound, resp.StatusCode)
	}

	location, err := resp.Location()
	if err != nil {
		return false, err
	}

	if location.String() != s.origin {
		return false, fmt.Errorf("url not equal. (expected: %s, actual: %s)", s.origin, location.String())
	}

	return true, nil
}

func (t *tester) redirect_400(s *session) (bool, error) {
	// 1. prepare request data
	domain := fmt.Sprintf("http://%s", s.tiny)
	req, err := http.NewRequest(http.MethodGet, domain, nil)
	if err != nil {
		return false, err
	}

	// 2. send request
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	// 3. valid response
	if resp.StatusCode != http.StatusBadRequest {
		return false, fmt.Errorf("response not equal. (expected: %v, actual: %v)", http.StatusBadRequest, resp.StatusCode)
	}

	return true, nil
}
