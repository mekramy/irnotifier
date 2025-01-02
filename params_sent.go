package irnotifier

import (
	"encoding/json"
	"time"
)

func SentParameter() *SentParams {
	return &SentParams{}
}

type SentParams struct {
	page     int
	limit    Limit
	sort     SentSort
	order    Order
	search   string
	pattern  string
	metadata string
	from     *time.Time
	to       *time.Time
}

func (params SentParams) ToJson() ([]byte, error) {
	return json.Marshal(map[string]any{
		"page":     params.page,
		"limit":    params.limit,
		"sort":     params.sort,
		"order":    params.order,
		"search":   params.search,
		"pattern":  params.pattern,
		"metadata": params.metadata,
		"from":     params.from,
		"to":       params.to,
	})
}

func (param *SentParams) Page(page int) *SentParams {
	param.page = page
	return param
}

func (param *SentParams) Limit(limit Limit) *SentParams {
	param.limit = limit
	return param
}

func (param *SentParams) Sort(sort SentSort) *SentParams {
	param.sort = sort
	return param
}

func (param *SentParams) Order(order Order) *SentParams {
	param.order = order
	return param
}

func (param *SentParams) Search(search string) *SentParams {
	param.search = search
	return param
}

func (param *SentParams) Pattern(pattern string) *SentParams {
	param.pattern = pattern
	return param
}

func (param *SentParams) Metadata(metadata string) *SentParams {
	param.metadata = metadata
	return param
}

func (param *SentParams) From(from time.Time) *SentParams {
	param.from = &from
	return param
}

func (param *SentParams) To(to time.Time) *SentParams {
	param.to = &to
	return param
}
