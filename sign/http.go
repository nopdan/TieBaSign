package sign

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	client = http.Client{Timeout: 5 * time.Second}
)

func httpget(url, cookie string) ([]byte, error) {
	reqs, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("httpget: %w", err)
	}
	reqs.Header.Set("Accept", "*/*")
	reqs.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	reqs.Header.Set("Cookie", cookie)
	rep, err := client.Do(reqs)
	if rep != nil {
		defer rep.Body.Close()
	}
	if err != nil {
		return nil, fmt.Errorf("httpget: %w", err)
	}
	if rep.StatusCode != 200 {
		return nil, Not200{rep.Status}
	}
	b, err := ioutil.ReadAll(rep.Body)
	if err != nil {
		return nil, fmt.Errorf("httpget: %w", err)
	}
	return b, nil
}

type Not200 struct {
	msg string
}

func (n Not200) Error() string {
	return "not 200 is :" + n.msg
}
