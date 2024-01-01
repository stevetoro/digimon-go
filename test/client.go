package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
)

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewHTTPClient(fn RoundTripFunc) *http.Client {
	return &http.Client{Transport: fn}
}

func LoadTestData(fileName string) io.ReadCloser {
	file, err := os.Open(fmt.Sprintf("../test/data/%s", fileName))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	return io.NopCloser(bytes.NewBufferString(string(b)))
}
