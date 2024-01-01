package digimon

import (
	"net/http"
)

const dapiUrl = "https://digi-api.com/api/v1"

type DigimonClient struct {
	client *http.Client

	Digimon DigimonService
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
	return c
}
