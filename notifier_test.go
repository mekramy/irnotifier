package irnotifier_test

import (
	"testing"
	"time"

	"github.com/mekramy/irnotifier"
)

const key = "WNJu3mzdvUX4T1fxK19JJ8v5BdavOk7Ez1iABa3mMUSXeEj2cvXLloRLBUiPRCp9dxe24CKyTubLleGB7UsBJV5U3jCBPyK9"

func TestInfo(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.Info(); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestInquiry(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.Inquiry("663a165868d199304c0db6eb"); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestSent(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.Sent(1, irnotifier.PerPage50, irnotifier.SortSentAt, irnotifier.OrderAsc, "", "MyMeta", "", "1402-03-31"); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestQueue(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.Queue("login", "", "09120003265", "", "", time.Now(), time.Now().Add(5*time.Minute), map[string]string{"code": "333333"}); err != nil {
		if vErr := irnotifier.ValidationErrors(err); vErr == nil {
			t.Error("error occurred ", err)
		} else {
			t.Errorf("Validation error: %+v", vErr)
		}
	} else {
		t.Logf("%+v", res)
	}
}

func TestReQueue(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.ReQueue("6640558a91bfe333529bef7a", "login", "", "09366661244", "meta", "", time.Now(), time.Now().Add(5*time.Minute), map[string]string{}); err != nil {
		if vErr := irnotifier.ValidationErrors(err); vErr == nil {
			t.Error("error occurred ", err)
		} else {
			t.Errorf("Validation error: %+v", vErr)
		}
	} else {
		t.Logf("%+v", res)
	}
}

func TestUnQueue(t *testing.T) {
	client := irnotifier.NewNotifier(key)
	if res, err := client.UnQueue("6640558a91bfe333529bef7a"); err != nil {
		if vErr := irnotifier.ValidationErrors(err); vErr == nil {
			t.Error("error occurred ", err)
		} else {
			t.Errorf("Validation error: %+v", vErr)
		}
	} else {
		t.Logf("%+v", res)
	}
}
