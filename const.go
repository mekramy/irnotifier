package irnotifier

// Version
type Version string

const V1 Version = "v1"

// MessageStatus
type MessageStatus string

const StatusQueue MessageStatus = "QUEUE"
const StatusSuspend MessageStatus = "SUSPEND"
const StatusPending MessageStatus = "PENDING"
const StatusInquiry MessageStatus = "INQUIRY"
const StatusFail MessageStatus = "FAIL"
const StatusSent MessageStatus = "SENT"
const StatusUnknown MessageStatus = "UNKNOWN"

// FailStatus
type FailStatus string

const FailExpire FailStatus = "expire"             // منقضی شده
const FailReject FailStatus = "reject"             // رد شده
const FailInsufficient FailStatus = "insufficient" // کمبود اعتبار

// FailSort
type FailSort string

const FailState FailSort = "state"
const FailNumber FailSort = "number"
const FailPattern FailSort = "pattern"
const FailTitle FailSort = "title"
const FailLastTry FailSort = "last_try"
const FailReceiver FailSort = "receiver"
const FailMetadata FailSort = "metadata"
const FailSendAt FailSort = "send_at"
const FailExpiration FailSort = "expiration"

// SentSort
type SentSort string

const SentMessageId SentSort = "message_id"
const SentNumber SentSort = "number"
const SentPattern SentSort = "pattern"
const SentTitle SentSort = "title"
const SentPages SentSort = "pages"
const SentTotal SentSort = "total"
const SentReceiver SentSort = "receiver"
const SentMetadata SentSort = "metadata"
const SentSendAt SentSort = "send_at"

// Order
type Order string

const OrderAsc Order = "asc"
const OrderDesc Order = "desc"

// Limit
type Limit uint

const Limit25 Limit = 25
const Limit50 Limit = 50
const Limit100 Limit = 100
