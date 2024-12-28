package irnotifier

import (
	"time"
)

type Report struct {
	Id       string  `json:"id"`
	Metadata *string `json:"metadata"`
	Pages    int     `json:"pages"`
	Cost     int     `json:"cost"`
}

type Information struct {
	Name     string   `json:"name"`
	Balance  int64    `json:"balance"`
	Numbers  []string `json:"numbers"`
	Patterns []string `json:"patterns"`
}

type Statistics struct {
	Queues    int64 `json:"queues"`
	Suspends  int64 `json:"suspends"`
	Pendings  int64 `json:"pendings"`
	Inquiries int64 `json:"inquiries"`
	Fails     int64 `json:"fails"`
	Sent      int64 `json:"sent"`
}

type FailMessage struct {
	Id         string     `json:"id"`
	State      FailStatus `json:"state"`
	Number     string     `json:"number"`
	Key        string     `json:"pattern"`
	Title      string     `json:"title"`
	LastTry    *time.Time `json:"last_try"`
	Receiver   string     `json:"receiver"`
	Metadata   *string    `json:"metadata"`
	SendAt     time.Time  `json:"send_at"`
	Expiration time.Time  `json:"expiration"`
}

type SentMessage struct {
	Id        string    `json:"id"`
	MessageId int64     `json:"message_id"`
	Number    string    `json:"number"`
	Key       string    `json:"pattern"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	Pages     int       `json:"pages"`
	Total     int       `json:"total"`
	Receiver  string    `json:"receiver"`
	Metadata  *string   `json:"metadata"`
	SendAt    time.Time `json:"send_at"`
}

type SearchResult[T any] struct {
	Search   string `json:"search"`
	Sort     string `json:"sort"`
	Order    string `json:"order"`
	Limit    int    `json:"limit"`
	Page     int    `json:"page"`
	Pages    int    `json:"pages"`
	From     int    `json:"from"`
	To       int    `json:"to"`
	Total    int    `json:"total"`
	Messages []T    `json:"data"`
}
