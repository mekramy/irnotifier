package irnotifier

import (
	"encoding/json"
	"time"
)

func DequeueParameter() *DequeueParams {
	return &DequeueParams{}
}

type DequeueParams struct {
	receiver *string
	metadata *string
	number   *string
	pattern  *string
	from     *time.Time
	to       *time.Time
}

func (query DequeueParams) ToJson() ([]byte, error) {
	return json.Marshal(map[string]any{
		"receiver": query.receiver,
		"metadata": query.metadata,
		"number":   query.number,
		"pattern":  query.pattern,
		"from":     query.from,
		"to":       query.to,
	})
}

func (query *DequeueParams) Receiver(receiver string) *DequeueParams {
	query.receiver = &receiver
	return query
}

func (query *DequeueParams) Metadata(metadata string) *DequeueParams {
	query.metadata = &metadata
	return query
}

func (query *DequeueParams) Number(from string) *DequeueParams {
	query.number = &from
	return query
}

func (query *DequeueParams) Pattern(pattern string) *DequeueParams {
	query.pattern = &pattern
	return query
}

func (query *DequeueParams) From(from time.Time) *DequeueParams {
	query.from = &from
	return query
}

func (query *DequeueParams) To(to time.Time) *DequeueParams {
	query.to = &to
	return query
}
