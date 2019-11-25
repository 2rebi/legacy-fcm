package fcm // import "github.com/RebirthLee/legacy-fcm"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	NORMAL 	PriorityType = "normal"
	HIGH 	PriorityType = "high"
)

type (
	PriorityType string

	Firebase interface {
		Send(*Message) (*Response, error)
	}

	firebase struct {
		serverKey string
	}
	
	Message struct {
		To 			string 			`json:"to,omitempty"`
		Ids 			[]string		`json:"registration_ids,omitempty"`
		Condition 		string 			`json:"condition,omitempty"`
		CollapseKey 		string 			`json:"collapse_key,omitempty"`
		Priority 		PriorityType		`json:"priority,omitempty"`
		ContentAvailable 	bool 			`json:"content_available,omitempty"`
		MutableContent 		bool 			`json:"mutable_content,omitempty"`
		TTL			time.Duration		`json:"-"`
		TimeToLive 		string 			`json:"ttl,omitempty"`
		DryRun 			bool 			`json:"dry_run,omitempty"`
		Android			AndroidConfig		`json:"android,omitempty"`	// Android support
		//Apns			ApnsConfig		`json:"apns,omitempty"`		// ToDo : iOS support
		Webpush			WebpushConfig		`json:"webpush,omitempty"`	// Webpush Javascript support
		Data 			map[string]interface{}	`json:"data,omitempty"`
		Notification		Notification		`json:"notification,omitempty"`
	}

	Notification struct {
		Title 			string `json:"title,omitempty"` 	// All Platform Supported
		Body 			string `json:"body,omitempty"`		// All Platform Supported
		ClickAction 		string `json:"click_action,omitempty"` 	// All Platform Supported
		Icon 			string `json:"icon,omitempty"` 		// Support by Android, Web.
		AndroidChannelId	string `json:"badge,omitempty"` 	// Android Only
		Tag 			string `json:"tag,omitempty"` 		// Android Only
		Color 			string `json:"color,omitempty"` 	// Android Only
		SubTitle 		string `json:"subtitle,omitempty"` 	// iOS Only
		Badge 			string `json:"badge,omitempty"` 	// iOS Only


		Sound 			string `json:"sound,omitempty"`		// Support by Android, iOS.
		//Todo
		//https://firebase.google.com/docs/cloud-messaging/http-server-ref
		//sound	Optional, string or Dictionary
		//The sound to play when the device receives the notification.
		//
		//String specifying sound files in the main bundle of the client app or in the Library/Sounds folder of the app's data container. See the iOS Developer Library for more information.
		//
		//For critical notifications use a dictionary that contains sound information for critical alerts. See the iOS Developer Library for keys required. For regular notifications, use the sound string instead.

		//Todo
		//BodyLocKey string `json:"body_loc_key,omitempty"` // Support by Android, iOS.
		//body_loc_key	Optional, string
		//The key to the body string in the app's string resources to use to localize the body text to the user's current localization.
		//
		//Corresponds to loc-key in the APNs payload.
		//
		//See Payload Key Reference and Localizing the Content of Your Remote Notifications for more information.

		//Todo
		//BodyLocArgs string `json:"body_loc_args,omitempty"` // Support by Android, iOS.
		//body_loc_args	Optional, JSON array as string
		//Variable string values to be used in place of the format specifiers in body_loc_key to use to localize the body text to the user's current localization.
		//
		//Corresponds to loc-args in the APNs payload.
		//
		//See Payload Key Reference and Localizing the Content of Your Remote Notifications for more information.

		//Todo
		//TitleLocKey string `json:"title_loc_key,omitempty"` // Support by Android, iOS.
		//title_loc_key	Optional, string
		//The key to the title string in the app's string resources to use to localize the title text to the user's current localization.
		//
		//Corresponds to title-loc-key in the APNs payload.
		//
		//See Payload Key Reference and Localizing the Content of Your Remote Notifications for more information.

		//Todo
		//TitleLocArgs string `json:"title_loc_args,omitempty"` // Support by Android, iOS.
		//title_loc_args	Optional, JSON array as string
		//Variable string values to be used in place of the format specifiers in title_loc_key to use to localize the title text to the user's current localization.
		//
		//Corresponds to title-loc-args in the APNs payload.
		//
		//See Payload Key Reference and Localizing the Content of Your Remote Notifications for more information.

	}

	AndroidConfig struct {
		CollapseKey 		string 			`json:"collapse_key,omitempty"`
		Priority 		PriorityType 		`json:"priority,omitempty"`
		TTL			string			`json:"ttl,omitempty"`				// ex) 3600s
		RestrictedPackageName	string			`json:"restricted_package_name,omitempty"`
		Notification		AndroidNotification	`json:"notification,omitempty"`
		//FcmOptions		AndroidFcmOptions	`json:"fcm_options,omitempty"`	// ToDo : AndroidFcmOptions
	}
	
	AndroidNotification struct {
		Title			string		`json:"title,omitempty"`
		Body			string		`json:"body,omitempty"`
		Icon			string		`json:"icon,omitempty"`
		Color			string		`json:"color,omitempty"`
		Sound			string		`json:"sound,omitempty"`
		Tag			string		`json:"tag,omitempty"`
		ClickAction		string		`json:"click_action,omitempty"`
		BodyLocKey		string		`json:"body_loc_key,omitempty"`
		BodyLocArgs		[]string	`json:"body_loc_args,omitempty"`
		TitleLocKey		string		`json:"title_loc_key,omitempty"`
		TitleLocArgs		[]string	`json:"title_loc_args,omitempty"`
		ChannelID		string		`json:"channel_id,omitempty"`
		Ticker			string		`json:"ticker,omitempty"`
		Sticky			bool		`json:"sticky,omitempty"`
		EventTime		string		`json:"event_time,omitempty"`
		LocalOnly		bool		`json:"local_only,omitempty"`
		//NotificationPriority	uint		`json:"notication_priority,omitempty"`		// ToDo : NotificationPriority
		DefaultSound		bool		`json:"default_sound,omitempty"`
		DefaultVibrateTimings	bool		`json:"default_vibrate_timings,omitempty"`
		DefaultLightSettings	bool		`json:"default_light_settings,omitempty"`
		VibrateTimings		[]string	`json:"vibrate_timings,omitempty"`
		NotificationCount	uint		`json:"notification_count,omitempty"`
		//LightSettings		LightSettings	`json:"light_settings,omitempty"`		// ToDo : LightSettings
		Image			string		`json:"image,omitempty"`
	}
	
	WebpushConfig struct {
		Headers		map[string]string	`json:"headers,omitempty"`		// ex) {"TTL":"3600"}
		//Notification	WebpushNotification	`json:"notification,omitempty"`		// ToDo: WebpushNotification
		//FcmOptions	WebpushFcmOptions	`json:"fcm_options,omitempty"`
	}
	
	WebpushFcmOptions struct {
		Link		string			`json:"link"`	// The link to open when the user clicks on the notification. For all URL values, HTTPS is required.
	}
	
	Response struct {
		MulticastId	uint64			`json:"multicast_id"`
		Success		uint			`json:"success"`
		Failure		uint			`json:"failure"`
		CanonicalIds	uint			`json:"canonical_ids"`
		Results		[]MessageResult		`json:"results"`
	}

	MessageResult struct {
		MessageId 	string `json:"message_id,omitempty"`
		Error 		string `json:"error,omitempty"`
	}
)

func New(serverKey string) Firebase {
	if len(serverKey) > 0 {
		return &firebase{
			serverKey: serverKey,
		}
	}
	return nil
}

func (firebase *firebase) Send(message *Message) (*Response, error) {
	if message == nil {
		return nil, errors.New("message is nil")
	}

	if message.TTL >= time.Second {
		message.TimeToLive = fmt.Sprintf("%ds", message.TTL/time.Second)
	}

	body, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, "https://fcm.googleapis.com/fcm/send", bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "key="+firebase.serverKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		response := &Response{}
		body, _ := ioutil.ReadAll(resp.Body)
		_ = json.Unmarshal(body, response)
		return response, nil
	} else if resp.StatusCode == http.StatusUnauthorized {
		return nil, errors.New("There was an error authenticating the sender account.")
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(body))
	}
}
