package modela

import (
	"github.com/jinzhu/copier"
)

type Users struct {
	ID string `json:"id,omitempty" copier:"Id"`
}

func (this Users) Parse(graph interface{}) error {
	return copier.Copy(&this, &graph)
}

func (this Users) ToGraph() interface{} {
	var g interface{}
	copier.Copy(&g, &this)
	return g
}
