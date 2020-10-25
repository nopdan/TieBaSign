package sign

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Getbs(BDUSS string) (string, error) {
	cookie := "BDUSS=" + BDUSS
	//return {"tbs":"************","is_login":1}
	b, err := httpget(TBSURL, cookie)
	if err != nil {
		return "", fmt.Errorf("Getbs: %w", err)
	}
	var t TBS
	err = json.Unmarshal(b, &t)
	if err != nil {
		return "", fmt.Errorf("Getbs: %w", err)
	}
	if t.IsLogin != 1 {
		return "", BDUSSInvalid
	}
	return t.Tbs, nil
}

var BDUSSInvalid = errors.New("BDUSS Invalid")

type TBS struct {
	Tbs     string `json:"tbs"`
	IsLogin int    `json:"is_login"`
}
