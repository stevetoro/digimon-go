package digimon

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stevetoro/digimon-go/resources"
)

type AttributeService struct {
	path   string
	client *http.Client
	params *resources.QueryParams
}

type Attribute struct {
	resources.Attribute
}

type AttributePage struct {
	resources.AttributePage

	client  *http.Client
	baseURL string
}

func NewAttributeService(client *http.Client) AttributeService {
	return AttributeService{
		path:   "/attribute",
		client: client,
	}
}

func (d AttributeService) Name(s string) (res Attribute, err error) {
	err = do(d.client, fmt.Sprintf("%s/%s", d.Endpoint(), s), &res)
	return res, err
}

func (d AttributeService) ID(i int) (res Attribute, err error) {
	err = do(d.client, fmt.Sprintf("%s/%d", d.Endpoint(), i), &res)
	return res, err
}

func (d AttributeService) WithQueryParams(q resources.QueryParams) AttributeService {
	d.params = &q
	return d
}

func (d AttributeService) List() (res AttributePage, err error) {
	res.client = d.client
	err = do(d.client, d.EndpointWithParams(), &res)
	return res, err
}

func (d AttributeService) Endpoint() string {
	return dapiUrl + d.path
}

func (d AttributeService) EndpointWithParams() string {
	e := d.Endpoint()
	if d.params != nil {
		e += "?" + encodeQueryParams(d.params)
	}
	return e
}

func (d AttributePage) Next() (res AttributePage, err error) {
	res.client = d.client
	if d.Pageable.NextPage == "" {
		return d, NoNextPageErr
	}
	e := strings.Split(d.Pageable.NextPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}

func (d AttributePage) Prev() (res AttributePage, err error) {
	res.client = d.client
	if d.Pageable.PreviousPage == "" {
		return d, NoPrevPageErr
	}
	e := strings.Split(d.Pageable.PreviousPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}
