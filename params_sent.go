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
	sort     FailSort
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

func (param *SentParams) Page(page int) {
	param.page = page
}

func (param *SentParams) Limit(limit Limit) {
	param.limit = limit
}

func (param *SentParams) Sort(sort FailSort) {
	param.sort = sort
}

func (param *SentParams) Order(order Order) {
	param.order = order
}

func (param *SentParams) Search(search string) {
	param.search = search
}

func (param *SentParams) Pattern(pattern string) {
	param.pattern = pattern
}

func (param *SentParams) Metadata(metadata string) {
	param.metadata = metadata
}

func (param *SentParams) From(from time.Time) {
	param.from = &from
}

func (param *SentParams) To(to time.Time) {
	param.to = &to
}
