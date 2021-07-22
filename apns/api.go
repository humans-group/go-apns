package apns

import (
	"encoding/json"
	"time"
)

type Notification struct {
	// An optional canonical UUID that identifies the notification. The
	// canonical form is 32 lowercase hexadecimal digits, displayed in five
	// groups separated by hyphens in the form 8-4-4-4-12. An example UUID is as
	// follows:
	//
	//  123e4567-e89b-12d3-a456-42665544000
	//
	// If you don't set this, a new UUID is created by APNs and returned in the
	// response.
	ApnsID string

	// A string which allows multiple notifications with the same collapse
	// identifier to be displayed to the user as a single notification. The
	// value should not exceed 64 bytes.
	CollapseID string

	// A string containing hexadecimal bytes of the device token for the target
	// device.
	DeviceToken string

	// The topic of the remote notification, which is typically the bundle ApnsID
	// for your app. The certificate you create in the Apple Developer Member
	// Center must include the capability for this topic. If your certificate
	// includes multiple topics, you must specify a value for this header. If
	// you omit this header and your APNs certificate does not specify multiple
	// topics, the APNs server uses the certificate’s Subject as the default
	// topic.
	Topic string

	// An optional time at which the notification is no longer valid and can be
	// discarded by APNs. If this value is in the past, APNs treats the
	// notification as if it expires immediately and does not store the
	// notification or attempt to redeliver it. If this value is left as the
	// default (ie, Expiration.IsZero()) an expiration header will not added to
	// the http request.
	Expiration time.Time

	// The priority of the notification. Specify ether apns.PriorityHigh (10) or
	// apns.PriorityLow (5) If you don't set this, the APNs server will set the
	// priority to 10.
	Priority int

	// A byte array containing the JSON-encoded payload of this push notification.
	// Refer to "The Remote Notification Payload" section in the Apple Local and
	// Remote Notification Programming Guide for more info.
	Payload Payload

	// The pushtype of the push notification. If this values is left as the
	// default an apns-push-type header with value 'alert' will be added to the
	// http request.
	PushType PushType
}

// PushType defines the value for the apns-push-type header
type PushType string

const (
	// PushTypeAlert is used for notifications that trigger a user interaction —
	// for example, an alert, badge, or sound. If you set this push type, the
	// apns-topic header field must use your app’s bundle ApnsID as the topic. The
	// alert push type is required on watchOS 6 and later. It is recommended on
	// macOS, iOS, tvOS, and iPadOS.
	PushTypeAlert PushType = "alert"

	// PushTypeBackground is used for notifications that deliver content in the
	// background, and don’t trigger any user interactions. If you set this push
	// type, the apns-topic header field must use your app’s bundle ApnsID as the
	// topic. The background push type is required on watchOS 6 and later. It is
	// recommended on macOS, iOS, tvOS, and iPadOS.
	PushTypeBackground PushType = "background"

	// PushTypeVOIP is used for notifications that provide information about an
	// incoming Voice-over-IP (VoIP) call. If you set this push type, the
	// apns-topic header field must use your app’s bundle ApnsID with .voip appended
	// to the end. If you’re using certificate-based authentication, you must
	// also register the certificate for VoIP services. The voip push type is
	// not available on watchOS. It is recommended on macOS, iOS, tvOS, and
	// iPadOS.
	PushTypeVOIP PushType = "voip"

	// PushTypeComplication is used for notifications that contain update
	// information for a watchOS app’s complications. If you set this push type,
	// the apns-topic header field must use your app’s bundle ApnsID with
	// .complication appended to the end. If you’re using certificate-based
	// authentication, you must also register the certificate for WatchKit
	// services. The complication push type is recommended for watchOS and iOS.
	// It is not available on macOS, tvOS, and iPadOS.
	PushTypeComplication PushType = "complication"

	// PushTypeFileProvider is used to signal changes to a File Provider
	// extension. If you set this push type, the apns-topic header field must
	// use your app’s bundle ApnsID with .pushkit.fileprovider appended to the end.
	// The fileprovider push type is not available on watchOS. It is recommended
	// on macOS, iOS, tvOS, and iPadOS.
	PushTypeFileProvider PushType = "fileprovider"

	// PushTypeMDM is used for notifications that tell managed devices to
	// contact the MDM server. If you set this push type, you must use the topic
	// from the UID attribute in the subject of your MDM push certificate.
	PushTypeMDM PushType = "mdm"
)

// Payload repsresents a data structure for APN notification.
// easyjson:json
type Payload struct {
	APS  APS             `json:"aps"`
	Data json.RawMessage `json:"data"`
}

// APS is Apple's reserved payload.
// easyjson:json
type APS struct {
	// Alert dictionary.
	Alert Alert `json:"alert,omitempty"`

	// Badge to display on the app icon.
	Badge int `json:"badge,omitempty"`

	// Sound is the name of a sound file to play as an alert.
	Sound string `json:"sound,omitempty"`

	// Content available apps launched in the background or resumed.
	ContentAvailable int `json:"content-available,omitempty"`

	// Category identifier for custom actions in iOS 8 or newer.
	Category string `json:"category,omitempty"`

	// ThreadID presents the app-specific identifier for grouping notifications.
	ThreadID string `json:"thread-id,omitempty"`

	// MutableContent presents the flag that lets user’s device know
	// that it should run the corresponding service app extension (if the value equals 1).
	MutableContent int `json:"mutable-content,omitempty"`
}

// Alert represents alert dictionary.
// easyjson:json
type Alert struct {
	Title        string   `json:"title,omitempty"`
	Body         string   `json:"body,omitempty"`
	TitleLocKey  string   `json:"title-loc-key,omitempty"`
	TitleLocArgs []string `json:"title-loc-args,omitempty"`
	ActionLocKey string   `json:"action-loc-key,omitempty"`
	LocKey       string   `json:"loc-key,omitempty"`
	LocArgs      []string `json:"loc-args,omitempty"`
	LaunchImage  string   `json:"launch-image,omitempty"`
}

// easyjson:json
type Response struct {
	Reason ErrorReason `json:"reason"`
}

// ErrorReason API error reasons
// https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
type ErrorReason string

const (
	ReasonExpiredProviderToken ErrorReason = "ExpiredProviderToken"
	ReasonBadDeviceToken       ErrorReason = "BadDeviceToken"
	ReasonCodeUnregistered     ErrorReason = "Unregistered"
)
