package modelA

type Pagination struct {
	Page     uint         `json:"page,omitempty" copier:"Page"`
	Limit    uint         `json:"limit,omitempty" copier:"Limit"`
	OrderBy  string       `json:"orderBy,omitempty" copier:"OrderBy"`
	Asc      bool         `json:"asc,omitempty" copier:"Asc"`
	Keyword  []string     `json:"keyword,omitempty" copier:"Keyword"`
	Language LanguageType `json:"language"  copier:"Language"`
}

func (this Pagination) Init() Pagination {
	this.Page = 1
	this.Limit = 10
	this.OrderBy = "updated_at"
	this.Asc = false
	this.Language = TH

	return this
}
