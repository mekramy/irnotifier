package irnotifier

import (
	"encoding/json"
	"time"
)

func FailParameter() *FailParams {
	return &FailParams{}
}

type FailParams struct {
	page     int
	limit    Limit
	sort     FailSort
	order    Order
	search   string
	state    FailStatus
	pattern  string
	metadata string
	from     *time.Time
	to       *time.Time
}

func (params FailParams) ToJson() ([]byte, error) {
	return json.Marshal(map[string]any{
		"page":     params.page,
		"limit":    params.limit,
		"sort":     params.sort,
		"order":    params.order,
		"search":   params.search,
		"state":    params.state,
		"pattern":  params.pattern,
		"metadata": params.metadata,
		"from":     params.from,
		"to":       params.to,
	})
}

func (param *FailParams) Page(page int) *FailParams {
	param.page = page
	return param
}

func (param *FailParams) Limit(limit Limit) *FailParams {
	param.limit = limit
	return param
}

func (param *FailParams) Sort(sort FailSort) *FailParams {
	param.sort = sort
	return param
}

func (param *FailParams) Order(order Order) *FailParams {
	param.order = order
	return param
}

func (param *FailParams) Search(search string) *FailParams {
	param.search = search
	return param
}

func (param *FailParams) State(state FailStatus) *FailParams {
	param.state = state
	return param
}

func (param *FailParams) Pattern(pattern string) *FailParams {
	param.pattern = pattern
	return param
}

func (param *FailParams) Metadata(metadata string) *FailParams {
	param.metadata = metadata
	return param
}

func (param *FailParams) From(from time.Time) *FailParams {
	param.from = &from
	return param
}

func (param *FailParams) To(to time.Time) *FailParams {
	param.to = &to
	return param
}
