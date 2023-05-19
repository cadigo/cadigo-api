package modelapp

import (
	modelgraph "cadigo-api/graph/model"

	"github.com/jinzhu/copier"
)

type Caddy struct {
	ID           *string   `json:"id,omitempty"`
	Name         *string   `json:"name,omitempty"`
	Location     *string   `json:"location,omitempty"`
	Avialability *string   `json:"avialability,omitempty"`
	Skill        []*string `json:"skill,omitempty"`
	Start        *int      `json:"start,omitempty"`
	Description  *string   `json:"description,omitempty"`
	Time         []*string `json:"time,omitempty"`
	Cost         *float64  `json:"cost,omitempty"`
}

func (this *Caddy) parse(graph modelgraph.Caddy) {
	copier.Copy(this, graph)
}

func (this *Caddy) toGraph() modelgraph.Caddy {
	return modelgraph.Caddy{
		Cost: this.Cost,
	}
}
