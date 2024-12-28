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

func (param *FailParams) Page(page int) {
	param.page = page
}

func (param *FailParams) Limit(limit Limit) {
	param.limit = limit
}

func (param *FailParams) Sort(sort FailSort) {
	param.sort = sort
}

func (param *FailParams) Order(order Order) {
	param.order = order
}

func (param *FailParams) Search(search string) {
	param.search = search
}

func (param *FailParams) State(state FailStatus) {
	param.state = state
}

func (param *FailParams) Pattern(pattern string) {
	param.pattern = pattern
}

func (param *FailParams) Metadata(metadata string) {
	param.metadata = metadata
}

func (param *FailParams) From(from time.Time) {
	param.from = &from
}

func (param *FailParams) To(to time.Time) {
	param.to = &to
}
