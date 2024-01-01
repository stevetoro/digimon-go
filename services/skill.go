package services

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/stevetoro/digimon-go/resources"
)

type SkillService struct {
	path   string
	client *http.Client
	params *resources.QueryParams
}

type Skill struct {
	resources.Skill
}

type SkillPage struct {
	resources.SkillPage

	client *http.Client
}

func NewSkillService(client *http.Client) SkillService {
	return SkillService{
		path:   "/level",
		client: client,
	}
}

func (d SkillService) Name(s string) (res Skill, err error) {
	err = do(d.client, fmt.Sprintf("%s/%s", d.Endpoint(), s), &res)
	return res, err
}

func (d SkillService) ID(i int) (res Skill, err error) {
	err = do(d.client, fmt.Sprintf("%s/%d", d.Endpoint(), i), &res)
	return res, err
}

func (d SkillService) WithQueryParams(q resources.QueryParams) SkillService {
	d.params = &q
	return d
}

func (d SkillService) List() (res SkillPage, err error) {
	res.client = d.client
	err = do(d.client, d.EndpointWithParams(), &res)
	return res, err
}

func (d SkillService) Endpoint() string {
	return dapiUrl + d.path
}

func (d SkillService) EndpointWithParams() string {
	e := d.Endpoint()
	if d.params != nil {
		e += "?" + encodeQueryParams(d.params)
	}
	return e
}

func (d SkillPage) Next() (res SkillPage, err error) {
	res.client = d.client
	if d.Pageable.NextPage == "" {
		return d, ErrNoNextPage
	}
	e := strings.Split(d.Pageable.NextPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}

func (d SkillPage) Prev() (res SkillPage, err error) {
	res.client = d.client
	if d.Pageable.PreviousPage == "" {
		return d, ErrNoPrevPage
	}
	e := strings.Split(d.Pageable.PreviousPage, dapiUrl)
	err = do(d.client, e[1], &res)
	return res, err
}
