package irnotifier

import (
	"encoding/json"
	"regexp"
)

// IsAPIError check if error is api related error
func IsAPIError(err error) bool {
	ok, _ := regexp.Match(`failed with status \d{3}`, []byte(err.Error()))
	return ok
}

// IsServerError check if error is 500 error
func IsServerError(err error) bool {
	return err.Error() == "failed with status 500"
}

// IsUnavailableError check if error is 501 or 503 error
func IsUnavailableError(err error) bool {
	return err.Error() == "failed with status 502" ||
		err.Error() == "failed with status 503"
}

// IsAuthError check if error is 401 or 403 error
func IsAuthError(err error) bool {
	return err.Error() == "failed with status 401" || err.Error() == "failed with status 403"
}

// IsNotFoundErr check if error is 404 error
func IsNotFoundErr(err error) bool {
	return err.Error() == "failed with status 404"
}

// IsRequestLimitErr check if error is 429 error
func IsRequestLimitErr(err error) bool {
	return err.Error() == "failed with status 429"
}

// ValidationErr
func ValidationErr(err error) map[string][]string {
	decoded := make(map[string]map[string]string)
	if err := json.Unmarshal([]byte(err.Error()), &decoded); err == nil {
		res := make(map[string][]string)
		for field, errs := range decoded {
			res[field] = make([]string, 0)
			for e := range errs {
				res[field] = append(res[field], e)
			}
		}
		return res
	}
	return nil
}
