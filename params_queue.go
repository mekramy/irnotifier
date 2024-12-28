package irnotifier

import (
	"encoding/json"
	"time"
)

func QueueParameter() *QueueParams {
	return &QueueParams{}
}

type QueueParams struct {
	receiver   string
	metadata   *string
	sendAt     time.Time
	expiration time.Time
	callback   *string
	params     map[string]string
	number     *string
	pattern    string
}

func (query QueueParams) ToJson() ([]byte, error) {
	return json.Marshal(map[string]any{
		"receiver":   query.receiver,
		"metadata":   query.metadata,
		"send_at":    query.sendAt,
		"expiration": query.expiration,
		"callback":   query.callback,
		"parameters": query.params,
		"number":     query.number,
		"pattern":    query.pattern,
	})
}

func (query *QueueParams) To(to string) *QueueParams {
	query.receiver = to
	return query
}

func (query *QueueParams) From(from string) *QueueParams {
	query.number = &from
	return query
}

func (query *QueueParams) Pattern(pattern string) *QueueParams {
	query.pattern = pattern
	return query
}

func (query *QueueParams) Metadata(metadata string) *QueueParams {
	query.metadata = &metadata
	return query
}

func (query *QueueParams) SendAt(sendAt time.Time) *QueueParams {
	query.sendAt = sendAt
	return query
}

func (query *QueueParams) Expiration(expiration time.Time) *QueueParams {
	query.expiration = expiration
	return query
}

func (query *QueueParams) Callback(callback string) *QueueParams {
	query.callback = &callback
	return query
}

func (query *QueueParams) Parameters(params map[string]string) *QueueParams {
	query.params = params
	return query
}

func (query *QueueParams) AddParameter(name, value string) *QueueParams {
	if query.params == nil {
		query.params = make(map[string]string)
	}
	query.params[name] = value
	return query
}
