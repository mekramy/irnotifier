package irnotifier

import (
	"time"
)

type Report struct {
	ID    string `json:"id"`
	Pages int    `json:"pages"`
	Cost  int    `json:"cost"`
	Meta  string `json:"meta"`
}

type ClientInfo struct {
	Name     string            `json:"name"`
	Balance  int64             `json:"balance"`
	Queued   int64             `json:"queued"`
	Pendings int64             `json:"pendings"`
	Numbers  map[string]int    `json:"numbers"`
	Patterns map[string]string `json:"patterns"`
}

type Message struct {
	ID        string    `json:"_id"`
	MessageId int64     `json:"message_id"`
	Title     string    `json:"title"`
	From      string    `json:"from"`
	To        string    `json:"to"`
	Meta      string    `json:"meta"`
	Message   string    `json:"message"`
	State     string    `json:"state"`
	Pages     int       `json:"pages"`
	Cost      int       `json:"cost"`
	SentAt    time.Time `json:"sent_at"`
}

type SearchResult struct {
	Search   string    `json:"search"`
	Sort     string    `json:"sort"`
	Order    string    `json:"order"`
	Limit    int       `json:"limit"`
	Page     int       `json:"page"`
	Pages    int       `json:"pages"`
	From     int       `json:"from"`
	To       int       `json:"to"`
	Total    int       `json:"total"`
	Messages []Message `json:"data"`
}
