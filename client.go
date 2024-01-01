package digimon

import (
	"net/http"
)

const dapiUrl = "https://digi-api.com/api/v1"

type DigimonClient struct {
	client *http.Client

	Digimon   DigimonService
	Attribute AttributeService
	Level     LevelService
	Type      TypeService
}

func NewDigimonClient() DigimonClient {
	return NewDigimonClientWith(http.DefaultClient)
}

func NewDigimonClientWith(h *http.Client) DigimonClient {
	c := DigimonClient{client: http.DefaultClient}
	if h != nil {
		c.client = h
	}
	c.Digimon = NewDigimonService(c.client)
	c.Attribute = NewAttributeService(c.client)
	c.Level = NewLevelService(c.client)
	c.Type = NewTypeService(c.client)
	return c
}
