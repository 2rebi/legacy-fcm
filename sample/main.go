package main

import (
	fcm "github.com/RebirthLee/legacy-fcm"
)

func main() {
	firebase := fcm.New("{ServerKey}")
	_, _ = firebase.Send(&fcm.Message{
		MutableContent: true,			// iOS Only
		ContentAvailable: true,			// iOS Only
		Data: map[string]string{
			"Media-Type":"video/mp4",
			"Media":"http://video.url",
		},
		Notification: fcm.Notification{
			Title: "Hello",
			Body: "World",
		},
	})
}
