package main

import (
	"encoding/json"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/xxmdhs/tiebasign/sign"
)

func main() {
	var ispush bool
	sw := strings.Builder{}
	sw.WriteString("百度贴吧自动签到\n\n")
	msg := `一共需要给 ` + strconv.Itoa(len(BDUSS)) + ` 个账号签到。`
	sw.WriteString(msg + "\n\n")
	for zanhao, v := range BDUSS {
		var ok bool
	finish:
		for i := 0; i < 3; i++ {
			var err1, err2 error
			var wait sync.WaitGroup
			var tbs string
			var list []string
			wait.Add(2)
			go func() {
				tbs, err1 = sign.Getbs(v)
				if err1 != nil {
					log.Println(err1)
				}
				wait.Done()
			}()
			go func() {
				list, err2 = sign.GetFollow(v)
				if err2 != nil {
					log.Println(err2)
				}
				wait.Done()
			}()
			wait.Wait()
			if err1 != nil || err2 != nil {
				continue
			}
			errCh := make(chan error, 20)
			limit := make(chan struct{}, 10)
			msgCh := make(chan string, 20)
			sum := len(list)
			if i == 0 {
				msg := "第" + strconv.Itoa(zanhao+1) + "个账号需要给" + strconv.Itoa(sum) + "个贴吧签到。"
				sw.WriteString(msg + "\n")
			}
			if sum == 0 {
				ok = true
				break
			}
			go func() {
				for _, name := range list {
					limit <- struct{}{}
					go toSign(name, v, tbs, errCh, limit, msgCh)
				}
			}()
			var s int
			for {
				select {
				case msg := <-msgCh:
					sw.WriteString(msg + "\n")
					ispush = true
					s++
					if s == sum {
						ok = true
						break finish
					}
				case err := <-errCh:
					log.Println(err)
					continue finish
				}
			}
		}
		if !ok {
			panic("签到失败")
		}
		msg := "第" + strconv.Itoa(zanhao+1) + "个账号签到完成。"
		sw.WriteString(msg + "\n\n")
	}
	msg = `全部账号签到完成。`
	sw.WriteString(msg + "\n\n")
	if ispush && tgkey != "" {
		msg := strings.TrimSpace(sw.String())
		var ok bool
		for i := 0; i < 3; i++ {
			err := sign.Push(msg, tgchatID, tgkey)
			if err != nil {
				log.Println(err)
				continue
			}
			ok = true
			break
		}
		if !ok {
			panic("推送失败")
		}
	} else {
		log.Println(strings.TrimSpace(sw.String()))
	}
}

func toSign(name, bduss, tbs string, errCh chan<- error, limit <-chan struct{}, msgCh chan<- string) {
	err := sign.Tosign(name, bduss, tbs)
	if err != nil {
		errCh <- err
		<-limit
		return
	}
	msgCh <- name + "吧签到成功"
	<-limit
}

var (
	BDUSS    = make([]string, 0)
	tgkey    string
	tgchatID string
)

func init() {
	c := os.Getenv("BDUSS")
	err := json.Unmarshal([]byte(c), &BDUSS)
	if err != nil {
		panic(err)
	}
	tgkey = os.Getenv("tgkey")
	tgchatID = os.Getenv("tgchatID")
}
