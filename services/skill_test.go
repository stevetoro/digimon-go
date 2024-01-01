package services_test

import (
	"net/http"
	"testing"

	"github.com/stevetoro/digimon-go"
	"github.com/stevetoro/digimon-go/resources"
	"github.com/stevetoro/digimon-go/services"
	"github.com/stevetoro/digimon-go/test"
	"github.com/stretchr/testify/assert"
)

func TestSkillID(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("speed-star.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	skill, err := dc.Skill.ID(1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, skill.ID)
	assert.Equal(t, "Speed Star", skill.Name)
	expectedDesc := "Charging at ultra-high speed, it tears the opponent into pieces with the wing blades on its back. "
	assert.Equal(t, expectedDesc, skill.Description)
}

func TestSkillName(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("speed-star.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	skill, err := dc.Skill.Name("speed star")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, skill.ID)
	assert.Equal(t, "Speed Star", skill.Name)
	expectedDesc := "Charging at ultra-high speed, it tears the opponent into pieces with the wing blades on its back. "
	assert.Equal(t, expectedDesc, skill.Description)
}

func TestSkillList(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Skill.List()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Skill", page.Content.Name)
	expectedDesc := "It's often used for offensive means, however it can be used for other objectives such as healing itself or another digimon.\nIt may happen that certain skills that can be unique for a certain Digimon, however this is not a rule,  various Digimon have abilities that's shared between then."
	assert.Equal(t, expectedDesc, page.Content.Description)
	assert.Equal(t, 5, len(page.Content.Fields))
	assert.Equal(t, 1, page.Content.Fields[0].ID)
	assert.Equal(t, "Speed Star", page.Content.Fields[0].Name)
	assert.Equal(t, "https://digi-api.com/api/v1/skill/1", page.Content.Fields[0].Href)
}

func TestSkillListQueryParams(t *testing.T) {
	query := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		query = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-page.json"),
		}
	})

	qp := resources.QueryParams{
		Name: "test-name",
		Page: 5,
	}

	digimon.NewDigimonClientWith(tc).Skill.WithQueryParams(qp).List()
	assert.Equal(t, "name=test-name&page=5", query)
}

func TestSkillPageNext(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Skill.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	_, err = page.Next()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=1", path)
}

func TestSkillNoNextPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-last-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Skill.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Next()
	assert.Error(t, services.ErrNoNextPage, err)
}

func TestSkillPagePrev(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-last-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Level.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	_, err = page.Prev()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=705", path)
}

func TestSkillNoPrevPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("skill-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Skill.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Prev()
	assert.Error(t, services.ErrNoPrevPage, err)
}
