package fcm


import "testing"

const ServerKey = "ServerKey"

func TestSendWithoutTo(t *testing.T) {
	if len(ServerKey) <= 0 {
		t.Error("Please Server Key Defined")
		return
	}
	f := New(ServerKey)
	results, err := f.Send(&Message{
		Data: map[string]string{
			"DataKey": "DataValue",
		},
		Notification: Notification{
			Title: "title",
			Body: "body",
		},
	})
	if results != nil || err == nil {
		t.Error("results must nil or err is not nil")
	} else if err.Error() != "to\n" {
		t.Error(err.Error())
	}
}

func TestSend(t *testing.T) {
	f := &firebase{serverKey: "abcde"}
	results, err := f.Send(&Message{
		To: "someone",
		Data: map[string]string{
			"DataKey": "DataValue",
		},
		Notification: Notification{
			Title: "title",
			Body: "body",
		},
	})
	if results != nil || err == nil {
		t.Error("results must nil or err is not nil")
	} else if err.Error() != "There was an error authenticating the sender account." {
		t.Error(err.Error())
	}
}

func TestSendNil(t *testing.T) {
	f := &firebase{}
	results, err := f.Send(nil)
	if results != nil || err == nil {
		t.Error("results must nil or err is not nil")
	} else if err.Error() != "message is nil" {
		t.Error(err.Error())
	}
}

func TestNew(t *testing.T) {
	f := New("")
	if f != nil {
		t.Fail()
	}
}