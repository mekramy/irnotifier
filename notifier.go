package irnotifier

type Notifier interface {
	// IsValidRequest check if request auth access key is valid
	IsValidRequest(authHeader string) bool
	// ParseReport parse delivery report sent by POST method from irnotifier.ir
	ParseReport(body []byte) (*Report, error)
	// Information get client info
	// [401] unauthorized
	// [500] operation failed
	Information() (*Information, error)
	// Statistic get client statistics
	// [400] form bind failed
	// [401] unauthorized
	// [500] operation failed
	Statistic(metadata *string) (*Statistics, error)
	// Inquiry get message status
	// [401] unauthorized
	// [404] message not found
	// [500] operation failed
	Inquiry(id string) (MessageStatus, error)
	// Queue send new message
	// [201] queued message id
	// [400] form bind failed
	// [401] unauthorized
	// [422] validation failed
	// [500] operation failed
	Queue(parameter *QueueParams) (string, error)
	// Requeue update queued message
	// [400] form bind failed
	// [401] unauthorized
	// [404] message not found
	// [422] validation failed
	// [500] operation failed
	Requeue(id string, parameter *QueueParams) (bool, error)
	// Dequeue delete queued message
	// [200] ok
	// [401] unauthorized
	// [404] message not found
	// [500] operation failed
	Dequeue(id string) (bool, error)
	// Suspend suspend queued messages
	// [200] updated count
	// [400] form bind failed
	// [401] unauthorized
	// [500] operation failed
	Suspend(metadata string, force bool) (int64, error)
	// Resume resume suspended messages
	// [200] updated count
	// [400] form bind failed
	// [401] unauthorized
	// [500] operation failed
	Resume(metadata string) (int64, error)
	// DequeueAll delete queued messages
	// [200] updated count
	// [400] form bind failed
	// [401] unauthorized
	// [500] operation failed
	DequeueAll(parameter *DequeueParams) (int64, error)
	// [401] unauthorized
	// [500] operation failed
	FailList(parameter *FailParams) (*SearchResult[FailMessage], error)
	// SentList get sent message list
	// [401] unauthorized
	// [500] operation failed
	SentList(parameter *SentParams) (*SearchResult[SentMessage], error)
}
