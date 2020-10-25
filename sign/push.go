package sign

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Push(msg, chatID, key string) error {
	msg = url.QueryEscape(msg)
	msg = "chat_id=" + chatID + "&text=" + msg
	req, err := http.NewRequest("POST", "https://api.telegram.org/"+key+"/sendMessage", strings.NewReader(msg))
	if err != nil {
		return fmt.Errorf("Push: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	reps, err := client.Do(req)
	if reps != nil {
		defer reps.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("Push: %w", err)
	}
	t, err := ioutil.ReadAll(reps.Body)
	if err != nil {
		return fmt.Errorf("Push: %w", err)
	}
	var ok isok
	err = json.Unmarshal(t, &ok)
	if err != nil {
		return fmt.Errorf("Push: %w", err)
	}
	if !ok.OK {
		return Pusherr
	}
	return nil
}

var Pusherr = errors.New("推送失败")

type isok struct {
	OK bool `json:"ok"`
}
