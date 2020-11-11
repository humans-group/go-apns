package apns

// easyjson:json
type RawResp struct {
	Reason ErrorReason `json:"reason"`
}

// API error reasons
// https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
type ErrorReason string

const reasonExpiredProviderToken ErrorReason = "ExpiredProviderToken"
