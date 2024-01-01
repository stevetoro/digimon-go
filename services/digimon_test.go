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

func TestDigimonID(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("agumon.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	digimon, err := dc.Digimon.ID(289)
	if err != nil {
		panic(err)
	}

	assertAgumon(t, digimon)
}

func TestDigimonName(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("agumon.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	digimon, err := dc.Digimon.Name("agumon")
	if err != nil {
		panic(err)
	}

	assertAgumon(t, digimon)
}

func TestDigimonList(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Digimon.List()
	if err != nil {
		panic(err)
	}

	assert.Equal(t, 5, len(page.Content))
	assert.Equal(t, 1, page.Content[0].ID)
	assert.Equal(t, "Garummon", page.Content[0].Name)
	assert.Equal(t, "https://digi-api.com/api/v1/digimon/1", page.Content[0].Href)
	assert.Equal(t, "https://digi-api.com/images/digimon/w/Garummon.png", page.Content[0].Image)
	assert.Equal(t, 0, page.Pageable.CurrentPage)
	assert.Equal(t, 5, page.Pageable.ElementsOnPage)
	assert.Equal(t, 1422, page.Pageable.TotalElements)
	assert.Equal(t, 283, page.Pageable.TotalPages)
	assert.Equal(t, "", page.Pageable.PreviousPage)
	assert.Equal(t, "https://digi-api.com/api/v1/digimon?page=1", page.Pageable.NextPage)
}

func TestDigimonListQueryParams(t *testing.T) {
	query := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		query = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-page.json"),
		}
	})

	qp := resources.DigimonQueryParams{
		Name:      "test-name",
		Attribute: "test-attr",
		XAntibody: "test-x",
		Level:     "test-level",
		Page:      5,
		PageSize:  10,
	}

	digimon.NewDigimonClientWith(tc).Digimon.WithQueryParams(qp).List()
	assert.Equal(t, "attribute=test-attr&level=test-level&name=test-name&page=5&pageSize=10&xAntibody=test-x", query)
}

func TestDigimonPageNext(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Digimon.List()
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

func TestDigimonNoNextPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-last-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Digimon.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Next()
	assert.Error(t, services.ErrNoNextPage, err)
}

func TestDigimonPagePrev(t *testing.T) {
	path := ""
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		path = req.URL.Query().Encode()
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-next-page.json"),
		}
	})

	dc := digimon.NewDigimonClientWith(tc)
	page, err := dc.Digimon.List()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "", path)

	_, err = page.Prev()
	if err != nil {
		panic(err)
	}
	assert.Equal(t, "page=0", path)
}

func TestDigimonNoPrevPageErr(t *testing.T) {
	tc := test.NewHTTPClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 200,
			Header:     make(http.Header),
			Body:       test.LoadTestData("digimon-page.json"),
		}
	})

	page, err := digimon.NewDigimonClientWith(tc).Digimon.List()
	if err != nil {
		panic(err)
	}

	_, err = page.Prev()
	assert.Error(t, services.ErrNoPrevPage, err)
}

func assertAgumon(t *testing.T, digimon services.Digimon) {
	assert.Equal(t, 289, digimon.ID)
	assert.Equal(t, "Agumon", digimon.Name)
	assert.Equal(t, false, digimon.XAntibody)
	assert.Equal(t, "1997", digimon.ReleaseDate)
	assert.Equal(t, "https://digi-api.com/images/digimon/w/Agumon.png", digimon.Images[0].Href)
	assert.Equal(t, false, digimon.Images[0].Transparent)
	assert.Equal(t, 7, digimon.Levels[0].ID)
	assert.Equal(t, "Child", digimon.Levels[0].Level)
	assert.Equal(t, 45, digimon.Types[0].ID)
	assert.Equal(t, "Reptile", digimon.Types[0].Type)
	assert.Equal(t, 3, digimon.Attributes[0].ID)
	assert.Equal(t, "Vaccine", digimon.Attributes[0].Attribute)
	assert.Equal(t, 6, len(digimon.Fields))
	assert.Equal(t, 4, digimon.Fields[0].ID)
	assert.Equal(t, "Unknown", digimon.Fields[0].Field)
	assert.Equal(t, "https://digi-api.com/images/etc/fields/Unknown.png", digimon.Fields[0].Image)
	assert.Equal(t, 2, len(digimon.Descriptions))
	assert.Equal(t, "reference_book", digimon.Descriptions[0].Origin)
	assert.Equal(t, "en_us", digimon.Descriptions[0].Language)
	assert.Equal(t, "A Reptile Digimon with an appearance resembling a small dinosaur, it has grown and become able to walk on two legs. Its strength is weak as it is still in the process of growing, but it has a fearless and rather ferocious personality. Hard, sharp claws grow from both its hands and feet, and their power is displayed in battle. It also foreshadows an evolution into a great and powerful Digimon. Its Special Move is spitting a fiery breath from its mouth to attack the opponent (Baby Flame). ", digimon.Descriptions[0].Description)
	assert.Equal(t, 18, len(digimon.Skills))
	assert.Equal(t, 597, digimon.Skills[0].ID)
	assert.Equal(t, "Surudoi Tsume", digimon.Skills[0].Skill)
	assert.Equal(t, "Sharp Claw", digimon.Skills[0].Translation)
	assert.Equal(t, "Uses its claws to rend the opponent. ", digimon.Skills[0].Description)
	assert.Equal(t, 8, len(digimon.PriorEvolutions))
	assert.Equal(t, 532, digimon.PriorEvolutions[0].ID)
	assert.Equal(t, "Wanyamon", digimon.PriorEvolutions[0].Digimon)
	assert.Equal(t, "", digimon.PriorEvolutions[0].Condition)
	assert.Equal(t, "https://digi-api.com/images/digimon/w/Wanyamon.png", digimon.PriorEvolutions[0].Image)
	assert.Equal(t, "https://digi-api.com/api/v1/digimon/532", digimon.PriorEvolutions[0].URL)
	assert.Equal(t, 72, len(digimon.NextEvolutions))
	assert.Equal(t, 706, digimon.NextEvolutions[0].ID)
	assert.Equal(t, "Metal Greymon", digimon.NextEvolutions[0].Digimon)
	assert.Equal(t, "", digimon.NextEvolutions[0].Condition)
	assert.Equal(t, "https://digi-api.com/images/digimon/w/Metal_Greymon.png", digimon.NextEvolutions[0].Image)
	assert.Equal(t, "https://digi-api.com/api/v1/digimon/706", digimon.NextEvolutions[0].URL)
}
