package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stevetoro/digimon-go/resources"
)

type TypeService struct {
	path   string
	client *http.Client
	params *resources.QueryParams
}

type Type struct {
	resources.Type
}

type TypePage struct {
	resources.TypePage

	client *http.Client
}

func NewTypeService(client *http.Client) TypeService {
	return TypeService{
		path:   "/type",
		client: client,
	}
}

func (d TypeService) Name(s string) (res Type, err error) {
	err = do(d.client, fmt.Sprintf("%s/%s", d.Endpoint(), s), &res)
	return res, err
}

func (d TypeService) ID(i int) (res Type, err error) {
	err = do(d.client, fmt.Sprintf("%s/%d", d.Endpoint(), i), &res)
	return res, err
}

func (d TypeService) WithQueryParams(q resources.QueryParams) TypeService {
	d.params = &q
	return d
}

func (d TypeService) List() (res TypePage, err error) {
	res.client = d.client
	err = do(d.client, d.EndpointWithParams(), &res)
	return res, err
}

func (d TypeService) Endpoint() string {
	return dapiUrl + d.path
}

func (d TypeService) EndpointWithParams() string {
	e := d.Endpoint()
	if d.params != nil {
		e += "?" + encodeQueryParams(d.params)
	}
	return e
}

func (d TypePage) Next() (res TypePage, err error) {
	res.client = d.client
	if d.Pageable.NextPage == "" {
		return d, ErrNoNextPage
	}
	e := strings.Split(d.Pageable.NextPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}

func (d TypePage) Prev() (res TypePage, err error) {
	res.client = d.client
	if d.Pageable.PreviousPage == "" {
		return d, ErrNoPrevPage
	}
	e := strings.Split(d.Pageable.PreviousPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}
