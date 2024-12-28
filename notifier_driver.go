package irnotifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func NewNotifier(apiKey string, host string, v Version) Notifier {
	client := new(irNotifier)
	if host != "" && host != "/" {
		client.host = strings.TrimRight(host, "/")
	} else {
		client.host = "https://irnotifier.ir"
	}
	client.host = client.host + "/api/" + string(v) + "/"
	client.apiKey = apiKey
	return client
}

type irNotifier struct {
	host   string
	apiKey string
}

func (notifier irNotifier) route(path string) string {
	return notifier.host + path
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
	bearer := strings.Replace(authHeader, "Bearer ", "", 1)
	return notifier.apiKey != "" && bearer == notifier.apiKey
}

func (notifier irNotifier) ParseReport(body []byte) (*Report, error) {
	res := new(Report)
	if err := json.Unmarshal(body, res); err != nil {
		return nil, err
	}
	return res, nil
}

func (notifier irNotifier) Information() (*Information, error) {
	url := notifier.route("information")
	if response, err := notifier.doHttp("GET", url, nil); err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else if !isSuccess(response.StatusCode) {
		return nil, responseError(response.StatusCode)
	} else {
		res := new(Information)
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (notifier irNotifier) Statistic(metadata *string) (*Statistics, error) {
	url := notifier.route("statistic")
	if data, err := json.Marshal(map[string]any{"metadata": metadata}); err != nil {
		return nil, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else if !isSuccess(response.StatusCode) {
		return nil, responseError(response.StatusCode)
	} else {
		res := new(Statistics)
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (notifier irNotifier) Inquiry(id string) (MessageStatus, error) {
	url := notifier.route("inquiry/" + alter(&id, "-"))
	if response, err := notifier.doHttp("GET", url, nil); err != nil {
		return StatusUnknown, err
	} else if response == nil {
		return StatusUnknown, nil
	} else if !isSuccess(response.StatusCode) {
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

func (notifier irNotifier) Queue(parameter *QueueParams) (string, error) {
	url := notifier.route("queue")
	if parameter == nil {
		return "", fmt.Errorf("parameter is required")
	} else if data, err := parameter.ToJson(); err != nil {
		return "", err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return "", err
	} else if response == nil {
		return "", nil
	} else if response.StatusCode == 422 {
		body, _ := io.ReadAll(response.Body)
		return "", fmt.Errorf("%s", string(body))
	} else if !isSuccess(response.StatusCode) {
		return "", responseError(response.StatusCode)
	} else {
		var res string
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&res); err != nil {
			return "", err
		}
		return res, nil
	}
}

func (notifier irNotifier) Requeue(id string, parameter *QueueParams) (bool, error) {
	url := notifier.route("queue/" + alter(&id, "-"))
	if parameter == nil {
		return false, fmt.Errorf("parameter is required")
	} else if data, err := parameter.ToJson(); err != nil {
		return false, err
	} else if response, err := notifier.doHttp("PUT", url, bytes.NewBuffer(data)); err != nil {
		return false, err
	} else if response == nil {
		return false, nil
	} else if response.StatusCode == 422 {
		body, _ := io.ReadAll(response.Body)
		return false, fmt.Errorf("%s", string(body))
	} else if !isSuccess(response.StatusCode) {
		return false, responseError(response.StatusCode)
	} else {
		return true, nil
	}
}

func (notifier irNotifier) Dequeue(id string) (bool, error) {
	url := notifier.route("queue/" + alter(&id, "-"))
	if response, err := notifier.doHttp("DELETE", url, nil); err != nil {
		return false, err
	} else if response == nil {
		return false, nil
	} else if !isSuccess(response.StatusCode) {
		return false, responseError(response.StatusCode)
	} else {
		return true, nil
	}
}

func (notifier irNotifier) Suspend(metadata string, force bool) (int64, error) {
	url := notifier.route("suspend")
	if data, err := json.Marshal(map[string]any{"metadata": metadata, "force": force}); err != nil {
		return 0, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return 0, err
	} else if response == nil {
		return 0, nil
	} else if response.StatusCode == 422 {
		body, _ := io.ReadAll(response.Body)
		return 0, fmt.Errorf("%s", string(body))
	} else if !isSuccess(response.StatusCode) {
		return 0, responseError(response.StatusCode)
	} else {
		var res int64
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&res); err != nil {
			return 0, err
		}
		return res, nil
	}
}

func (notifier irNotifier) Resume(metadata string) (int64, error) {
	url := notifier.route("resume")
	if data, err := json.Marshal(map[string]any{"metadata": metadata}); err != nil {
		return 0, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return 0, err
	} else if response == nil {
		return 0, nil
	} else if response.StatusCode == 422 {
		body, _ := io.ReadAll(response.Body)
		return 0, fmt.Errorf("%s", string(body))
	} else if !isSuccess(response.StatusCode) {
		return 0, responseError(response.StatusCode)
	} else {
		var res int64
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&res); err != nil {
			return 0, err
		}
		return res, nil
	}
}

func (notifier irNotifier) DequeueAll(parameter *DequeueParams) (int64, error) {
	url := notifier.route("dequeue")
	if parameter == nil {
		parameter = DequeueParameter()
	}

	if data, err := parameter.ToJson(); err != nil {
		return 0, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return 0, err
	} else if response == nil {
		return 0, nil
	} else if response.StatusCode == 422 {
		body, _ := io.ReadAll(response.Body)
		return 0, fmt.Errorf("%s", string(body))
	} else if !isSuccess(response.StatusCode) {
		return 0, responseError(response.StatusCode)
	} else {
		var res int64
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(&res); err != nil {
			return 0, err
		}
		return res, nil
	}
}

func (notifier irNotifier) FailList(parameter *FailParams) (*SearchResult[FailMessage], error) {
	url := notifier.route("query/fail")
	if parameter == nil {
		parameter = FailParameter()
	}

	if data, err := parameter.ToJson(); err != nil {
		return nil, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else if !isSuccess(response.StatusCode) {
		return nil, responseError(response.StatusCode)
	} else {
		res := new(SearchResult[FailMessage])
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}

func (notifier irNotifier) SentList(parameter *SentParams) (*SearchResult[SentMessage], error) {
	url := notifier.route("query/sent")
	if parameter == nil {
		parameter = SentParameter()
	}

	if data, err := parameter.ToJson(); err != nil {
		return nil, err
	} else if response, err := notifier.doHttp("POST", url, bytes.NewBuffer(data)); err != nil {
		return nil, err
	} else if response == nil {
		return nil, nil
	} else if !isSuccess(response.StatusCode) {
		return nil, responseError(response.StatusCode)
	} else {
		res := new(SearchResult[SentMessage])
		decoder := json.NewDecoder(response.Body)
		if err := decoder.Decode(res); err != nil {
			return nil, err
		}
		return res, nil
	}
}
