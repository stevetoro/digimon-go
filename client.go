package digimon

import (
	"net/http"

	"github.com/stevetoro/digimon-go/services"
)

type DigimonClient struct {
	client *http.Client

	Digimon   services.DigimonService
	Attribute services.AttributeService
	Level     services.LevelService
	Type      services.TypeService
	Skill     services.SkillService
}

func NewDigimonClient() DigimonClient {
	return NewDigimonClientWith(http.DefaultClient)
}

func NewDigimonClientWith(h *http.Client) DigimonClient {
	c := DigimonClient{client: http.DefaultClient}
	if h != nil {
		c.client = h
	}
	c.Digimon = services.NewDigimonService(c.client)
	c.Attribute = services.NewAttributeService(c.client)
	c.Level = services.NewLevelService(c.client)
	c.Type = services.NewTypeService(c.client)
	c.Skill = services.NewSkillService(c.client)
	return c
}
