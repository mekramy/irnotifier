package irnotifier

const BASE = "https://irnotifier.ir/api/v1/"

type MessageStatus string

const (
	StatusQueue   MessageStatus = "QUEUE"
	StatusFailed  MessageStatus = "FAILED"
	StatusPending MessageStatus = "PENDING"
	StatusSent    MessageStatus = "SUCCESS"
	StatusUnknown MessageStatus = "UNKNOWN"
)

// Sort
type Sort string

const (
	SortMessageId Sort = "message_id"
	SortTitle     Sort = "title"
	SortFrom      Sort = "from"
	SortTo        Sort = "to"
	SortState     Sort = "state"
	SortPages     Sort = "pages"
	SortCost      Sort = "cost"
	SortSentAt    Sort = "sent_at"
)

// Order
type Order string

const (
	OrderAsc  Order = "asc"
	OrderDesc Order = "desc"
)

// PerPage
type PerPage uint

const (
	PerPage25  PerPage = 25
	PerPage50  PerPage = 50
	PerPage100 PerPage = 100
)
