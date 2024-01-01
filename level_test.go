package digimon_test

import (
	"net/http"
	"testing"

	"github.com/stevetoro/digimon-go"
	"github.com/stevetoro/digimon-go/resources"
	"github.com/stevetoro/digimon-go/test"
	"github.com/stretchr/testify/assert"
)

func TestLevelID(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("hybrid.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	attribute, err := dc.Level.ID(1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, attribute.ID)
	assert.Equal(t, "Hybrid", attribute.Name)
	expectedDesc := "The Evolution Stage designation of Digimon created using a Legendary Spirit. Hybrids are divided into five sub-categories called Forms, whose respective power levels are loosely equivalent to particular natural Evolution Stages."
	assert.Equal(t, expectedDesc, attribute.Description)
}

func TestLevelName(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("hybrid.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	attribute, err := dc.Level.Name("hybrid")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, attribute.ID)
	assert.Equal(t, "Hybrid", attribute.Name)
	expectedDesc := "The Evolution Stage designation of Digimon created using a Legendary Spirit. Hybrids are divided into five sub-categories called Forms, whose respective power levels are loosely equivalent to particular natural Evolution Stages."
	assert.Equal(t, expectedDesc, attribute.Description)
}

func TestLevelList(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Level.List()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Level", page.Content.Name)
	expectedDesc := "An Evolution Stage, also referred to as Level and Generation, is the level of development a Digimon is at."
	assert.Equal(t, expectedDesc, page.Content.Description)
	assert.Equal(t, 5, len(page.Content.Fields))
	assert.Equal(t, 1, page.Content.Fields[0].ID)
	assert.Equal(t, "Hybrid", page.Content.Fields[0].Name)
	assert.Equal(t, "https://digi-api.com/api/v1/level/1", page.Content.Fields[0].Href)
}

func TestLevelListQueryParams(t *testing.T) {
	query := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		query = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-page.json"),
		}
	})

	qp := resources.QueryParams{
		Name: "test-name",
		Page: 5,
	}

	digimon.NewDigimonClientWith(tc).Level.WithQueryParams(qp).List()
	assert.Equal(t, "name=test-name&page=5", query)
}

func TestLevelPageNext(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Level.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	page, err = page.Next()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=1", path)
}

func TestLevelNoNextPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-last-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Level.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Next()
	assert.Error(t, digimon.NoNextPageErr, err)
}

func TestLevelPagePrev(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-last-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Level.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	page, err = page.Prev()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=0", path)
}

func TestLevelNoPrevPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("level-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Level.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Prev()
	assert.Error(t, digimon.NoPrevPageErr, err)
}
