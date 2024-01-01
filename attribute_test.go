package digimon_test

import (
	"net/http"
	"testing"

	"github.com/stevetoro/digimon-go"
	"github.com/stevetoro/digimon-go/resources"
	"github.com/stevetoro/digimon-go/test"
	"github.com/stretchr/testify/assert"
)

func TestAttributeID(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("variable.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	attribute, err := dc.Attribute.ID(1)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, attribute.ID)
	assert.Equal(t, "Variable", attribute.Name)
	expectedDesc := "Digimon with the Variable attribute, which at least so far means Hybrid, change attribute to match that of their opponent. Related to the Digimon known as 'Ancient Species'."
	assert.Equal(t, expectedDesc, attribute.Description)
}

func TestAttributeName(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("variable.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	attribute, err := dc.Attribute.Name("variable")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 1, attribute.ID)
	assert.Equal(t, "Variable", attribute.Name)
	expectedDesc := "Digimon with the Variable attribute, which at least so far means Hybrid, change attribute to match that of their opponent. Related to the Digimon known as 'Ancient Species'."
	assert.Equal(t, expectedDesc, attribute.Description)
}

func TestAttributeList(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Attribute.List()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "Attribute", page.Content.Name)
	expectedDesc := "Attribute refers to the type of computer file a Digimon represents. The Digimon are generally separated into three different attributes 'Data', 'Vaccine' and 'Virus', depending on the influence that they cause in the Environment. Digimon with the same name and different attributes sometimes look different and in some cases have different attacks or types as well. There are also 'Free' and 'Variable' attributes that are related to the Digimon known as 'Ancient Species'. Digimon with the Variable attribute, which at least so far means Hybrid, change attribute to match that of their opponent. Vaccine digimon are supposed to be more powerful against Virus digimon, Virus digimon are supposed to be more powerful against Data digimon, and Data digimon are supposed to be more powerful against Vaccine digimon. Digimon of the same attribute will generally get along well with each other. However, all of the above is not set in stone and there have been heroic Virus Digimon, as well as villainous Vaccine Digimon. The Attribute system started out as a sort of 'Rock, Paper, Scissors' determinator in the virtual pets. This is still, in essence, in use in most mediums (especially the Card Game), though less strongly stressed."
	assert.Equal(t, expectedDesc, page.Content.Description)
	assert.Equal(t, 5, len(page.Content.Fields))
	assert.Equal(t, 1, page.Content.Fields[0].ID)
	assert.Equal(t, "Variable", page.Content.Fields[0].Name)
	assert.Equal(t, "https://digi-api.com/api/v1/attribute/1", page.Content.Fields[0].Href)
}

func TestAttributeListQueryParams(t *testing.T) {
	query := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		query = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-page.json"),
		}
	})

	qp := resources.QueryParams{
		Name: "test-name",
		Page: 5,
	}

	digimon.NewDigimonClientWith(tc).Attribute.WithQueryParams(qp).List()
	assert.Equal(t, "name=test-name&page=5", query)
}

func TestAttributePageNext(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Attribute.List()
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

func TestAttributeNoNextPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-last-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Attribute.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Next()
	assert.Error(t, digimon.NoNextPageErr, err)
}

func TestAttributePagePrev(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-last-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Attribute.List()
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

func TestAttributeNoPrevPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("attribute-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Attribute.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Prev()
	assert.Error(t, digimon.NoPrevPageErr, err)
}
