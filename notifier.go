package irnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Notifier interface {
	// IsValidRequest check if request auth access key is valid
	IsValidRequest(authHeader string) bool
	// ParseReport parse delivery report sent by POST method from irnotifier.ir
	ParseReport(body []byte) (*Report, error)
	// Info get client info
	Info() (*ClientInfo, error)
	// Inquiry get message status
	Inquiry(id string) (MessageStatus, error)
	// Sent get sent message list
	Sent(page int, perPage PerPage, sort Sort, order Order, search, meta, from, to string) (*SearchResult, error)
	// Queue send new message
	Queue(pattern, from, to, meta, callback string, sentAt, expiration time.Time, params map[string]string) (string, error)
	// ReQueue update sent message
	ReQueue(messageId, pattern, from, to, meta, callback string, sentAt, expiration time.Time, params map[string]string) (bool, error)
	// UnQueue update sent message
	UnQueue(messageId string) (bool, error)
}

func NewNotifier(apiKey string) Notifier {
	client := new(irNotifier)
	client.apiKey = apiKey
	return client
}

type irNotifier struct {
	apiKey string
}

func (notifier irNotifier) doHttp(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+notifier.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	if response, err := client.Do(req); err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

func (notifier irNotifier) IsValidRequest(authHeader string) bool {
	return notifier.apiKey != "" && strings.Replace(authHeader, "Bearer ", "", 1) == notifier.apiKey
}

func (notifier irNotifier) ParseReport(body []byte) (*Report, error) {
	res := new(Report)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (notifier irNotifier) Info() (*ClientInfo, error) {
	if response, err := notifier.doHttp("GET", BASE+"info", nil); err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else if response.StatusCode != 200 {
		return nil, fmt.Errorf("failed with status %d", response.StatusCode)
	} else {
		res := new(ClientInfo)
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (notifier irNotifier) Inquiry(id string) (MessageStatus, error) {
	if response, err := notifier.doHttp("GET", BASE+"inquiry/"+id, nil); err != nil {
		return StatusUnknown, err
	} else if response == nil {
		return StatusUnknown, nil
	} else if response.StatusCode != 200 {
		return StatusUnknown, fmt.Errorf("failed with status %d", response.StatusCode)
	} else {
		var res MessageStatus
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&res); err != nil {
			return StatusUnknown, err
		}
		return res, nil
	}
}

func (notifier irNotifier) Sent(page int, perPage PerPage, sort Sort, order Order, search, meta, from, to string) (*SearchResult, error) {
	// Generate query
	filters := make(map[string]string)
	if meta != "" {
		filters["meta"] = meta
	}
	if from != "" {
		filters["from"] = from
	}
	if to != "" {
		filters["to"] = to
	}
	if query, err := json.Marshal(map[string]any{
		"page":    page,
		"limit":   perPage,
		"sort":    sort,
		"order":   order,
		"search":  search,
		"filters": filters,
	}); err != nil {
		return nil, err
	} else {
		if response, err := notifier.doHttp("GET", BASE+"sent?qs="+string(query), nil); err != nil {
			return nil, err
		} else if response == nil {
			return nil, nil
		} else if response.StatusCode != 200 {
			return nil, fmt.Errorf("failed with status %d", response.StatusCode)
		} else {
			res := new(SearchResult)
			decoder := json.NewDecoder(response.Body)
			if err := decoder.Decode(res); err != nil {
				return nil, err
			}
			return res, nil
		}
	}
}

func (notifier irNotifier) Queue(pattern, from, to, meta, callback string, sentAt, expiration time.Time, params map[string]string) (string, error) {
	if body, err := json.Marshal(map[string]any{
		"meta":       meta,
		"callback":   callback,
		"number":     from,
		"pattern":    pattern,
		"to":         to,
		"parameters": params,
		"send_at":    sentAt,
		"expiration": expiration,
	}); err != nil {
		return "", err
	} else {
		if response, err := notifier.doHttp("POST", BASE+"queue", bytes.NewBuffer(body)); err != nil {
			return "", err
		} else if response == nil {
			return "", nil
		} else if response.StatusCode == 422 {
			body, _ := io.ReadAll(response.Body)
			return "", fmt.Errorf("%s", string(body))
		} else if response.StatusCode != 200 {
			return "", fmt.Errorf("failed with status %d", response.StatusCode)
		} else {
			var res string
			decoder := json.NewDecoder(response.Body)
			if err := decoder.Decode(&res); err != nil {
				return "", err
			}
			return res, nil
		}
	}
}

func (notifier irNotifier) ReQueue(messageId, pattern, from, to, meta, callback string, sentAt, expiration time.Time, params map[string]string) (bool, error) {
	if body, err := json.Marshal(map[string]any{
		"meta":       meta,
		"callback":   callback,
		"number":     from,
		"pattern":    pattern,
		"to":         to,
		"parameters": params,
		"send_at":    sentAt,
		"expiration": expiration,
	}); err != nil {
		return false, err
	} else {
		if response, err := notifier.doHttp("PUT", BASE+"queue/"+messageId, bytes.NewBuffer(body)); err != nil {
			return false, err
		} else if response == nil {
			return false, nil
		} else if response.StatusCode == 422 {
			body, _ := io.ReadAll(response.Body)
			return false, fmt.Errorf("%s", string(body))
		} else if response.StatusCode != 200 {
			return false, fmt.Errorf("failed with status %d", response.StatusCode)
		} else {
			return true, nil
		}
	}
}

func (notifier irNotifier) UnQueue(messageId string) (bool, error) {
	if response, err := notifier.doHttp("DELETE", BASE+"queue/"+messageId, nil); err != nil {
		return false, err
	} else if response == nil {
		return false, nil
	} else if response.StatusCode != 200 {
		return false, fmt.Errorf("failed with status %d", response.StatusCode)
	} else {
		return true, nil
	}
}
