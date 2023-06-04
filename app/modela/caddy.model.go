package modela

import (
	"cadigo-api/graph/modelgraph"

	"github.com/jinzhu/copier"
)

type Caddy struct {
	ID            string   `json:"id,omitempty" copier:"Id"`
	Reference     string   `json:"reference,omitempty" copier:"Reference"`
	Name          string   `json:"name,omitempty" copier:"Name"`
	Location      string   `json:"location,omitempty" copier:"Location"`
	Avialability  string   `json:"avialability,omitempty" copier:"Avialability"`
	Skill         []string `json:"skill,omitempty" copier:"Skill"`
	Star          int      `json:"star,omitempty" copier:"Star"`
	Description   string   `json:"description,omitempty" copier:"Description,nopanic"`
	Time          []string `json:"time,omitempty" copier:"Time"`
	Cost          float64  `json:"cost,omitempty" copier:"Cost"`
	CourseGolfIDs []string `json:"courseGolfIDs,omitempty" copier:"CourseGolfIDs"`
}

func (this Caddy) Parse(graph modelgraph.CaddyInput) error {
	return copier.Copy(&this, &graph)
}

func (this Caddy) ToGraph() modelgraph.Caddy {
	g := modelgraph.Caddy{}
	copier.Copy(&g, &this)
	return g
}

type CaddyFilter struct {
	Skill         []string `json:"skill,omitempty" copier:"Skill"`
	CourseGolfIDs []string `json:"courseGolfIDs,omitempty" copier:"CourseGolfIDs"`
	Cost          *float64 `json:"cost,omitempty" copier:"Cost"`
	Star          *int     `json:"star,omitempty" copier:"Star"`
	Ids           []string `json:"ids,omitempty" copier:"Ids"`
}

func (this CaddyFilter) Init() CaddyFilter {
	this.Skill = []string{}
	this.CourseGolfIDs = []string{}
	this.Cost = nil
	this.Star = nil
	this.Ids = []string{}

	return this
}
