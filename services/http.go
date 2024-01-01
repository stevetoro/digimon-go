package services

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

const dapiUrl = "https://digi-api.com/api/v1"

func do(client *http.Client, url string, obj any) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &obj)
}

func encodeQueryParams(opt any) string {
	v, _ := query.Values(opt)
	return v.Encode()
}
