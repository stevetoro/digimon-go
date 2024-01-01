package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stevetoro/digimon-go/resources"
)

type DigimonService struct {
	path   string
	client *http.Client
	params *resources.DigimonQueryParams
}

type Digimon struct {
	resources.Digimon
}

type DigimonPage struct {
	resources.DigimonPage

	client *http.Client
}

func NewDigimonService(client *http.Client) DigimonService {
	return DigimonService{
		path:   "/digimon",
		client: client,
	}
}

func (d DigimonService) Name(s string) (res Digimon, err error) {
	err = do(d.client, fmt.Sprintf("%s/%s", d.Endpoint(), s), &res)
	return res, err
}

func (d DigimonService) ID(i int) (res Digimon, err error) {
	err = do(d.client, fmt.Sprintf("%s/%d", d.Endpoint(), i), &res)
	return res, err
}

func (d DigimonService) WithQueryParams(q resources.DigimonQueryParams) DigimonService {
	d.params = &q
	return d
}

func (d DigimonService) List() (res DigimonPage, err error) {
	res.client = d.client
	err = do(d.client, d.EndpointWithParams(), &res)
	return res, err
}

func (d DigimonService) Endpoint() string {
	return dapiUrl + d.path
}

func (d DigimonService) EndpointWithParams() string {
	e := d.Endpoint()
	if d.params != nil {
		e += "?" + encodeQueryParams(d.params)
	}
	return e
}

func (d DigimonPage) Next() (res DigimonPage, err error) {
	res.client = d.client
	if d.Pageable.NextPage == "" {
		return d, ErrNoNextPage
	}
	e := strings.Split(d.Pageable.NextPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}

func (d DigimonPage) Prev() (res DigimonPage, err error) {
	res.client = d.client
	if d.Pageable.PreviousPage == "" {
		return d, ErrNoPrevPage
	}
	e := strings.Split(d.Pageable.PreviousPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}
