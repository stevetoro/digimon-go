package digimon_test

import (
	"net/http"
	"testing"

	"github.com/stevetoro/digimon-go"
	"github.com/stevetoro/digimon-go/resources"
	"github.com/stevetoro/digimon-go/test"
	"github.com/stretchr/testify/assert"
)

func TestTypeID(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("cyborg.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	typ, err := dc.Type.ID(1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, typ.ID)
	assert.Equal(t, "Cyborg", typ.Name)
}

func TestTypeName(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("cyborg.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	typ, err := dc.Type.Name("cyborg")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, typ.ID)
	assert.Equal(t, "Cyborg", typ.Name)
}

func TestTypeList(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("type-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Type.List()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Type", page.Content.Name)
	expectedDesc := "A Digimon's Type indicates what sort of category a Digimon's specific species belongs to. Many of these simply indicate what a Digimon is based on and only come into play under certain situations - some Digimon may have a certain advantage or disadvantage to a Digimon of another type. Or an item will work on a Digimon of one type or not the other."
	assert.Equal(t, expectedDesc, page.Content.Description)
	assert.Equal(t, 5, len(page.Content.Fields))
	assert.Equal(t, 1, page.Content.Fields[0].ID)
	assert.Equal(t, "Cyborg", page.Content.Fields[0].Name)
	assert.Equal(t, "https://digi-api.com/api/v1/type/1", page.Content.Fields[0].Href)
}

func TestTypeListQueryParams(t *testing.T) {
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

	digimon.NewDigimonClientWith(tc).Type.WithQueryParams(qp).List()
	assert.Equal(t, "name=test-name&page=5", query)
}

func TestTypePageNext(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("type-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Type.List()
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

func TestTypeNoNextPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("type-last-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Type.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Next()
	assert.Error(t, digimon.NoNextPageErr, err)
}

func TestTypePagePrev(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("type-last-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Type.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	page, err = page.Prev()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=28", path)
}

func TestTypeNoPrevPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("type-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Type.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Prev()
	assert.Error(t, digimon.NoPrevPageErr, err)
}
