package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stevetoro/digimon-go/resources"
)

type LevelService struct {
	path   string
	client *http.Client
	params *resources.QueryParams
}

type Level struct {
	resources.Level
}

type LevelPage struct {
	resources.LevelPage

	client *http.Client
}

func NewLevelService(client *http.Client) LevelService {
	return LevelService{
		path:   "/level",
		client: client,
	}
}

func (d LevelService) Name(s string) (res Level, err error) {
	err = do(d.client, fmt.Sprintf("%s/%s", d.Endpoint(), s), &res)
	return res, err
}

func (d LevelService) ID(i int) (res Level, err error) {
	err = do(d.client, fmt.Sprintf("%s/%d", d.Endpoint(), i), &res)
	return res, err
}

func (d LevelService) WithQueryParams(q resources.QueryParams) LevelService {
	d.params = &q
	return d
}

func (d LevelService) List() (res LevelPage, err error) {
	res.client = d.client
	err = do(d.client, d.EndpointWithParams(), &res)
	return res, err
}

func (d LevelService) Endpoint() string {
	return dapiUrl + d.path
}

func (d LevelService) EndpointWithParams() string {
	e := d.Endpoint()
	if d.params != nil {
		e += "?" + encodeQueryParams(d.params)
	}
	return e
}

func (d LevelPage) Next() (res LevelPage, err error) {
	res.client = d.client
	if d.Pageable.NextPage == "" {
		return d, ErrNoNextPage
	}
	e := strings.Split(d.Pageable.NextPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}

func (d LevelPage) Prev() (res LevelPage, err error) {
	res.client = d.client
	if d.Pageable.PreviousPage == "" {
		return d, ErrNoPrevPage
	}
	e := strings.Split(d.Pageable.PreviousPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}
