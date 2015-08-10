package flowpush

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const flowdockAPIURL = "https://api.flowdock.com"

func flowpushGET(apiKey, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(apiKey, "BATMAN")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

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
