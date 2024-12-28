package irnotifier_test

import (
	"testing"
	"time"

	"github.com/mekramy/irnotifier"
)

const host = "http://127.0.0.1:8888"
const id = "e47d0d70-dd82-4dc5-b80d-2f719311bdda"
const key = "N0EBCZ0H6SCZ9JCULZJGHGRWJFIQNO2SLF7AGP5RYQP5YEJO6HVJTJ4WWET0VS5M"

func notifier() irnotifier.Notifier {
	return irnotifier.NewNotifier(key, host, irnotifier.V1)
}

func TestInformation(t *testing.T) {
	if res, err := notifier().Information(); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestStatistic(t *testing.T) {
	if res, err := notifier().Statistic(nil); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestInquiry(t *testing.T) {
	if res, err := notifier().Inquiry(id); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%+v", res)
	}
}

func TestQueue(t *testing.T) {
	params := irnotifier.QueueParameter().
		To("09121230004").
		Metadata("test").
		Pattern("test").
		SendAt(time.Now()).
		Expiration(time.Now().Add(time.Minute*5)).
		AddParameter("name", "John Doe").
		AddParameter("number", "123").
		AddParameter("date", "Jun")
	if res, err := notifier().Queue(params); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%s", res)
	}
}

func TestReQueue(t *testing.T) {
	params := irnotifier.QueueParameter().
		To("09121230000").
		Metadata("test").
		Pattern("test").
		SendAt(time.Now()).
		Expiration(time.Now().Add(time.Minute*5)).
		AddParameter("name", "John Doe").
		AddParameter("number", "123").
		AddParameter("date", "Jun")
	if res, err := notifier().Requeue("e610ca06-a98c-48b5-aeb1-732094af5620", params); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%v", res)
	}
}

func TestDequeue(t *testing.T) {
	if res, err := notifier().Dequeue("1ad7a837-7089-4789-a15a-23c2501a1551"); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%v", res)
	}
}

func TestSuspend(t *testing.T) {
	if res, err := notifier().Suspend("Hi", true); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%d", res)
	}
}

func TestResume(t *testing.T) {
	if res, err := notifier().Resume(""); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%d", res)
	}
}

func TestDequeueAll(t *testing.T) {
	params := irnotifier.DequeueParameter().
		Receiver("09121230004").
		Metadata("test").
		Number("123").
		Pattern("test").
		From(time.Now()).
		To(time.Now().Add(time.Minute * 5))
	if res, err := notifier().DequeueAll(params); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%v", res)
	}
}

func TestFail(t *testing.T) {
	if res, err := notifier().FailList(nil); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%v", res.Total)
	}
}

func TestSent(t *testing.T) {
	if res, err := notifier().SentList(nil); err != nil {
		t.Error("error occurred ", err)
	} else {
		t.Logf("%v", res.Total)
	}
}
