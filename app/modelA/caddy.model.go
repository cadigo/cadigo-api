package modelA

import (
	"cadigo-api/graph/modelgraph"

	"github.com/jinzhu/copier"
)

type Caddy struct {
	ID           string   `json:"id,omitempty" copier:"Id"`
	Name         string   `json:"name,omitempty" copier:"Name"`
	Location     string   `json:"location,omitempty" copier:"Location"`
	Avialability string   `json:"avialability,omitempty" copier:"Avialability"`
	Skill        []string `json:"skill,omitempty" copier:"Skill"`
	Start        int      `json:"start,omitempty" copier:"Start"`
	Description  string   `json:"description,omitempty" copier:"Description,nopanic"`
	Time         []string `json:"time,omitempty" copier:"Time"`
	Cost         float64  `json:"cost,omitempty" copier:"Cost"`
}

func (this Caddy) Parse(graph modelgraph.CaddyInput) error {
	return copier.Copy(&this, &graph)
}

func (this Caddy) ToGraph() modelgraph.Caddy {
	return modelgraph.Caddy{
		Cost: &this.Cost,
	}
}
