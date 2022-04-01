package internal

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

func DoRequest(method string, endpoint string, body []byte, token string) ([]byte, error) {
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(body))
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")

	resp, err := do(request)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func do(req *http.Request) ([]byte, error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
