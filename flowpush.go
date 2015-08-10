package flowpush

import (
	"fmt"
	"net/http"
	"net/url"
)

func pushMessage(flowAPIKey, message, sender string, threadID int64) error {
	v := url.Values{}
	v.Set("content", message)
	v.Set("external_user_name", sender)
	if threadID != 0 {
		v.Set("message_id", string(threadID))
	}

	pushURL := fmt.Sprintf("https://api.flowdock.com/v1/messages/chat/%s", flowAPIKey)
	resp, err := http.PostForm(pushURL, v)
	defer resp.Body.Close()

	if err != nil {
		return err
	}
	return nil
}

func PushMessageToFlowWithKey(flowAPIKey, message, sender string) error {
	return pushMessage(flowAPIKey, message, sender, 0)
}

func ReplyToThreadInFlowWithKey(flowAPIKey, message, sender string, threadID int64) error {
	return pushMessage(flowAPIKey, message, sender, threadID)
}
