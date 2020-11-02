package sign

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Tosign(cxt context.Context, name, BDUSS, tbs string) error {
	body := "kw=" + name + "&tbs=" + tbs + "&sign=" + enCodeMd5("kw="+name+"tbs="+tbs+"tiebaclient!!!")
	reqs, err := http.NewRequestWithContext(cxt, "POST", SIGNUEL, strings.NewReader(body))
	if err != nil {
		return fmt.Errorf("Tosign: %w", err)
	}
	reqs.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	reqs.Header.Set("Cookie", "BDUSS="+BDUSS)
	rep, err := client.Do(reqs)
	if rep.Body != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return fmt.Errorf("Tosign: %w", err)
	}
	if rep.StatusCode != 200 {
		return Not200
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return fmt.Errorf("Tosign: %w", err)
	}
	var e errcode
	err = json.Unmarshal(b, &e)
	if err != nil {
		return fmt.Errorf("Tosign: %w", err)
	}
	if e.ErrCode != "0" {
		return SignErr
	}
	return nil
}

var SignErr = errors.New("签到失败")

func enCodeMd5(msg string) string {
	h := md5.Sum([]byte(msg))
	return hex.EncodeToString(h[:])
}

type errcode struct {
	ErrCode string `json:"error_code"`
}
