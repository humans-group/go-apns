package apns

// easyjson:json
type rawResp struct {
	Reason errorReason `json:"reason"`
}

// API error reasons
// https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/handling_notification_responses_from_apns
type errorReason string

const reasonExpiredProviderToken errorReason = "ExpiredProviderToken"
const reasonBadDeviceToken errorReason = "BadDeviceToken"
